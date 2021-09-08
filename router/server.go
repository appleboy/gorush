package router

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/core"
	"github.com/appleboy/gorush/logx"
	"github.com/appleboy/gorush/metric"
	"github.com/appleboy/gorush/notify"
	"github.com/appleboy/gorush/status"

	api "github.com/appleboy/gin-status-api"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-queue/queue"
	"github.com/mattn/go-isatty"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/thoas/stats"
	"golang.org/x/crypto/acme/autocert"
)

var doOnce sync.Once

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

func pushHandler(cfg *config.ConfYaml, q *queue.Queue) gin.HandlerFunc {
	return func(c *gin.Context) {
		var form notify.RequestPush
		var msg string

		if err := c.ShouldBindWith(&form, binding.JSON); err != nil {
			msg = "Missing notifications field."
			logx.LogAccess.Debug(err)
			abortWithError(c, http.StatusBadRequest, msg)
			return
		}

		if len(form.Notifications) == 0 {
			msg = "Notifications field is empty."
			logx.LogAccess.Debug(msg)
			abortWithError(c, http.StatusBadRequest, msg)
			return
		}

		if int64(len(form.Notifications)) > cfg.Core.MaxNotification {
			msg = fmt.Sprintf("Number of notifications(%d) over limit(%d)", len(form.Notifications), cfg.Core.MaxNotification)
			logx.LogAccess.Debug(msg)
			abortWithError(c, http.StatusBadRequest, msg)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			// Deprecated: the CloseNotifier interface predates Go's context package.
			// New code should use Request.Context instead.
			// Change to context package
			<-c.Request.Context().Done()
			// Don't send notification after client timeout or disconnected.
			// See the following issue for detail information.
			// https://github.com/appleboy/gorush/issues/422
			if cfg.Core.Sync {
				cancel()
			}
		}()

		counts, logs := handleNotification(ctx, cfg, form, q)

		c.JSON(http.StatusOK, gin.H{
			"success": "ok",
			"counts":  counts,
			"logs":    logs,
		})
	}
}

func configHandler(cfg *config.ConfYaml) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.YAML(http.StatusCreated, cfg)
	}
}

func metricsHandler(c *gin.Context) {
	promhttp.Handler().ServeHTTP(c.Writer, c.Request)
}

func appStatusHandler(q *queue.Queue) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := status.App{}

		result.Version = GetVersion()
		result.QueueMax = q.Capacity()
		result.QueueUsage = q.Usage()
		result.TotalCount = status.StatStorage.GetTotalCount()
		result.Ios.PushSuccess = status.StatStorage.GetIosSuccess()
		result.Ios.PushError = status.StatStorage.GetIosError()
		result.Android.PushSuccess = status.StatStorage.GetAndroidSuccess()
		result.Android.PushError = status.StatStorage.GetAndroidError()
		result.Huawei.PushSuccess = status.StatStorage.GetHuaweiSuccess()
		result.Huawei.PushError = status.StatStorage.GetHuaweiError()

		c.JSON(http.StatusOK, result)
	}
}

func sysStatsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, status.Stats.Data())
	}
}

// StatMiddleware response time, status code count, etc.
func StatMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		beginning, recorder := status.Stats.Begin(c.Writer)
		c.Next()
		status.Stats.End(beginning, stats.WithRecorder(recorder))
	}
}

func autoTLSServer(cfg *config.ConfYaml, q *queue.Queue) *http.Server {
	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(cfg.Core.AutoTLS.Host),
		Cache:      autocert.DirCache(cfg.Core.AutoTLS.Folder),
	}

	return &http.Server{
		Addr:      ":https",
		TLSConfig: &tls.Config{GetCertificate: m.GetCertificate},
		Handler:   routerEngine(cfg, q),
	}
}

func routerEngine(cfg *config.ConfYaml, q *queue.Queue) *gin.Engine {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cfg.Core.Mode == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

	isTerm := isatty.IsTerminal(os.Stdout.Fd())
	if isTerm {
		log.Logger = log.Output(
			zerolog.ConsoleWriter{
				Out:     os.Stdout,
				NoColor: false,
			},
		)
	}

	// Support metrics
	doOnce.Do(func() {
		m := metric.NewMetrics(func() int {
			return q.Usage()
		})
		prometheus.MustRegister(m)
	})

	// set server mode
	gin.SetMode(cfg.Core.Mode)

	r := gin.New()

	// Global middleware
	r.Use(logger.SetLogger(
		logger.WithUTC(true),
		logger.WithSkipPath([]string{
			cfg.API.HealthURI,
			cfg.API.MetricURI,
		}),
	))
	r.Use(gin.Recovery())
	r.Use(VersionMiddleware())
	r.Use(StatMiddleware())

	r.GET(cfg.API.StatGoURI, api.GinHandler)
	r.GET(cfg.API.StatAppURI, appStatusHandler(q))
	r.GET(cfg.API.ConfigURI, configHandler(cfg))
	r.GET(cfg.API.SysStatURI, sysStatsHandler())
	r.POST(cfg.API.PushURI, pushHandler(cfg, q))
	r.GET(cfg.API.MetricURI, metricsHandler)
	r.GET(cfg.API.HealthURI, heartbeatHandler)
	r.HEAD(cfg.API.HealthURI, heartbeatHandler)
	r.GET("/version", versionHandler)
	r.GET("/", rootHandler)

	return r
}

// markFailedNotification adds failure logs for all tokens in push notification
func markFailedNotification(cfg *config.ConfYaml, notification *notify.PushNotification, reason string) []logx.LogPushEntry {
	logx.LogError.Error(reason)
	logs := make([]logx.LogPushEntry, 0)
	for _, token := range notification.Tokens {
		logs = append(logs, logx.GetLogPushEntry(&logx.InputLog{
			ID:        notification.ID,
			Status:    core.FailedPush,
			Token:     token,
			Message:   notification.Message,
			Platform:  notification.Platform,
			Error:     errors.New(reason),
			HideToken: cfg.Log.HideToken,
			Format:    cfg.Log.Format,
		}))
	}

	return logs
}

// HandleNotification add notification to queue list.
func handleNotification(ctx context.Context, cfg *config.ConfYaml, req notify.RequestPush, q *queue.Queue) (int, []logx.LogPushEntry) {
	var count int
	wg := sync.WaitGroup{}
	newNotification := []*notify.PushNotification{}

	if cfg.Core.Sync && !core.IsLocalQueue(core.Queue(cfg.Queue.Engine)) {
		cfg.Core.Sync = false
	}

	for i := range req.Notifications {
		notification := &req.Notifications[i]
		switch notification.Platform {
		case core.PlatFormIos:
			if !cfg.Ios.Enabled {
				continue
			}
		case core.PlatFormAndroid:
			if !cfg.Android.Enabled {
				continue
			}
		case core.PlatFormHuawei:
			if !cfg.Huawei.Enabled {
				continue
			}
		}
		newNotification = append(newNotification, notification)
	}

	logs := make([]logx.LogPushEntry, 0, count)
	for _, notification := range newNotification {
		if cfg.Core.Sync {
			wg.Add(1)
		}

		if core.IsLocalQueue(core.Queue(cfg.Queue.Engine)) && cfg.Core.Sync {
			func(msg *notify.PushNotification, cfg *config.ConfYaml) {
				if err := q.QueueTask(func(ctx context.Context) error {
					defer wg.Done()
					resp, err := notify.SendNotification(msg, cfg)
					if err != nil {
						return err
					}

					// add log
					logs = append(logs, resp.Logs...)

					return nil
				}); err != nil {
					logx.LogError.Error(err)
				}
			}(notification, cfg)
		} else if err := q.Queue(notification); err != nil {
			resp := markFailedNotification(cfg, notification, "max capacity reached")
			// add log
			logs = append(logs, resp...)
			wg.Done()
		}

		count += len(notification.Tokens)
		// Count topic message
		if notification.To != "" {
			count++
		}
	}

	if cfg.Core.Sync {
		wg.Wait()
	}

	status.StatStorage.AddTotalCount(int64(count))

	return count, logs
}
