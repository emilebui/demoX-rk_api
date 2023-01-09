package repositories

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type RedisRepo struct {
	c        *redis.Client
	queueKey string
}

func NewRedisRepo(client *redis.Client, queueKey string) Repo {
	return &RedisRepo{
		c:        client,
		queueKey: queueKey,
	}
}

func (r *RedisRepo) PushMessage(msg string) error {

	err := r.c.LPush(context.Background(), r.queueKey, msg).Err()
	if err != nil {
		fmt.Printf("Error pushing message: %s - %v\n", msg, err)
		return err
	}

	return nil
}
