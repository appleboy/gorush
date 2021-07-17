package queue

// Worker interface
type Worker interface {
	BeforeRun() error
	Run(chan struct{}) error
	AfterRun() error

	Shutdown() error
	Queue(job QueuedMessage) error
	Capacity() int
	Usage() int
}

// QueuedMessage ...
type QueuedMessage interface {
	Bytes() []byte
}
