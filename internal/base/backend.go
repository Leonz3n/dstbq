package base

type Backend interface {
	Ping() error
	Close() error
}
