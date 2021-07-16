package simple

import (
	"errors"

	"github.com/appleboy/gorush/gorush"
	"github.com/appleboy/gorush/queue"
)

var _ queue.Worker = (*Worker)(nil)

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
		return errors.New("max capacity reached")
	}
}

// NewWorker for struct
func NewWorker(num int) *Worker {
	return &Worker{
		queueNotification: make(chan gorush.PushNotification, num),
	}
}
