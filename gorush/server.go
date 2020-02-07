package gorush

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"regexp"

	api "github.com/appleboy/gin-status-api"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/acme/autocert"
)

var (
	rxURL = regexp.MustCompile(`^/healthz$`)
)

func init() {
	// Support metrics
	m := NewMetrics()
	prometheus.MustRegister(m)
}

func abortWithError(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"text": "Welcome to notification server.",
	})
}

func heartbeatHandler(c *gin.Context) {
	c.AbortWithStatus(http.StatusOK)
}

func versionHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"source":  "https://github.com/appleboy/gorush",
		"version": GetVersion(),
	})
}

func pushHandler(c *gin.Context) {
	var form RequestPush
	var msg string

	if err := c.ShouldBindWith(&form, binding.JSON); err != nil {
		msg = "Missing notifications field."
		LogAccess.Debug(err)
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

	ctx, cancel := context.WithCancel(context.Background())
	notifier := c.Writer.CloseNotify()
	go func(closer <-chan bool) {
		<-closer
		// Don't send notification after client timeout or disconnected.
		// See the following issue for detail information.
		// https://github.com/appleboy/gorush/issues/422
		if PushConf.Core.Sync {
			cancel()
		}
	}(notifier)

	counts, logs := queueNotification(ctx, form)

	c.JSON(http.StatusOK, gin.H{
		"success": "ok",
		"counts":  counts,
		"logs":    logs,
	})
}

func configHandler(c *gin.Context) {
	c.YAML(http.StatusCreated, PushConf)
}

func metricsHandler(c *gin.Context) {
	promhttp.Handler().ServeHTTP(c.Writer, c.Request)
}

func autoTLSServer() *http.Server {
	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(PushConf.Core.AutoTLS.Host),
		Cache:      autocert.DirCache(PushConf.Core.AutoTLS.Folder),
	}

	return &http.Server{
		Addr:      ":https",
		TLSConfig: &tls.Config{GetCertificate: m.GetCertificate},
		Handler:   routerEngine(),
	}
}

func routerEngine() *gin.Engine {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if PushConf.Core.Mode == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

	if isTerm {
		log.Logger = log.Output(
			zerolog.ConsoleWriter{
				Out:     os.Stdout,
				NoColor: false,
			},
		)
	}

	// set server mode
	gin.SetMode(PushConf.Core.Mode)

	r := gin.New()

	// Global middleware
	r.Use(logger.SetLogger(logger.Config{
		UTC:            true,
		SkipPathRegexp: rxURL,
	}))
	r.Use(gin.Recovery())
	r.Use(VersionMiddleware())
	r.Use(StatMiddleware())

	r.GET(PushConf.API.StatGoURI, api.GinHandler)
	r.GET(PushConf.API.StatAppURI, appStatusHandler)
	r.GET(PushConf.API.ConfigURI, configHandler)
	r.GET(PushConf.API.SysStatURI, sysStatsHandler)
	r.POST(PushConf.API.PushURI, pushHandler)
	r.GET(PushConf.API.MetricURI, metricsHandler)
	r.GET(PushConf.API.HealthURI, heartbeatHandler)
	r.GET("/version", versionHandler)
	r.GET("/", rootHandler)

	return r
}
