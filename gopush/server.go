package gopush

import (
	api "github.com/appleboy/gin-status-api"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func AbortWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
	c.Abort()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"text": "Welcome to notification server.",
	})
}

func pushHandler(c *gin.Context) {
	var form RequestPush
	var msg string

	if err := c.BindJSON(&form); err != nil {
		msg = "Missing notifications field."
		LogAccess.Debug(msg)
		AbortWithError(c, http.StatusBadRequest, msg)
		return
	}

	if len(form.Notifications) == 0 {
		msg = "Notification field is empty."
		LogAccess.Debug(msg)
		AbortWithError(c, http.StatusBadRequest, msg)
		return
	}

	if len(form.Notifications) > PushConf.Core.MaxNotification {
		msg = fmt.Sprintf("Number of notifications(%d) over limit(%d)", len(form.Notifications), PushConf.Core.MaxNotification)
		LogAccess.Debug(msg)
		AbortWithError(c, http.StatusBadRequest, msg)
		return
	}

	// process notification.
	go SendNotification(form)

	c.JSON(http.StatusOK, gin.H{
		"text": "Welcome to notification server.",
	})
}

func GetMainEngine() *gin.Engine {
	// set server mode
	gin.SetMode(PushConf.Core.Mode)

	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(VersionMiddleware())
	r.Use(LogMiddleware())

	r.GET(PushConf.Api.StatGoUri, api.StatusHandler)
	r.POST(PushConf.Api.PushUri, pushHandler)
	r.GET("/", rootHandler)

	return r
}

func RunHTTPServer() error {
	var err error
	if PushConf.Core.SSL && PushConf.Core.CertPath != "" && PushConf.Core.KeyPath != "" {
		err = GetMainEngine().RunTLS(":"+PushConf.Core.Port, PushConf.Core.CertPath, PushConf.Core.KeyPath)
	} else {
		err = GetMainEngine().Run(":" + PushConf.Core.Port)
	}

	return err
}
