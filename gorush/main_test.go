package gorush

import (
	"context"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/appleboy/gorush/config"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	PushConf, _ = config.LoadConf("")
	if err := InitLog(); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(int(PushConf.Core.WorkerNum))
	InitWorkers(ctx, wg, PushConf.Core.WorkerNum, PushConf.Core.QueueNum)

	if err := InitAppStatus(); err != nil {
		log.Fatal(err)
	}

	defer func() {
		close(QueueNotification)
		cancel()
		wg.Wait()
		Stats.Close()
		time.Sleep(1 * time.Second)
		goleak.VerifyTestMain(m)
	}()
}
