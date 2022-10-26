package redis

import "github.com/go-redis/redis/v8"

type Backend struct {
	rclient redis.UniversalClient
}

// New a redis backend.
func New(addrs []string) *Backend {
	ropt := &redis.UniversalOptions{
		Addrs:    addrs,
		DB:       0,
		Password: "",
	}

	return &Backend{
		rclient: redis.NewUniversalClient(ropt),
	}
}
