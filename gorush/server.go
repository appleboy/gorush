package gorush

import (
	"fmt"
	api "github.com/appleboy/gin-status-api"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"net/http"
)

func abortWithError(c *gin.Context, code int, message string) {
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
		abortWithError(c, http.StatusBadRequest, msg)
		return
	}

	if len(form.Notifications) == 0 {
		msg = "Notifications field is empty."
		LogAccess.Debug(msg)
		abortWithError(c, http.StatusBadRequest, msg)
		return
	}

	if len(form.Notifications) > PushConf.Core.MaxNotification {
		msg = fmt.Sprintf("Number of notifications(%d) over limit(%d)", len(form.Notifications), PushConf.Core.MaxNotification)
		LogAccess.Debug(msg)
		abortWithError(c, http.StatusBadRequest, msg)
		return
	}

	// queue notification.
	go queueNotification(form)

	c.JSON(http.StatusOK, gin.H{
		"success": "ok",
	})
}

func configHandler(c *gin.Context) {
	c.YAML(http.StatusCreated, PushConf)
}

func routerEngine() *gin.Engine {
	// set server mode
	gin.SetMode(PushConf.Core.Mode)

	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(VersionMiddleware())
	r.Use(LogMiddleware())

	r.GET(PushConf.API.StatGoURI, api.StatusHandler)
	r.GET(PushConf.API.StatAppURI, appStatusHandler)
	r.GET(PushConf.API.ConfigURI, configHandler)
	r.POST(PushConf.API.PushURI, pushHandler)
	r.GET("/", rootHandler)

	return r
}

// RunHTTPServer provide run http or https protocol.
func RunHTTPServer() error {
	var err error
	if PushConf.Core.SSL && PushConf.Core.CertPath != "" && PushConf.Core.KeyPath != "" {
		err = endless.ListenAndServeTLS(":"+PushConf.Core.Port, PushConf.Core.CertPath, PushConf.Core.KeyPath, routerEngine())
	} else {
		err = endless.ListenAndServe(":"+PushConf.Core.Port, routerEngine())
	}

	return err
}
