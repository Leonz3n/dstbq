package task

import (
	"errors"
	"fmt"
)

// State denotes the state of a task.
type State int

const (
	StatePending    State = iota + 1 // initial state of a task
	StateProcessing                  // when the worker starts processing the task
	StateRetry                       // when failed task has been scheduled for retry
	StateCompleted                   // when the task is completed
)

func (s State) String() string {
	switch s {
	case StatePending:
		return "pending"
	case StateProcessing:
		return "processing"
	case StateRetry:
		return "retry"
	case StateCompleted:
		return "completed"
	}
	panic(fmt.Sprintf("internal error: unknown task state %d", s))
}

func StateFromString(s string) (State, error) {
	switch s {
	case "pending":
		return StatePending, nil
	case "processing":
		return StateProcessing, nil
	case "retry":
		return StateRetry, nil
	case "completed":
		return StateCompleted, nil
	}
	return 0, errors.New(fmt.Sprintf("%q is not supported task state", s))
}
