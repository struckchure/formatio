package lib

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type IRedis interface {
	SetItem(string, string) error
	GetItem(string) (*string, error)
	DeleteItem(string) error
}

type Redis struct {
	rdb *redis.Client
}

// SetItem implements RedisInterface.
func (r *Redis) SetItem(key string, value string) error {
	err := r.rdb.Set(context.Background(), key, value, 30*time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

// GetItem implements RedisInterface.
func (r *Redis) GetItem(key string) (*string, error) {
	value, err := r.rdb.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("%s does not exist", key)
	} else if err != nil {
		return nil, err
	}

	return &value, nil
}

// DeleteItem implements RedisInterface.
func (r *Redis) DeleteItem(key string) error {
	result, err := r.rdb.Del(context.Background(), key).Result()
	if err != nil {
		return err
	}

	if result == 0 {
		return fmt.Errorf("key '%s' was not found", key)
	}

	return nil
}

func NewRedis(rdb *redis.Client) IRedis {
	return &Redis{rdb: rdb}
}

func NewRedisConnection(env Env) *redis.Client {
	opts, err := redis.ParseURL(env.REDIS_URL)
	if err != nil {
		log.Fatal("err connecting to redis: ", err)
	}

	log.Println("redis is connected ...")

	return redis.NewClient(opts)
}
