package nsq

import (
	"log"
	"testing"
	"time"

	"github.com/appleboy/gorush/logx"
	"github.com/appleboy/gorush/queue"
	"github.com/nsqio/go-nsq"
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

func TestShutdown(t *testing.T) {
	w := NewWorker(
		WithAddr("127.0.0.1:4150"),
		WithTopic("test"),
	)
	q := queue.NewQueue(w, 2)
	q.Start()
	time.Sleep(1 * time.Second)
	q.Shutdown()
	q.Wait()
}

func TestCustomFunc(t *testing.T) {
	m := mockMessage{
		msg: "foo",
	}
	w := NewWorker(
		WithAddr("127.0.0.1:4150"),
		WithTopic("test"),
		WithRunFunc(func(msg *nsq.Message) error {
			logx.LogAccess.Infof("get message: %s", msg.Body)
			time.Sleep(5 * time.Second)
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
	time.Sleep(6000 * time.Millisecond)
	q.Shutdown()
	q.Wait()
	// you will see the execute time > 10s
}