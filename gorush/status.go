package gorush

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/atomic"
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
	RushStatus.TotalCount = 0
	RushStatus.Ios.PushSuccess = 0
	RushStatus.Ios.PushError = 0
	RushStatus.Android.PushSuccess = 0
	RushStatus.Android.PushError = 0
}

func addTotalCount(count int64) {
	atomic.AddInt64(&RushStatus.TotalCount, count)
}

func addIosSuccess(count int64) {
	atomic.AddInt64(&RushStatus.Ios.PushSuccess, count)
}

func addIosError(count int64) {
	atomic.AddInt64(&RushStatus.Ios.PushError, count)
}

func addAndroidSuccess(count int64) {
	atomic.AddInt64(&RushStatus.Android.PushSuccess, count)
}

func addAndroidError(count int64) {
	atomic.AddInt64(&RushStatus.Android.PushError, count)
}

func appStatusHandler(c *gin.Context) {
	result := StatusApp{}

	result.QueueMax = cap(QueueNotification)
	result.QueueUsage = len(QueueNotification)
	result.TotalCount = atomic.LoadInt64(&RushStatus.TotalCount)
	result.Ios.PushSuccess = atomic.LoadInt64(&RushStatus.Ios.PushSuccess)
	result.Ios.PushError = atomic.LoadInt64(&RushStatus.Ios.PushError)
	result.Android.PushSuccess = atomic.LoadInt64(&RushStatus.Android.PushSuccess)
	result.Android.PushError = atomic.LoadInt64(&RushStatus.Android.PushError)

	c.JSON(http.StatusOK, result)
}
