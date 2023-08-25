package redis

import (
	"context"

	"github.com/Leonz3n/kulery/log"

	"github.com/go-redis/redis/v8"
)

type BackendGR struct {
	logger log.Logger
	client redis.UniversalClient
}

// NewBackendGR new a redis backend.
func NewBackendGR(logger log.Logger, client redis.UniversalClient) *BackendGR {
	return &BackendGR{logger: logger, client: client}
}

// Ping redis connect.
func (b *BackendGR) Ping() error {
	return b.client.Ping(context.Background()).Err()
}

// Close redis connect.
func (b *BackendGR) Close() error {
	return b.client.Close()
}
