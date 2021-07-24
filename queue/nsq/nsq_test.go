package nsq

import (
	"testing"
	"time"

	"github.com/appleboy/gorush/logx"
	"github.com/appleboy/gorush/queue"

	"github.com/nsqio/go-nsq"
	"github.com/stretchr/testify/assert"
)

var host = "nsq"

type mockMessage struct {
	msg string
}

func (m mockMessage) Bytes() []byte {
	return []byte(m.msg)
}

func TestMain(m *testing.M) {
	m.Run()
}

func TestShutdown(t *testing.T) {
	w := NewWorker(
		WithAddr(host+":4150"),
		WithTopic("test"),
	)
	q, err := queue.NewQueue(
		queue.WithWorker(w),
		queue.WithWorkerCount(2),
	)
	assert.NoError(t, err)
	q.Start()
	time.Sleep(1 * time.Second)
	q.Shutdown()
	// check shutdown once
	q.Shutdown()
	q.Wait()
}

func TestCustomFuncAndWait(t *testing.T) {
	m := mockMessage{
		msg: "foo",
	}
	w := NewWorker(
		WithAddr(host+":4150"),
		WithTopic("test"),
		WithMaxInFlight(2),
		WithRunFunc(func(msg *nsq.Message) error {
			logx.LogAccess.Infof("get message: %s", msg.Body)
			time.Sleep(500 * time.Millisecond)
			return nil
		}),
	)
	q, err := queue.NewQueue(
		queue.WithWorker(w),
		queue.WithWorkerCount(2),
	)
	assert.NoError(t, err)
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
