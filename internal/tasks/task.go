package tasks

import (
	"context"
	"time"
)

type Task struct {
	UUID    string
	Context context.Context
	ETA     *time.Time
	Image   string
}