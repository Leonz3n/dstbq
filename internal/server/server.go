package server

import (
	"github.com/Leonz3n/k8s-job-massage/internal/base"
	"github.com/Leonz3n/k8s-job-massage/internal/http"
	"github.com/Leonz3n/k8s-job-massage/internal/worker"
	"k8s.io/client-go/kubernetes"
	"sync"
)

// Config specifies the server's processing behavior.
type Config struct {
	// Maximum number of concurrent processing of tasks.
	//
	// If set to a zero or negative value, NewServer will overwrite the value
	// to the number of CPUs usable by the current process.
	Concurrency int

	// Logger specifies the logger used by the server instance.
	//
	// If unset, default logger is used.
	Logger base.Logger

	Broker     base.Broker
	Backend    base.Backend
	KClientSet *kubernetes.Clientset
}

type Server struct {
	logger base.Logger

	worker  *worker.Worker
	httpsvr *http.Server

	// wait group to wait for all goroutines to finish.
	wg sync.WaitGroup
}

// NewServer new a server.
func NewServer(cfg *Config) *Server {
	return &Server{
		logger:  cfg.Logger,
		httpsvr: http.NewServer(cfg.Logger, cfg.Broker),
		worker: worker.NewWorker(&worker.Params{
			Logger:      cfg.Logger,
			Broker:      cfg.Broker,
			KClientSet:  cfg.KClientSet,
			Concurrency: cfg.Concurrency,
		}),
	}
}

func (s *Server) Start() {
	s.httpsvr.Start(&s.wg)
	s.worker.Start(&s.wg)

	s.wg.Wait()
}

// Stop gracefully shuts down the server.
//
// If worker didn't finish processing a task during the timeout, the task will be pushed back to Redis.
func (s *Server) Stop() {

}
