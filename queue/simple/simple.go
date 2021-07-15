package simple

import (
	"errors"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/gorush"
)

// Worker for simple queue using channel
type Worker struct {
	cfg               config.ConfYaml
	queueNotification chan gorush.PushNotification
}

// Run start the worker
func (s *Worker) Run(_ chan struct{}) {
	for notification := range s.queueNotification {
		gorush.SendNotification(s.cfg, notification)
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
func NewWorker(cfg config.ConfYaml) *Worker {
	return &Worker{
		cfg:               cfg,
		queueNotification: make(chan gorush.PushNotification, cfg.Core.QueueNum),
	}
}
