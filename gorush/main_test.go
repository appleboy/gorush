package gorush

import (
	"log"
	"testing"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/status"
)

func TestMain(m *testing.M) {
	cfg, _ := config.LoadConf()
	if err := status.InitAppStatus(cfg); err != nil {
		log.Fatal(err)
	}

	m.Run()
}
