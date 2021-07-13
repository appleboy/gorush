package gorush

import (
	"context"
	"log"
	"sync"
	"testing"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/logx"
)

func TestMain(m *testing.M) {
	PushConf, _ = config.LoadConf("")
	if err := logx.InitLog(
		PushConf.Log.AccessLevel,
		PushConf.Log.AccessLog,
		PushConf.Log.ErrorLevel,
		PushConf.Log.ErrorLog,
	); err != nil {
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
	}()

	m.Run()
}
