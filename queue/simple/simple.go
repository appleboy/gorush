package simple

import (
	"errors"
	"runtime"

	"github.com/appleboy/gorush/notify"
	"github.com/appleboy/gorush/queue"
)

var _ queue.Worker = (*Worker)(nil)

// Option for queue system
type Option func(*Worker)

var errMaxCapacity = errors.New("max capacity reached")

// Worker for simple queue using channel
type Worker struct {
	queueNotification chan queue.QueuedMessage
	runFunc           func(queue.QueuedMessage) error
}

// BeforeRun run script before start worker
func (s *Worker) BeforeRun() error {
	return nil
}

// AfterRun run script after start worker
func (s *Worker) AfterRun() error {
	return nil
}

// Run start the worker
func (s *Worker) Run(_ chan struct{}) error {
	for notification := range s.queueNotification {
		s.runFunc(notification)
	}
	return nil
}

// Shutdown worker
func (s *Worker) Shutdown() error {
	close(s.queueNotification)
	return nil
}

// Capacity for channel
func (s *Worker) Capacity() int {
	return cap(s.queueNotification)
}

// Usage for count of channel usage
func (s *Worker) Usage() int {
	return len(s.queueNotification)
}

// Queue send notification to queue
func (s *Worker) Queue(job queue.QueuedMessage) error {
	select {
	case s.queueNotification <- job:
		return nil
	default:
		return errMaxCapacity
	}
}

// WithQueueNum setup the capcity of queue
func WithQueueNum(num int) Option {
	return func(w *Worker) {
		w.queueNotification = make(chan queue.QueuedMessage, num)
	}
}

// WithRunFunc setup the run func of queue
func WithRunFunc(fn func(queue.QueuedMessage) error) Option {
	return func(w *Worker) {
		w.runFunc = fn
	}
}

// NewWorker for struc
func NewWorker(opts ...Option) *Worker {
	w := &Worker{
		queueNotification: make(chan queue.QueuedMessage, runtime.NumCPU()<<1),
		runFunc: func(msg queue.QueuedMessage) error {
			notify.SendNotification(msg)
			return nil
		},
	}

	// Loop through each option
	for _, opt := range opts {
		// Call the option giving the instantiated
		opt(w)
	}

	return w
}
