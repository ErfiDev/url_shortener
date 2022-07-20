package redis

import (
	"context"
	"time"
)

func (rdb RDB) Set(ctx context.Context, key string, value string, exp time.Duration) error {
	_, err := rdb.DB.Set(ctx, key, value, exp).Result()
	return err
}

func (rdb RDB) Get(ctx context.Context, key string) (string, error) {
	res := rdb.DB.Get(ctx, key)

	return res.Result()
}

func (rdb RDB) Delete(ctx context.Context, key string) error {
	_, err := rdb.DB.Del(ctx, key).Result()
	return err
}

func (rdb RDB) GetKeys(ctx context.Context, pattern string) ([]string, error) {
	res := rdb.DB.Keys(ctx, pattern)

	return res.Result()
}

func (rdb RDB) GetM(ctx context.Context, keys []string) ([]interface{}, error) {
	res := rdb.DB.MGet(ctx, keys...)

	return res.Result()
}
