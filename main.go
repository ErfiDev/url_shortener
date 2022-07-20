package main

import (
	"github.com/erfidev/url_shortener/adaptors/redis"
	"github.com/erfidev/url_shortener/delivery"
	"github.com/erfidev/url_shortener/env"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	// loading required environments
	redisEnv := env.LoadRedisEnv()
	appEnv := env.LoadAppEnv()

	// loading required app dependencies
	rdb := redis.Connect(&redisEnv)

	delivery.Setup(rdb, &appEnv)
}
