package gopush

import (
	"github.com/appleboy/gofight"
	"github.com/buger/jsonparser"
	"github.com/stretchr/testify/assert"
	"net/http"
	"runtime"
	"testing"
)

var go_version = runtime.Version()

func initTest() {
	PushConf = BuildDefaultPushConf()
	PushConf.Core.Mode = "test"
}

func TestRootHandler(t *testing.T) {
	initTest()

	r := gofight.New()

	r.GET("/").
		Run(GetMainEngine(), func(r gofight.HttpResponse, rq gofight.HttpRequest) {
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
		Run(GetMainEngine(), func(r gofight.HttpResponse, rq gofight.HttpRequest) {
			data := []byte(r.Body.String())

			value, _ := jsonparser.GetString(data, "go_version")

			assert.Equal(t, go_version, value)
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestMissingParameterPushHandler(t *testing.T) {
	initTest()

	r := gofight.New()

	// missing some parameter.
	r.POST("/api/push").
		SetJSON(gofight.D{
			"platform": 1,
		}).
		Run(GetMainEngine(), func(r gofight.HttpResponse, rq gofight.HttpRequest) {

			assert.Equal(t, http.StatusBadRequest, r.Code)
		})
}

func TestDisabledIosPushHandler(t *testing.T) {
	initTest()

	PushConf.Ios.Enabled = false
	InitAPNSClient()

	r := gofight.New()

	r.POST("/api/push").
		SetJSON(gofight.D{
			"tokens":   []string{"11aa01229f15f0f0c52029d8cf8cd0aeaf2365fe4cebc4af26cd6d76b7919ef7"},
			"platform": 1,
			"message":  "Welcome",
		}).
		Run(GetMainEngine(), func(r gofight.HttpResponse, rq gofight.HttpRequest) {

			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestIosPushHandler(t *testing.T) {
	initTest()

	PushConf.Ios.Enabled = true
	PushConf.Ios.PemKeyPath = "../certificate/certificate-valid.pem"
	InitAPNSClient()

	r := gofight.New()

	r.POST("/api/push").
		SetJSON(gofight.D{
			"tokens":   []string{"11aa01229f15f0f0c52029d8cf8cd0aeaf2365fe4cebc4af26cd6d76b7919ef7"},
			"platform": 1,
			"message":  "Welcome",
		}).
		Run(GetMainEngine(), func(r gofight.HttpResponse, rq gofight.HttpRequest) {

			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestDisabledAndroidPushHandler(t *testing.T) {
	initTest()

	PushConf.Android.Enabled = false

	r := gofight.New()

	r.POST("/api/push").
		SetJSON(gofight.D{
			"tokens":   []string{"aaaaaa", "bbbbb"},
			"platform": 2,
			"message":  "Welcome",
		}).
		Run(GetMainEngine(), func(r gofight.HttpResponse, rq gofight.HttpRequest) {

			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestAndroidPushHandler(t *testing.T) {
	initTest()

	PushConf.Android.Enabled = true

	r := gofight.New()

	r.POST("/api/push").
		SetJSON(gofight.D{
			"tokens":   []string{"aaaaaa", "bbbbb"},
			"platform": 2,
			"message":  "Welcome",
		}).
		Run(GetMainEngine(), func(r gofight.HttpResponse, rq gofight.HttpRequest) {

			assert.Equal(t, http.StatusOK, r.Code)
		})
}
