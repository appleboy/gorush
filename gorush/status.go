package gorush

import (
	"github.com/appleboy/gorush/gorush/storage/boltdb"
	"github.com/appleboy/gorush/gorush/storage/memory"
	"github.com/appleboy/gorush/gorush/storage/redis"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
