package gorush

import (
	"github.com/appleboy/gorush/storage/boltdb"
	"github.com/appleboy/gorush/storage/memory"
	"github.com/appleboy/gorush/storage/redis"
	"github.com/gin-gonic/gin"
	"github.com/thoas/stats"
	"net/http"
)

// Stats provide response time, status code count, etc.
var Stats = stats.New()

// StatusApp is app status structure
type StatusApp struct {
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
	switch PushConf.Stat.Engine {
	case "memory":
		StatStorage = memory.New()
	case "redis":
		StatStorage = redis.New(PushConf)
		err := StatStorage.Init()

		if err != nil {
			LogError.Error("redis error: " + err.Error())

			return err
		}

	case "boltdb":
		StatStorage = boltdb.New(PushConf)
	default:
		StatStorage = memory.New()
	}

	return nil
}

func appStatusHandler(c *gin.Context) {
	result := StatusApp{}

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
		Stats.End(beginning, recorder)
	}
}
