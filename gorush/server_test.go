package gopush

import (
	"github.com/appleboy/gofight"
	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
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

func testRequest(t *testing.T, url string) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	assert.NoError(t, err)

	_, ioerr := ioutil.ReadAll(resp.Body)
	assert.NoError(t, ioerr)
	assert.Equal(t, "200 OK", resp.Status, "should get a 200")
}

func TestPrintGoPushVersion(t *testing.T) {
	PrintGoPushVersion()
}

func TestRunNormalServer(t *testing.T) {
	initTest()

	router := gin.New()

	go func() {
		assert.NoError(t, RunHTTPServer())
	}()
	// have to wait for the goroutine to start and run the server
	// otherwise the main thread will complete
	time.Sleep(5 * time.Millisecond)

	assert.Error(t, router.Run(":8088"))
	testRequest(t, "http://localhost:8088/api/status")
}

// func TestRunTLSServer(t *testing.T) {
// 	initTest()

// 	PushConf.Core.SSL = true
// 	PushConf.Core.Port = "8087"
// 	PushConf.Core.CertPath = "../certificate/localhost.cert"
// 	PushConf.Core.KeyPath = "../certificate/localhost.key"
// 	router := gin.New()

// 	go func() {
// 		assert.NoError(t, RunHTTPServer())
// 	}()
// 	// have to wait for the goroutine to start and run the server
// 	// otherwise the main thread will complete
// 	time.Sleep(5 * time.Millisecond)

// 	assert.Error(t, router.Run(":8087"))
// 	testRequest(t, "https://localhost:8087/api/status")
// }

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

func TestMissingIosCertificate(t *testing.T) {
	initTest()

	PushConf.Ios.Enabled = true
	PushConf.Ios.PemKeyPath = "test"
	err := InitAPNSClient()

	assert.Error(t, err)
}

func TestIosPushDevelopment(t *testing.T) {
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

func TestIosPushProduction(t *testing.T) {
	initTest()

	PushConf.Ios.Enabled = true
	PushConf.Ios.Production = true
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

func TestAndroidWrongAPIKey(t *testing.T) {
	initTest()

	PushConf.Android.Enabled = true
	PushConf.Android.ApiKey = os.Getenv("ANDROID_API_KEY") + "a"

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

func TestAndroidWrongToken(t *testing.T) {
	initTest()

	PushConf.Android.Enabled = true
	PushConf.Android.ApiKey = os.Getenv("ANDROID_API_KEY")

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

func TestAndroidRightToken(t *testing.T) {
	initTest()

	PushConf.Android.Enabled = true
	PushConf.Android.ApiKey = os.Getenv("ANDROID_API_KEY")

	android_token := os.Getenv("ANDROID_TEST_TOKEN")

	r := gofight.New()

	r.POST("/api/push").
		SetJSON(gofight.D{
			"tokens":   []string{android_token, "bbbbb"},
			"platform": 2,
			"message":  "Welcome",
		}).
		Run(GetMainEngine(), func(r gofight.HttpResponse, rq gofight.HttpRequest) {

			assert.Equal(t, http.StatusOK, r.Code)
		})
}
