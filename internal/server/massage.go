package server

import (
	brokersiface "github.com/Leonz3n/dstbq/internal/brokers/iface"
	backendsiface "github.com/Leonz3n/dstbq/internal/backends/iface"
)

type Massage struct {
	broker brokersiface.Broker
	backend backendsiface.Backend
}

// NewMassageServer new a Massge server.
func NewMassageServer(brokerServer brokersiface.Broker, backendServer backendsiface.Backend) *Massage {
	return &Massage{
		broker:  brokerServer,
		backend: backendServer,
	}
}