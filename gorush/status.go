package gorush

import (
	"sync/atomic"
)

type StatusApp struct {
	QueueMax   int           `json:"queue_max"`
	QueueUsage int           `json:"queue_usage"`
	TotalCount int64         `json:"total_count"`
	Ios        IosStatus     `json:"ios"`
	Android    AndroidStatus `json:"android"`
}

type AndroidStatus struct {
	PushSuccess int64 `json:"push_success"`
	PushError   int64 `json:"push_error"`
}

type IosStatus struct {
	PushSuccess int64 `json:"push_success"`
	PushError   int64 `json:"push_error"`
}

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
