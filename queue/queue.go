package queue

import (
	"runtime"
)

type (
	// A Queue is a message queue.
	Queue struct {
		logger       Logger
		workerCount  int
		routineGroup *routineGroup
		quit         chan struct{}
		worker       Worker
	}
)

// NewQueue returns a Queue.
func NewQueue(w Worker, workerNum int) *Queue {
	q := &Queue{
		workerCount:  runtime.NumCPU(),
		routineGroup: newRoutineGroup(),
		quit:         make(chan struct{}),
		worker:       w,
		logger:       new(defaultLogger),
	}

	if workerNum > 0 {
		q.workerCount = workerNum
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

// Shutdown stops all queues.
func (q *Queue) Shutdown() {
	q.worker.Shutdown()
	close(q.quit)
}

// Wait all process
func (q *Queue) Wait() {
	q.routineGroup.Wait()
}

// Queue to queue all job
func (q *Queue) Queue(job QueuedMessage) error {
	return q.worker.Queue(job)
}

func (q *Queue) work(num int) {
	if err := q.worker.BeforeRun(); err != nil {
		q.logger.Fatal(err)
	}
	q.routineGroup.Run(func() {
		// to handle panic cases from inside the worker
		// in such case, we start a new goroutine
		defer func() {
			if err := recover(); err != nil {
				q.logger.Error(err)
				go q.work(num)
			}
		}()

		q.logger.Info("started the worker num ", num)
		q.worker.Run(q.quit)
		q.logger.Info("closed the worker num ", num)
	})
	if err := q.worker.AfterRun(); err != nil {
		q.logger.Fatal(err)
	}
}

func (q *Queue) startWorker() {
	for i := 0; i < q.workerCount; i++ {
		go q.work(i)
	}
}
