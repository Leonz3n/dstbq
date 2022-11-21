package redis

import (
	"context"
	"github.com/Leonz3n/k8s-job-massage/internal/base"
	"github.com/go-redis/redis/v8"
)

type BackendGR struct {
	logger base.Logger
	client redis.UniversalClient
}

// NewBackendGR new a redis backend.
func NewBackendGR(logger base.Logger, client redis.UniversalClient) *BackendGR {
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
