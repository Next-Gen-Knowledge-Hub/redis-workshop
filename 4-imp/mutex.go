package ch11

import (
	"context"
	_ "embed"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

//go:embed lock_acquire.lua
var lockScrypt string

//go:embed lock_release.lua
var unlockScrypt string

type Mutex struct {
	client *redis.Client
	key    string
}

func NewMutex() *Mutex {
	m := Mutex{
		client: redis.NewClient(&redis.Options{
			Addr: "127.0.0.1:6379",
			DB:   4,
		}),
		key: uuid.NewString()[0:8],
	}

	return &m
}

func (m *Mutex) AcquireLock(ctx context.Context) (lockId string, err error) {
	lockId = uuid.NewString()
	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
			{
				var ok bool
				ok, err = m.client.Eval(ctx, lockScrypt, []string{fmt.Sprintf("lock:%s", m.key)}, lockId, 10).Bool()

				if err != nil && err != redis.Nil {
					return "", err
				}

				if !ok {
					time.Sleep(50 * time.Millisecond)
					continue
				}

				return lockId, nil
			}
		}
	}
}

func (m *Mutex) ReleaseLock(ctx context.Context, id string) (err error) {
	var ok bool
	ok, err = m.client.Eval(ctx, unlockScrypt, []string{fmt.Sprintf("lock:%s", m.key)}, id).Bool()
	if err != nil && err != redis.Nil {
		return err
	}

	if !ok {
		return fmt.Errorf("not ok !!")
	}

	return nil
}
