package redis

import (
	"github.com/go-redis/redis"
	"log"
)

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "GTO4mjZQXZkWYgspMWHHgla0Lf5yNew8zlgRyq",
		DB:       0,
		PoolSize: 10,
	})
	// 通过Ping测试连接
	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}

func GetRedisClient() *redis.Client {
	return redisClient
}
