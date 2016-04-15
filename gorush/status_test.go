package gorush

import (
	"github.com/stretchr/testify/assert"
	"sync/atomic"
	"testing"
)

func TestAddTotalCount(t *testing.T) {
	InitAppStatus()
	addTotalCount(1000)

	val := atomic.LoadInt64(&RushStatus.TotalCount)

	assert.Equal(t, int64(1000), val)
}

func TestAddIosSuccess(t *testing.T) {
	InitAppStatus()
	addIosSuccess(1000)

	val := atomic.LoadInt64(&RushStatus.Ios.PushSuccess)

	assert.Equal(t, int64(1000), val)
}

func TestAddIosError(t *testing.T) {
	InitAppStatus()
	addIosError(1000)

	val := atomic.LoadInt64(&RushStatus.Ios.PushError)

	assert.Equal(t, int64(1000), val)
}

func TestAndroidSuccess(t *testing.T) {
	InitAppStatus()
	addAndroidSuccess(1000)

	val := atomic.LoadInt64(&RushStatus.Android.PushSuccess)

	assert.Equal(t, int64(1000), val)
}

func TestAddAndroidError(t *testing.T) {
	InitAppStatus()
	addAndroidError(1000)

	val := atomic.LoadInt64(&RushStatus.Android.PushError)

	assert.Equal(t, int64(1000), val)
}
