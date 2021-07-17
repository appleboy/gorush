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
	queueNotification chan gorush.PushNotification
}

// Run start the worker
func (s *Worker) Run(_ chan struct{}) error {
	for notification := range s.queueNotification {
		gorush.SendNotification(notification)
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
func (s *Worker) Queue(job interface{}) error {
	select {
	case s.queueNotification <- job.(gorush.PushNotification):
		return nil
	default:
		return errMaxCapacity
	}
}

// WithQueueNum setup the capcity of queue
func WithQueueNum(num int) Option {
	return func(w *Worker) {
		w.queueNotification = make(chan gorush.PushNotification, num)
	}
}

// NewWorker for struc
func NewWorker(opts ...Option) *Worker {
	w := &Worker{
		queueNotification: make(chan gorush.PushNotification, runtime.NumCPU()<<1),
	}

	// Loop through each option
	for _, opt := range opts {
		// Call the option giving the instantiated
		opt(w)
	}

	return w
}
