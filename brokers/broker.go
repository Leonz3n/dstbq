package brokers

type Broker interface {
	Ping() error
}
