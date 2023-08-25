package redis

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Leonz3n/kulery/log"

	"github.com/go-redis/redis/v8"
)

// DefaultQueueName is the queue name used if none are specified by user.
const DefaultQueueName = "default"

// Global Redis keys.
const (
	AllSchedulers = "massage:schedulers" // ZSET
	AllQueues     = "massage:queues"     // SET
)

type BrokerGR struct {
	logger  log.Logger
	rclient redis.UniversalClient
}

// NewBrokerGR new a redis broker
func NewBrokerGR(logger log.Logger, client redis.UniversalClient) *BrokerGR {
	return &BrokerGR{rclient: client, logger: logger}
}

// Ping redis connect.
func (b *BrokerGR) Ping() error {
	return b.rclient.Ping(context.Background()).Err()
}

// Close redis connect.
func (b *BrokerGR) Close() error {
	return b.rclient.Close()
}

// Consume a message from redis.
func (b *BrokerGR) Consume() (*task.Signature, error) {
	signature := new(task.Signature)
	items, err := b.rclient.BLPop(context.Background(), time.Duration(1000)*time.Millisecond, QueueKeyPrefix("default")).Result()
	if err != nil {
		return nil, err
	}

	// items[0] - the name of the key where an element was popped
	// items[1] - the value of the popped element
	if len(items) != 2 {
		return nil, redis.Nil
	}

	decoder := json.NewDecoder(bytes.NewReader([]byte(items[1])))
	decoder.UseNumber()
	if err = decoder.Decode(signature); err != nil {
		return nil, err
	}

	return signature, nil
}

func (b *BrokerGR) Publish(signature *task.Signature) error {
	msg, err := json.Marshal(signature)
	if err != nil {
		return fmt.Errorf("JSON marshal error: %s", err)
	}

	err = b.rclient.RPush(context.Background(), QueueKeyPrefix("default"), msg).Err()
	return err
}

// QueueKeyPrefix returns a prefix for all keys in the given queue.
func QueueKeyPrefix(qname string) string {
	return fmt.Sprintf("k8s-job-massage:{%s}:", qname)
}

// PendingKey returns a redis key for the given queue name.
func PendingKey(qname string) string {
	return fmt.Sprintf("%spending", QueueKeyPrefix(qname))
}
