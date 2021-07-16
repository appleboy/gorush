package simple

import (
	"errors"

	"github.com/appleboy/gorush/gorush"
)

// Worker for simple queue using channel
type Worker struct {
	queueNotification chan gorush.PushNotification
}

// Run start the worker
func (s *Worker) Run(_ chan struct{}) {
	for notification := range s.queueNotification {
		gorush.SendNotification(notification)
	}
}

// Stop worker
func (s *Worker) Stop() {
	close(s.queueNotification)
}

// Capacity for channel
func (s *Worker) Capacity() int {
	return cap(s.queueNotification)
}

// Usage for count of channel usage
func (s *Worker) Usage() int {
	return len(s.queueNotification)
}

// Enqueue send notification to queue
func (s *Worker) Enqueue(job interface{}) error {
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
