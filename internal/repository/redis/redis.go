package redis

import (
	"blog/config"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var redisClient *Redis

type Redis struct {
	client *redis.Client
}

func Init() {
	c := redis.NewClient(&redis.Options{
		DB:           config.Conf.Redis.Db,
		Addr:         config.Conf.Redis.Addr,
		Password:     config.Conf.Redis.Password,
		PoolSize:     config.Conf.Redis.PoolSize,
		MinIdleConns: config.Conf.Redis.MinIdleConn,
	})
	_, err := c.Ping(context.TODO()).Result()
	if err != nil {
		panic(err)
	}
	redisClient = &Redis{
		client: c,
	}
}

func GetClient() *Redis {
	if nil == redisClient {
		panic("Please initialize the Redis client first!")
	}
	return redisClient
}

func Close() {
	if nil != redisClient {
		_ = redisClient.client.Close()
	}
}
func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}
func (r *Redis) Set(ctx context.Context, key string, value interface{}, expiration int64) error {
	return r.client.Set(ctx, key, value, time.Duration(expiration)).Err()
}
func (r *Redis) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}
func (r *Redis) Hset(ctx context.Context, key string, field string, value interface{}) error {
	return redisClient.client.HSet(ctx, key, field, value).Err()
}
func (r *Redis) Hget(ctx context.Context, key string, field string) (string, error) {
	return redisClient.client.HGet(ctx, key, field).Result()
}
func (r *Redis) Hdel(ctx context.Context, key string, field string) error {
	return redisClient.client.HDel(ctx, key, field).Err()
}
func (r *Redis) Hincrby(ctx context.Context, key string, field string, incr int64) error {
	return redisClient.client.HIncrBy(ctx, key, field, incr).Err()
}
