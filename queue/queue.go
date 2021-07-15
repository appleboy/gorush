package queue

import (
	"runtime"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/queue/simple"
)

type (
	// A Queue is a message queue.
	Queue struct {
		workerCount  int
		queueCount   int
		routineGroup *routineGroup
		quit         chan struct{}
		worker       Worker
	}
)

// NewQueue returns a Queue.
func NewQueue(cfg config.ConfYaml) *Queue {
	q := &Queue{
		workerCount:  int(cfg.Core.WorkerNum),
		queueCount:   int(cfg.Core.QueueNum),
		routineGroup: newRoutineGroup(),
		quit:         make(chan struct{}),
		worker:       simple.NewWorker(cfg),
	}

	if q.workerCount != 0 {
		q.workerCount = runtime.NumCPU()
	}

	if q.queueCount == 0 {
		q.queueCount = runtime.NumCPU() << 1
	}

	return q
}

// Capacity for queue max size
func (q *Queue) Capacity() int {
	return q.worker.Capacity()
}

// Usage for count of queue usage
func (q *Queue) Usage() int {
	return q.worker.Usage()
}

// Config update current config
func (q *Queue) Config(cfg config.ConfYaml) {
	q.worker.Config(cfg)
}

// Start to enable all worker
func (q *Queue) Start() {
	q.startWorker()
}

// Stop stops q.
func (q *Queue) Stop() {
	q.worker.Stop()
	close(q.quit)
}

// Wait all process
func (q *Queue) Wait() {
	q.routineGroup.Wait()
}

// Enqueue queue all job
func (q *Queue) Enqueue(job interface{}) error {
	return q.worker.Enqueue(job)
}

func (q *Queue) startWorker() {
	for i := 0; i < q.workerCount; i++ {
		go func(num int) {
			q.routineGroup.Run(func() {
				// logx.LogAccess.Info("started the worker num ", num)
				q.worker.Run(q.quit)
				// logx.LogAccess.Info("closed the worker num ", num)
			})
		}(i)
	}
}
