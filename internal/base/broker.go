package base

type Broker interface {
	Ping() error
	Close() error
}
