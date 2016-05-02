package gorush

import (
	"github.com/appleboy/gorush/gorush/storage/memory"
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
func InitAppStatus() {
	switch PushConf.Stat.Engine {
	case "memory":
		StatStorage = memory.New()
	// case "redis":
	// 	initRedis()
	// case "boltdb":
	// 	initBoltDB()
	default:
		StatStorage = memory.New()
	}
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
