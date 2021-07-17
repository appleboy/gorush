package simple

import (
	"runtime"
	"testing"

	"github.com/appleboy/gorush/gorush"

	"github.com/stretchr/testify/assert"
)

func TestQueueUsage(t *testing.T) {
	w := NewWorker()
	assert.Equal(t, runtime.NumCPU()<<1, w.Capacity())
	assert.Equal(t, 0, w.Usage())

	w.Queue(gorush.PushNotification{})
	assert.Equal(t, 1, w.Usage())
}

func TestMaxCapacity(t *testing.T) {
	w := NewWorker(WithQueueNum(2))
	assert.Equal(t, 2, w.Capacity())
	assert.Equal(t, 0, w.Usage())

	assert.NoError(t, w.Queue(gorush.PushNotification{}))
	assert.Equal(t, 1, w.Usage())
	assert.NoError(t, w.Queue(gorush.PushNotification{}))
	assert.Equal(t, 2, w.Usage())
	assert.Error(t, w.Queue(gorush.PushNotification{}))
	assert.Equal(t, 2, w.Usage())

	err := w.Queue(gorush.PushNotification{})
	assert.Equal(t, errMaxCapacity, err)
}
