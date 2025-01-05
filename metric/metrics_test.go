package metric

import (
	"context"
	"testing"
	"time"

	"github.com/golang-queue/queue"
	"github.com/stretchr/testify/assert"
)

var noTask = func(ctx context.Context) error { return nil }

func TestNewMetrics(t *testing.T) {
	q := queue.NewPool(10)
	assert.NoError(t, q.QueueTask(noTask))
	assert.NoError(t, q.QueueTask(noTask))
	time.Sleep(10 * time.Millisecond)
	defer q.Release()
	m := NewMetrics(q)
	assert.Equal(t, 2, m.q.SubmittedTasks())
	assert.Equal(t, 2, m.q.SuccessTasks())
}
