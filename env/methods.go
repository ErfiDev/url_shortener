package env

import "os"

func LoadRedisEnv() RedisEnv {
	return RedisEnv{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       os.Getenv("REDIS_DB"),
	}
}

func LoadAppEnv() AppEnv {
	return AppEnv{
		Port:   os.Getenv("PORT"),
		Domain: os.Getenv("DOMAIN"),
	}
}
