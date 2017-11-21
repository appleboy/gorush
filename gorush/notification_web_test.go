package gorush

import (
	"log"
	"os"
	"testing"

	"github.com/appleboy/gorush/config"
	"github.com/stretchr/testify/assert"
)

func init() {
	PushConf, _ = config.LoadConf("")
	if err := InitLog(); err != nil {
		log.Fatal(err)
	}

	InitWorkers(PushConf.Core.WorkerNum, PushConf.Core.QueueNum)

	if err := InitAppStatus(); err != nil {
		log.Fatal(err)
	}
}

func TestMissingWebAPIKey(t *testing.T) {
	PushConf, _ = config.LoadConf("")

	PushConf.Web.Enabled = true
	PushConf.Web.APIKey = ""

	err := CheckPushConf()

	assert.Error(t, err)
	assert.Equal(t, "Missing GCM API Key for Chrome", err.Error())
}

func TestPushToWebWrongSubscription(t *testing.T) {
	PushConf, _ = config.LoadConf("")

	PushConf.Web.Enabled = true
	PushConf.Web.APIKey = os.Getenv("ANDROID_API_KEY")

	req := PushNotification{
		Subscriptions: []Subscription{{"aaaaaa", "bbbbbb", "cccccc"}},
		Platform:      PlatformWeb,
		Message:       "Welcome",
	}

	// Web Success count: 0, Failure count: 1
	isError := PushToWeb(req)
	assert.True(t, isError)
}
