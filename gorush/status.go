package gorush

import (
	"errors"
	"net/http"

	"github.com/appleboy/gorush/storage/badger"
	"github.com/appleboy/gorush/storage/boltdb"
	"github.com/appleboy/gorush/storage/buntdb"
	"github.com/appleboy/gorush/storage/leveldb"
	"github.com/appleboy/gorush/storage/memory"
	"github.com/appleboy/gorush/storage/redis"
	"github.com/gin-gonic/gin"
	"github.com/thoas/stats"
)

// Stats provide response time, status code count, etc.
var Stats = stats.New()

// StatusApp is app status structure
type StatusApp struct {
	Version    string        `json:"version"`
	QueueMax   int           `json:"queue_max"`
	QueueUsage int           `json:"queue_usage"`
	TotalCount int64         `json:"total_count"`
	Ios        IosStatus     `json:"ios"`
	Android    AndroidStatus `json:"android"`
}

// AndroidStatus is android structure
type AndroidStatus struct {
	PushSuccess int64 `json:"push_success"`
	PushError   int64 `json:"push_error"`
}

// IosStatus is iOS structure
type IosStatus struct {
	PushSuccess int64 `json:"push_success"`
	PushError   int64 `json:"push_error"`
}

// InitAppStatus for initialize app status
func InitAppStatus() error {
	LogAccess.Debug("Init App Status Engine as ", PushConf.Stat.Engine)
	switch PushConf.Stat.Engine {
	case "memory":
		StatStorage = memory.New()
	case "redis":
		StatStorage = redis.New(PushConf)
	case "boltdb":
		StatStorage = boltdb.New(PushConf)
	case "buntdb":
		StatStorage = buntdb.New(PushConf)
	case "leveldb":
		StatStorage = leveldb.New(PushConf)
	case "badger":
		StatStorage = badger.New(PushConf)
	default:
		LogError.Error("storage error: can't find storage driver")
		return errors.New("can't find storage driver")
	}

	if err := StatStorage.Init(); err != nil {
		LogError.Error("storage error: " + err.Error())

		return err
	}

	return nil
}

func appStatusHandler(c *gin.Context) {
	result := StatusApp{}

	result.Version = GetVersion()
	result.QueueMax = cap(QueueNotification)
	result.QueueUsage = len(QueueNotification)
	result.TotalCount = StatStorage.GetTotalCount()
	result.Ios.PushSuccess = StatStorage.GetIosSuccess()
	result.Ios.PushError = StatStorage.GetIosError()
	result.Android.PushSuccess = StatStorage.GetAndroidSuccess()
	result.Android.PushError = StatStorage.GetAndroidError()

	c.JSON(http.StatusOK, result)
}

func sysStatsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, Stats.Data())
}

// StatMiddleware response time, status code count, etc.
func StatMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		beginning, recorder := Stats.Begin(c.Writer)
		c.Next()
		Stats.End(beginning, stats.WithRecorder(recorder))
	}
}
