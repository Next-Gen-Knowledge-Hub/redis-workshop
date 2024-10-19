package imp

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type PushRedisRepo struct {
	client *redis.Client
}

func NewPushRedisRepo() *PushRedisRepo {
	p := PushRedisRepo{
		client: redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			DB:   4,
		}),
	}

	return &p
}

func (p *PushRedisRepo) SendPush(ctx context.Context, uid string, message string) error {
	messageKey := fmt.Sprintf("msg:%s", uuid.NewString()[0:5])

	pipe := p.client.Pipeline()
	_, err := pipe.Set(ctx, messageKey, message, 0).Result()
	if err != nil {
		return err
	}

	clientPushListKey := fmt.Sprintf("pushs:%s", uid)
	_, err = pipe.RPush(ctx, clientPushListKey, messageKey).Result()
	if err != nil {
		return err
	}

	cmd, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}

	for _, c := range cmd {
		if c.Err() != nil {
			return c.Err()
		}
	}

	return nil
}

func (p *PushRedisRepo) GetPushs(ctx context.Context, uid string, count int) (msgs []string, err error) {
	clientPushListKey := fmt.Sprintf("pushs:%s", uid)

	pushCount, err := p.client.LLen(ctx, clientPushListKey).Result()
	if err != nil {
		return nil, err
	}

	if pushCount == 0 {
		return nil, nil
	}

	messageKey := make([]string, 0, count)

	for i := 0; i < count && i < int(pushCount); i++ {
		key, err := p.client.LPop(ctx, clientPushListKey).Result()
		if err != nil {
			return nil, err
		}

		messageKey = append(messageKey, key)
	}

	msgs = make([]string, 0, count)
	for _, msgKey := range messageKey {
		msg, err := p.client.Get(ctx, msgKey).Result()
		if err != nil {
			return nil, err
		}

		msgs = append(msgs, msg)

		_, _ = p.client.Expire(ctx, msgKey, time.Second).Result()
	}

	return msgs, nil
}
