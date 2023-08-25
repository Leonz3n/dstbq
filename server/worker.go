package server

import (
	"sync"

	"github.com/Leonz3n/kulery/brokers"
	"github.com/Leonz3n/kulery/log"
)

type Worker struct {
	broker brokers.Broker

	logger log.Logger

	// channel to communicate back to the long running "processor" goroutine.
	done chan struct{}

	// quit channel is closed when the shutdown of the "processor" goroutine starts.
	quit chan struct{}

	// abort channel communicates to the in-flight worker goroutines to stop.
	abort chan struct{}

	// sema is a counting semaphore to ensure the number of active workers
	// does not exceed the limit.
	sema chan struct{}
}

func NewWorker() *Worker {
	return &Worker{
		done:  make(chan struct{}),
		quit:  make(chan struct{}),
		abort: make(chan struct{}),
		sema:  make(chan struct{}, 3),
	}
}

func (w *Worker) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-w.done:
				w.logger.Info("worker done")
			default:
				w.exec()
			}
		}
	}()
}

func (w *Worker) Stop(wg *sync.WaitGroup) {

}

func (w *Worker) Shutdown() {
	for i := 0; i < cap(w.sema); i++ {
		w.sema <- struct{}{}
	}
}

func (w *Worker) exec() {
	select {
	case <-w.quit:
		return
	case w.sema <- struct{}{}:

	}
}
