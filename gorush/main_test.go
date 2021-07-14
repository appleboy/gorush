package gorush

import (
	"log"
	"testing"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/logx"
	"github.com/appleboy/gorush/status"
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

	if err := status.InitAppStatus(PushConf); err != nil {
		log.Fatal(err)
	}

	m.Run()
}
