package worker

import (
	"sync"
	"time"

	"github.com/Leonz3n/k8s-job-massage/internal/base"
	"k8s.io/client-go/kubernetes"
)

type Worker struct {
	logger     base.Logger
	kclientset *kubernetes.Clientset
	broker     base.Broker
	backend    base.Backend

	shutdownTimeout time.Duration

	once sync.Once
	// sema is a counting semaphore to ensure the number of active workers
	// does not exceed the limit.
	sema chan struct{}
	quit chan struct{}
	done chan struct{}
}

type Params struct {
	Logger      base.Logger
	Broker      base.Broker
	Backend     base.Backend
	KClientSet  *kubernetes.Clientset
	Concurrency int
}

// NewWorker new a worker.
func NewWorker(params *Params) *Worker {
	return &Worker{
		logger:     params.Logger,
		kclientset: params.KClientSet,
		broker:     params.Broker,
		backend:    params.Backend,
		sema:       make(chan struct{}, params.Concurrency),
		quit:       make(chan struct{}),
		done:       make(chan struct{}),
	}
}

func (w *Worker) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-w.done:
				w.logger.Debug("Processor done")
				return
			default:
				w.Exec()
			}
		}
	}()
}

func (w *Worker) Exec() {
	select {
	case <-w.quit:
		return
	case w.sema <- struct{}{}:
		go func() {
			defer func() {
				<-w.sema // release token
			}()

			// TODO process task
		}()
	}
}

func (w *Worker) Stop() {
	w.once.Do(func() {
		w.logger.Debug("Processor shutting down...")
		// Unblock if processor is waiting for sema token.
		close(w.quit)
		// Signal the processor goroutine to stop processing tasks
		// from the queue.
		w.done <- struct{}{}
	})
}

func (w *Worker) Shutdown() {
	w.Stop()

	w.logger.Info("Waiting for all workers to finish...")
	// block until all workers have released the token
	for i := 0; i < cap(w.sema); i++ {
		w.sema <- struct{}{}
	}
	w.logger.Info("All workers have finished")
}
