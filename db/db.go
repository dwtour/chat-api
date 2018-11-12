package db

import "github.com/go-redis/redis"

var (
	Conn *redis.Client
)

func Connect() (string, error) {
	redisdb := redis.NewClient(&redis.Options{
		Addr:     "redis-18104.c9.us-east-1-4.ec2.cloud.redislabs.com:18104",
		Password: "vUH8ylqhWIaXhOTk7QcsLkKikTPjn8Pf",
		DB:       0,  // use default DB
	})
	Conn = redisdb
	pong, err := redisdb.Ping().Result()
	return pong, err
}