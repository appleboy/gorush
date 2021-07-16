package queue

// Worker interface
type Worker interface {
	Run(chan struct{})
	Stop()
	Enqueue(job interface{}) error
	Capacity() int
	Usage() int
}
