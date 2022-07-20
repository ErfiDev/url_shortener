package contract

import (
	"context"
	"time"
)

type DbContract interface {
	Set(ctx context.Context, key string, value string, exp time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
	GetKeys(ctx context.Context, pattern string) ([]string, error)
	GetM(ctx context.Context, keys []string) ([]interface{}, error)
}
