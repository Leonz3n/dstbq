package broker

type Broker interface {
	// GetConfig returns config
	GetConfig()

	// StartConsuming enters a loop and waits for incoming messages
	StartConsuming()

	// StopConsuming quits the loop
	StopConsuming()

	// Publish places a new message on the default queue
	Publish()

	// GetPendingTasks returns a slice of task signatures waiting in the queue
	GetPendingTasks()

	// GetDelayedTasks returns a slice of task signatures that are scheduled, but not yet in the queue
	GetDelayedTasks()
}
