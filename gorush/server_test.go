package gopush

import (
	"github.com/appleboy/gofight"
	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func initTest() {
	gin.SetMode(gin.TestMode)
	PushConf = BuildDefaultPushConf()
}

func TestGinHelloWorld(t *testing.T) {
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
