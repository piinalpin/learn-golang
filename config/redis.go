package config

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/go-redis/redis/v9"
	"github.com/joho/godotenv"
)

func InitRedis() *redis.Client {
	godotenv.Load()

	var dsn = os.ExpandEnv("${redis.host}:${redis.port}")

	client := redis.NewClient(&redis.Options{
		Addr: dsn,
		DB: 0,
	})

	var ctx = context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Error("Failed to connect to redis. Error: ", err)
		panic(err)
	}

	return client
}