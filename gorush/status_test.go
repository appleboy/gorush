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

func TestRedisServerSuccess(t *testing.T) {
	PushConf.Stat.Redis.Addr = "localhost:6379"

	err := initRedis()

	assert.NoError(t, err)
}

func TestRedisServerError(t *testing.T) {
	PushConf.Stat.Redis.Addr = "localhost:6370"

	err := initRedis()

	assert.Error(t, err)
}

func TestStatForRedisEngine(t *testing.T) {
	var val int64
	PushConf.Stat.Engine = "redis"
	PushConf.Stat.Redis.Addr = "localhost:6379"
	InitAppStatus()

	addTotalCount(1000)
	addIosSuccess(1000)
	addIosError(1000)
	addAndroidSuccess(1000)
	addAndroidError(1000)

	val = getTotalCount()
	assert.Equal(t, int64(1000), val)
	val = getIosSuccess()
	assert.Equal(t, int64(1000), val)
	val = getIosError()
	assert.Equal(t, int64(1000), val)
	val = getAndroidSuccess()
	assert.Equal(t, int64(1000), val)
	val = getAndroidError()
	assert.Equal(t, int64(1000), val)
}

func TestDefaultEngine(t *testing.T) {
	var val int64
	PushConf.Stat.Engine = "test"
	InitAppStatus()

	addTotalCount(1000)
	addIosSuccess(1000)
	addIosError(1000)
	addAndroidSuccess(1000)
	addAndroidError(1000)

	val = getTotalCount()
	assert.Equal(t, int64(1000), val)
	val = getIosSuccess()
	assert.Equal(t, int64(1000), val)
	val = getIosError()
	assert.Equal(t, int64(1000), val)
	val = getAndroidSuccess()
	assert.Equal(t, int64(1000), val)
	val = getAndroidError()
	assert.Equal(t, int64(1000), val)
}
