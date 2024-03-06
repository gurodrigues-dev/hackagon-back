package repository

import (
	"context"
	"fmt"
	"gin/config"
	"time"

	"github.com/go-redis/redis"
)

type Redis struct {
	conn *redis.Client
}

func NewRedisClient() (*Redis, error) {

	conf := config.Get()

	client := redis.NewClient(&redis.Options{
		Addr:     conf.Cache.Address,
		Password: conf.Cache.Password,
		DB:       0,
	})

	repo := &Redis{
		conn: client,
	}

	return repo, nil

}

func (r *Redis) SaveRedis(ctx context.Context, key, value string) error {

	err := r.conn.Set(key, value, 10*time.Minute).Err()

	if err != nil {
		return err
	}

	return nil

}

func (r *Redis) VerifyToken(ctx context.Context, token, email string) error {

	value, err := r.conn.Get(token).Result()

	if err != nil {
		return err
	}

	matchValue := value == email

	if !matchValue {
		return fmt.Errorf("email not true")
	}

	return nil
}
