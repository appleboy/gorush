package queue

// Worker interface
type Worker interface {
	Run(chan struct{}) error
	Shutdown() error
	Queue(job interface{}) error
	Capacity() int
	Usage() int
}
