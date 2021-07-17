package queue

// Worker interface
type Worker interface {
	BeforeRun() error
	Run(chan struct{}) error
	AfterRun() error

	Shutdown() error
	Queue(job interface{}) error
	Capacity() int
	Usage() int
}
