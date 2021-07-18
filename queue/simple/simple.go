package simple

import (
	"errors"
	"runtime"

	"github.com/appleboy/gorush/gorush"
	"github.com/appleboy/gorush/queue"
)

var _ queue.Worker = (*Worker)(nil)

// Option for queue system
type Option func(*Worker)

var errMaxCapacity = errors.New("max capacity reached")

// Worker for simple queue using channel
type Worker struct {
	QueueNotification chan queue.QueuedMessage
	runFunc           func(*Worker) error
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
	return s.runFunc(s)
}

// Shutdown worker
func (s *Worker) Shutdown() error {
	close(s.QueueNotification)
	return nil
}

// Capacity for channel
func (s *Worker) Capacity() int {
	return cap(s.QueueNotification)
}

// Usage for count of channel usage
func (s *Worker) Usage() int {
	return len(s.QueueNotification)
}

// Queue send notification to queue
func (s *Worker) Queue(job queue.QueuedMessage) error {
	select {
	case s.QueueNotification <- job:
		return nil
	default:
		return errMaxCapacity
	}
}

// WithQueueNum setup the capcity of queue
func WithQueueNum(num int) Option {
	return func(w *Worker) {
		w.QueueNotification = make(chan queue.QueuedMessage, num)
	}
}

// WithRunFunc setup the run func of queue
func WithRunFunc(fn func(w *Worker) error) Option {
	return func(w *Worker) {
		w.runFunc = fn
	}
}

// NewWorker for struc
func NewWorker(opts ...Option) *Worker {
	w := &Worker{
		QueueNotification: make(chan queue.QueuedMessage, runtime.NumCPU()<<1),
		runFunc: func(w *Worker) error {
			for notification := range w.QueueNotification {
				gorush.SendNotification(notification)
			}
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
