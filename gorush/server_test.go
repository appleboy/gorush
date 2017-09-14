package gorush

import (
	"net/http"
	"os"
	"runtime"
	"testing"
	"time"

	"github.com/axiomzen/gorush/config"
	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gopkg.in/appleboy/gofight.v2"
)

var goVersion = runtime.Version()

func initTest() {
	PushConf = config.BuildDefaultPushConf()
	PushConf.Core.Mode = "test"
}

func TestPrintGoRushVersion(t *testing.T) {
	SetVersion("3.0.0")
	ver := GetVersion()
	PrintGoRushVersion()

	assert.Equal(t, "3.0.0", ver)
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
	gofight.TestRequest(t, "http://localhost:8088/api/stat/go")
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

	gofight.TestRequest(t, "https://localhost:8087/api/stat/go")
}

func TestRunAutoTLSServer(t *testing.T) {
	initTest()
	PushConf.Core.AutoTLS.Enabled = true
	go func() {
		assert.NoError(t, RunHTTPServer())
	}()
	// have to wait for the goroutine to start and run the server
	// otherwise the main thread will complete
	time.Sleep(5 * time.Millisecond)
}

func TestLoadTLSCertError(t *testing.T) {
	initTest()

	PushConf.Core.SSL = true
	PushConf.Core.Port = "8087"
	PushConf.Core.CertPath = "../config/config.yml"
	PushConf.Core.KeyPath = "../config/config.yml"

	assert.Error(t, RunHTTPServer())
}

func TestRootHandler(t *testing.T) {
	initTest()

	r := gofight.New()

	// log for json
	PushConf.Log.Format = "json"

	r.GET("/").
		Run(routerEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			data := []byte(r.Body.String())

			value, _ := jsonparser.GetString(data, "text")

			assert.Equal(t, "Welcome to notification server.", value)
			assert.Equal(t, http.StatusOK, r.Code)
			assert.Equal(t, "application/json; charset=utf-8", r.HeaderMap.Get("Content-Type"))
		})
}

func TestAPIStatusGoHandler(t *testing.T) {
	initTest()

	r := gofight.New()

	r.GET("/api/stat/go").
		Run(routerEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			data := []byte(r.Body.String())

			value, _ := jsonparser.GetString(data, "go_version")

			assert.Equal(t, goVersion, value)
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestAPIStatusAppHandler(t *testing.T) {
	initTest()

	r := gofight.New()

	appVersion := "v1.0.0"
	SetVersion(appVersion)

	r.GET("/api/stat/app").
		Run(routerEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			data := []byte(r.Body.String())

			value, _ := jsonparser.GetString(data, "version")

			assert.Equal(t, appVersion, value)
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestAPIConfigHandler(t *testing.T) {
	initTest()

	r := gofight.New()

	r.GET("/api/config").
		Run(routerEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusCreated, r.Code)
		})
}

func TestMissingNotificationsParameter(t *testing.T) {
	initTest()

	r := gofight.New()

	// missing notifications parameter.
	r.POST("/api/push").
		Run(routerEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

			assert.Equal(t, http.StatusBadRequest, r.Code)
			assert.Equal(t, "application/json; charset=utf-8", r.HeaderMap.Get("Content-Type"))
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
		Run(routerEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

			assert.Equal(t, http.StatusBadRequest, r.Code)
		})
}

func TestOutOfRangeMaxNotifications(t *testing.T) {
	initTest()

	PushConf.Core.MaxNotification = int64(1)

	r := gofight.New()

	// notifications is empty.
	r.POST("/api/push").
		SetJSON(gofight.D{
			"notifications": []gofight.D{
				{
					"tokens":   []string{"aaaaa", "bbbbb"},
					"platform": PlatFormAndroid,
					"message":  "Welcome",
				},
				{
					"tokens":   []string{"aaaaa", "bbbbb"},
					"platform": PlatFormAndroid,
					"message":  "Welcome",
				},
			},
		}).
		Run(routerEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

			assert.Equal(t, http.StatusBadRequest, r.Code)
		})
}

func TestSuccessPushHandler(t *testing.T) {
	initTest()

	PushConf.Android.Enabled = true
	PushConf.Android.APIKey = os.Getenv("ANDROID_API_KEY")

	androidToken := os.Getenv("ANDROID_TEST_TOKEN")

	r := gofight.New()

	r.POST("/api/push").
		SetJSON(gofight.D{
			"notifications": []gofight.D{
				{
					"tokens":   []string{androidToken, "bbbbb"},
					"platform": PlatFormAndroid,
					"message":  "Welcome",
				},
			},
		}).
		Run(routerEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestSysStatsHandler(t *testing.T) {
	initTest()

	r := gofight.New()

	r.GET("/sys/stats").
		Run(routerEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestMetricsHandler(t *testing.T) {
	initTest()

	r := gofight.New()

	r.GET("/metrics").
		Run(routerEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestDisabledHTTPServer(t *testing.T) {
	initTest()
	PushConf.Core.Enabled = false
	err := RunHTTPServer()
	PushConf.Core.Enabled = true

	assert.Nil(t, err)
}
