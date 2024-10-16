package imp

import (
	"context"
	_ "embed"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

//go:embed semaphor_acquire.lua
var acquireSemaphorScrypt string

//go:embed semaphor_release.lua
var releaseSemaphorScrypt string

type Semaphor struct {
	client *redis.Client
	count  int
	id     string
}

func NewSemaphor(count int) *Semaphor {
	s := Semaphor{
		client: redis.NewClient(&redis.Options{
			Addr: "127.0.0.1:6379",
			DB:   4,
		}),
		count: count,
		id:    uuid.NewString()[0:8],
	}

	return &s
}

func (s *Semaphor) AcquireLock(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			{
				ok, err := s.client.Eval(ctx, acquireSemaphorScrypt, []string{fmt.Sprintf("smaphor:%s", s.id)}, s.count, 10).Bool()

				if err != nil && err != redis.Nil {
					return err
				}

				if !ok {
					time.Sleep(50 * time.Microsecond)
					continue
				}

				return nil
			}
		}
	}
}

func (s *Semaphor) ReleaseLock(ctx context.Context) error {
	ok, err := s.client.Eval(ctx, releaseSemaphorScrypt, []string{fmt.Sprintf("smaphor:%s", s.id)}).Bool()
	if err != nil && err != redis.Nil {
		return err
	}

	if !ok {
		return fmt.Errorf("semaphor no ok !!")
	}

	return nil
}
