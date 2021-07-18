package simple

import (
	"log"
	"runtime"
	"testing"
	"time"

	"github.com/appleboy/gorush/gorush"
	"github.com/appleboy/gorush/logx"
	"github.com/appleboy/gorush/queue"

	"github.com/stretchr/testify/assert"
)

type mockMessage struct {
	msg string
}

func (m mockMessage) Bytes() []byte {
	return []byte(m.msg)
}

func TestMain(m *testing.M) {
	if err := logx.InitLog(
		"debug",
		"stdout",
		"debug",
		"stdout",
	); err != nil {
		log.Fatalf("Can't load log module, error: %v", err)
	}

	m.Run()
}

func TestQueueUsage(t *testing.T) {
	w := NewWorker()
	assert.Equal(t, runtime.NumCPU()<<1, w.Capacity())
	assert.Equal(t, 0, w.Usage())

	w.Queue(&gorush.PushNotification{})
	assert.Equal(t, 1, w.Usage())
}

func TestMaxCapacity(t *testing.T) {
	w := NewWorker(WithQueueNum(2))
	assert.Equal(t, 2, w.Capacity())
	assert.Equal(t, 0, w.Usage())

	assert.NoError(t, w.Queue(&gorush.PushNotification{}))
	assert.Equal(t, 1, w.Usage())
	assert.NoError(t, w.Queue(&gorush.PushNotification{}))
	assert.Equal(t, 2, w.Usage())
	assert.Error(t, w.Queue(&gorush.PushNotification{}))
	assert.Equal(t, 2, w.Usage())

	err := w.Queue(&gorush.PushNotification{})
	assert.Equal(t, errMaxCapacity, err)
}

func TestCustomFuncAndWait(t *testing.T) {
	m := mockMessage{
		msg: "foo",
	}
	w := NewWorker(
		WithRunFunc(func(msg queue.QueuedMessage) error {
			logx.LogAccess.Infof("get message: %s", msg.Bytes())
			time.Sleep(500 * time.Millisecond)
			return nil
		}),
	)
	q := queue.NewQueue(w, 2)
	q.Start()
	time.Sleep(100 * time.Millisecond)
	q.Queue(m)
	q.Queue(m)
	q.Queue(m)
	q.Queue(m)
	time.Sleep(600 * time.Millisecond)
	q.Shutdown()
	q.Wait()
	// you will see the execute time > 1000ms
}
