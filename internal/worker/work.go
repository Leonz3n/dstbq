package worker

import (
	"github.com/Leonz3n/dstbq/internal/tasks"
)

type Worker struct {
	Queue string
}

// NewWorker new a Wokrer instance.
func NewWorker(queue string) *Worker {
	return &Worker{
		Queue: queue,
	}
}

func (worker *Worker) Launch() error {
	return nil
}

func (worker *Worker) Process(task *tasks.Task) error {
	return nil
}