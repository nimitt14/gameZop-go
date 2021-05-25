package database

import (
    "context"
    "github.com/go-redis/redis/v8"
)

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
        Addr:     "redis-19115.c14.us-east-1-3.ec2.cloud.redislabs.com:19115",
        Password: "fJPvOIHVwxArsbwyHDfND8x9HuAxuK9d", 
        DB:       0, 
	})
	// blocking connect
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return rdb
}

