package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	api "github.com/appleboy/gin-status-api"
)

func rootHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"text": "hello world",
	})
}

func GetMainEngine() *gin.Engine {
	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/api/status", api.StatusHandler)
	r.GET("/", rootHandler)

	return r
}

func main() {
	GetMainEngine().Run(":8088")
}
