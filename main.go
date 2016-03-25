package main

import (
	api "github.com/appleboy/gin-status-api"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
		"text": "Welcome to golang push server.",
	})
}

func pushHandler(c *gin.Context) {
	var form RequestPushNotification

	if err := c.BindJSON(&form); err != nil {
		log.Println(err)
		AbortWithError(c, http.StatusBadRequest, "Bad input request, please refer to README guide.")
		return
	}

	// process notification.
	pushNotification(form)

	c.JSON(http.StatusOK, gin.H{
		"text": "Welcome to golang push server.",
	})
}

func GetMainEngine() *gin.Engine {
	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(VersionMiddleware())

	r.GET("/api/status", api.StatusHandler)
	r.POST("/api/push", pushHandler)
	r.GET("/", rootHandler)

	return r
}

func main() {
	endless.ListenAndServe(":8088", GetMainEngine())
}
