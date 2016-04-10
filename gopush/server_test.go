package gopush

import (
	"github.com/appleboy/gofight"
	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"runtime"
	"testing"
	"time"
)

var go_version = runtime.Version()

func initTest() {
	PushConf = BuildDefaultPushConf()
	PushConf.Core.Mode = "test"
}

func TestPrintGoPushVersion(t *testing.T) {
	PrintGoPushVersion()
}

func TestRunNormalServer(t *testing.T) {
	initTest()

	gin.SetMode(gin.TestMode)

	go func() {
		assert.NoError(t, RunHTTPServer())
	}()
	// have to wait for the goroutine to start and run the server
	// otherwise the main thread will complete
	time.Sleep(5 * time.Millisecond)

	assert.Error(t, RunHTTPServer())
	gofight.TestRequest(t, "http://localhost:8088/api/status")
}

func TestRunTLSServer(t *testing.T) {
	initTest()

	PushConf.Core.SSL = true
	PushConf.Core.Port = "8087"
	PushConf.Core.CertPath = "../certificate/localhost.cert"
	PushConf.Core.KeyPath = "../certificate/localhost.key"

	go func() {
		assert.NoError(t, RunHTTPServer())
	}()
	// have to wait for the goroutine to start and run the server
	// otherwise the main thread will complete
	time.Sleep(5 * time.Millisecond)

	assert.Error(t, RunHTTPServer())
}

func TestRootHandler(t *testing.T) {
	initTest()

	r := gofight.New()

	// log for json
	PushConf.Log.Format = "json"

	r.GET("/").
		Run(GetMainEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			data := []byte(r.Body.String())

			value, _ := jsonparser.GetString(data, "text")

			assert.Equal(t, "Welcome to notification server.", value)
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestAPIStatusHandler(t *testing.T) {
	initTest()

	r := gofight.New()

	r.GET("/api/status").
		Run(GetMainEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			data := []byte(r.Body.String())

			value, _ := jsonparser.GetString(data, "go_version")

			assert.Equal(t, go_version, value)
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestMissingNotificationsParameter(t *testing.T) {
	initTest()

	r := gofight.New()

	// missing notifications parameter.
	r.POST("/api/push").
		Run(GetMainEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

			assert.Equal(t, http.StatusBadRequest, r.Code)
		})
}

func TestEmptyNotifications(t *testing.T) {
	initTest()

	r := gofight.New()

	// notifications is empty.
	r.POST("/api/push").
		SetJSON(gofight.D{
			"notifications": []PushNotification{},
		}).
		Run(GetMainEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

			assert.Equal(t, http.StatusBadRequest, r.Code)
		})
}

func TestSuccessPushHandler(t *testing.T) {
	initTest()

	PushConf.Android.Enabled = true
	PushConf.Android.ApiKey = os.Getenv("ANDROID_API_KEY")

	android_token := os.Getenv("ANDROID_TEST_TOKEN")

	r := gofight.New()

	r.POST("/api/push").
		SetJSON(gofight.D{
			"notifications": []gofight.D{
				gofight.D{
					"tokens":   []string{android_token, "bbbbb"},
					"platform": 2,
					"message":  "Welcome",
				},
			},
		}).
		Run(GetMainEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

			assert.Equal(t, http.StatusOK, r.Code)
		})
}
