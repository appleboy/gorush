package gorush

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jaraxasoftware/gorush/web"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	api "gopkg.in/appleboy/gin-status-api.v1"
	apns "github.com/sideshow/apns2"
)

func init() {
	// Support metrics
	m := NewMetrics()
	prometheus.MustRegister(m)
}

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
	var sync bool

	if err := c.BindJSON(&form); err != nil {
		msg = "Missing notifications field. "
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

	if int64(len(form.Notifications)) > PushConf.Core.MaxNotification {
		msg = fmt.Sprintf("Number of notifications(%d) over limit(%d)", len(form.Notifications), PushConf.Core.MaxNotification)
		LogAccess.Debug(msg)
		abortWithError(c, http.StatusBadRequest, msg)
		return
	}

	sync = form.Sync

	if sync {
		isError := false
		var apnsFailedResults map[string]*apns.Response
		var gcmFailedResults map[string]string
		var webFailedResults map[string]*web.Response
		apnsFailedResults = make(map[string]*apns.Response)
		gcmFailedResults = make(map[string]string)
		webFailedResults = make(map[string]*web.Response)
		for i := 0; i < len(form.Notifications); i++ {
			var isErrorLoop bool
			var apnsFailedResultsLoop *map[string]*apns.Response
			var gcmFailedResultsLoop *map[string]string
			var webFailedResultsLoop *map[string]*web.Response
			notification := form.Notifications[i]

			switch notification.Platform {
			case PlatformIos:
				var enabled bool
				var notEnabledMessage string
				if notification.Voip {
					enabled = PushConf.Ios.VoipEnabled
					notEnabledMessage = "VoIP iOS is not enabled in configuration"
				} else {
					enabled = PushConf.Ios.Enabled
					notEnabledMessage = "iOS is not enabled in configuration"
				}
				if enabled {
					apnsFailedResultsLoop, isErrorLoop = PushToIOSWithErrorResult(notification)
					if apnsFailedResultsLoop != nil {
						for k, v := range *apnsFailedResultsLoop {
							apnsFailedResults[k] = v
						}
					}
				} else {
					for _, token := range notification.Tokens {
						time := apns.Time{time.Now()}
						response := apns.Response{StatusCode: 500, Reason: notEnabledMessage, Timestamp: time}
						apnsFailedResults[token] = &response
					}
				}
			case PlatformAndroid:
				if PushConf.Android.Enabled {
					gcmFailedResultsLoop, isErrorLoop = PushToAndroidWithErrorResult(notification)
					if gcmFailedResultsLoop != nil {
						for k, v := range *gcmFailedResultsLoop {
							gcmFailedResults[k] = v
						}
					}
				} else {
					for _, token := range notification.Tokens {
						gcmFailedResults[token] = "Android is not enabled in configuration"
					}					
				}
			case PlatformWeb:
				if PushConf.Web.Enabled {
					webFailedResultsLoop, isErrorLoop = PushToWebWithErrorResult(notification)
					if webFailedResultsLoop != nil {
						for k, v := range *webFailedResultsLoop {
							webFailedResults[k] = v
						}
					}
				} else {
					for _, token := range notification.Tokens {
						response := web.Response{StatusCode: 500, Body: "Web is not enabled in configuration"}
						webFailedResults[token] = &response
					}					
				}
			}
			isError = isError || isErrorLoop
		}
		c.JSON(http.StatusOK, gin.H{
			"success": "ok",
			"apnsFailedResults": apnsFailedResults,
			"gcmFailedResults": gcmFailedResults,
			"webFailedResults": webFailedResults,
		})
	} else {
		// queue notification.
		go queueNotification(form)

		c.JSON(http.StatusOK, gin.H{
			"success": "ok",
		})
	}
}

func configHandler(c *gin.Context) {
	c.YAML(http.StatusCreated, PushConf)
}

func metricsHandler(c *gin.Context) {
	promhttp.Handler().ServeHTTP(c.Writer, c.Request)
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
	r.Use(StatMiddleware())

	r.GET(PushConf.API.StatGoURI, api.StatusHandler)
	r.GET(PushConf.API.StatAppURI, appStatusHandler)
	r.GET(PushConf.API.ConfigURI, configHandler)
	r.GET(PushConf.API.SysStatURI, sysStatsHandler)
	r.POST(PushConf.API.PushURI, pushHandler)
	r.GET(PushConf.API.MetricURI, metricsHandler)
	r.GET("/", rootHandler)

	return r
}
