package server

import (
	"sync"
)

type Server struct {
	// wait for all goroutines to finish
	wg sync.WaitGroup

	worker *Worker
}

func NewServer() *Server {
	return &Server{}
}

func (srv *Server) Start() error {

	srv.worker.Start(&srv.wg)
	return nil
}

// Shutdown gracefully shuts down the server.
func (srv *Server) Shutdown() {

}

// Stop signals the server to stop pulling new tasks off queues.
func (srv *Server) Stop() {

}
