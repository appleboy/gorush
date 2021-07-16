package queue

import (
	"runtime"

	"github.com/appleboy/gorush/logx"
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
func NewQueue(w Worker) *Queue {
	q := &Queue{
		workerCount:  runtime.NumCPU(),
		queueCount:   runtime.NumCPU() << 1,
		routineGroup: newRoutineGroup(),
		quit:         make(chan struct{}),
		worker:       w,
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
				logx.LogAccess.Info("started the worker num ", num)
				q.worker.Run(q.quit)
				logx.LogAccess.Info("closed the worker num ", num)
			})
		}(i)
	}
}
