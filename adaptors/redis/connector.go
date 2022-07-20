package redis

import (
	"context"
	"github.com/erfidev/url_shortener/env"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
)

func Connect(con *env.RedisEnv) *RDB {
	dbToInt, err := strconv.Atoi(con.DB)
	if err != nil {
		log.Fatalf("Error on converting redis db to int: %s", err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     con.Host + ":" + con.Port,
		Password: con.Password,
		DB:       dbToInt,
	})

	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("redis connection error: %s", err)
	}

	log.Println("redis connection success")

	return &RDB{
		DB: client,
	}
}
