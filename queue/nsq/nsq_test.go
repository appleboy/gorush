package nsq

import (
	"log"
	"testing"
	"time"

	"github.com/appleboy/gorush/logx"
	"github.com/appleboy/gorush/queue"
)

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
		WithAddr("nsq:4150"),
		WithTopic("test"),
	)
	q := queue.NewQueue(w, 2)
	q.Start()
	time.Sleep(1 * time.Second)
	q.Shutdown()
}
