package redis

import "github.com/go-redis/redis/v8"

type Broker struct {
	rclient redis.UniversalClient
}

// New a redis broker.
func New(addrs []string) *Broker {
	ropt := &redis.UniversalOptions{
		Addrs:    addrs,
		DB:       0,
		Password: "",
	}

	return &Broker{
		rclient: redis.NewUniversalClient(ropt),
	}
}