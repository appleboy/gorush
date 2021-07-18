package nsq

import (
	"log"
	"testing"
	"time"

	"github.com/appleboy/gorush/logx"
	"github.com/appleboy/gorush/queue"
	"github.com/nsqio/go-nsq"
)

var host = "nsq"

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
		WithAddr(host+":4150"),
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
		WithAddr(host+":4150"),
		WithTopic("test"),
		WithRunFunc(func(msg *nsq.Message) error {
			logx.LogAccess.Infof("get message: %s", msg.Body)
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
