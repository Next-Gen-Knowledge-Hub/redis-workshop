package imp

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// pres
// 1seconds, 2seconds, 5seconds, ...
var pres = []int{1, 2, 5, 10, 60, 300}

type RedisCounter struct {
	client *redis.Client
}

func NewRedisCounter() *RedisCounter {
	c := RedisCounter{
		client: redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			DB:   0,
		}),
	}

	return &c
}

func (c *RedisCounter) UpdateCounter(name string, count int64) {
	now := time.Now().Unix()

	pipe := c.client.Pipeline()
	for i := 0; i < len(pres); i++ {
		pnow := int(now/int64(pres[i])) * pres[i]
		hash := fmt.Sprintf("%d:%s", pres[i], name)
		pipe.ZAdd(context.TODO(), "known", redis.Z{
			Score:  0,
			Member: hash,
		})
		pipe.HIncrBy(context.TODO(), fmt.Sprintf("counter:%s", hash), fmt.Sprintf("%d", pnow), count)
	}

	_, err := pipe.Exec(context.TODO())
	if err != nil {
		panic(err)
	}
}

func (c *RedisCounter) GetCounter(name string, perci int) {
	hash := fmt.Sprintf("counter:%d:%s", pres[perci], name)
	data, err := c.client.HGetAll(context.TODO(), hash).Result()
	if err != nil {
		panic(err)
	}

	for k, v := range data {
		fmt.Println(k, v)
	}
}
