package gorush

import (
	"github.com/stretchr/testify/assert"
	// "sync/atomic"
	"testing"
)

// func TestAddTotalCount(t *testing.T) {
// 	InitAppStatus()
// 	addTotalCount(1000)

// 	val := atomic.LoadInt64(&RushStatus.TotalCount)

// 	assert.Equal(t, int64(1000), val)
// }

// func TestAddIosSuccess(t *testing.T) {
// 	InitAppStatus()
// 	addIosSuccess(1000)

// 	val := atomic.LoadInt64(&RushStatus.Ios.PushSuccess)

// 	assert.Equal(t, int64(1000), val)
// }

// func TestAddIosError(t *testing.T) {
// 	InitAppStatus()
// 	addIosError(1000)

// 	val := atomic.LoadInt64(&RushStatus.Ios.PushError)

// 	assert.Equal(t, int64(1000), val)
// }

// func TestAndroidSuccess(t *testing.T) {
// 	InitAppStatus()
// 	addAndroidSuccess(1000)

// 	val := atomic.LoadInt64(&RushStatus.Android.PushSuccess)

// 	assert.Equal(t, int64(1000), val)
// }

// func TestAddAndroidError(t *testing.T) {
// 	InitAppStatus()
// 	addAndroidError(1000)

// 	val := atomic.LoadInt64(&RushStatus.Android.PushError)

// 	assert.Equal(t, int64(1000), val)
// }

func TestStatForMemoryEngine(t *testing.T) {
	var val int64
	PushConf.Stat.Engine = "memory"
	InitAppStatus()

	StatStorage.AddTotalCount(100)
	StatStorage.AddIosSuccess(200)
	StatStorage.AddIosError(300)
	StatStorage.AddAndroidSuccess(400)
	StatStorage.AddAndroidError(500)

	val = StatStorage.GetTotalCount()
	assert.Equal(t, int64(100), val)
	val = StatStorage.GetIosSuccess()
	assert.Equal(t, int64(200), val)
	val = StatStorage.GetIosError()
	assert.Equal(t, int64(300), val)
	val = StatStorage.GetAndroidSuccess()
	assert.Equal(t, int64(400), val)
	val = StatStorage.GetAndroidError()
	assert.Equal(t, int64(500), val)
}

// func TestRedisServerSuccess(t *testing.T) {
// 	PushConf.Stat.Redis.Addr = "localhost:6379"

// 	err := initRedis()

// 	assert.NoError(t, err)
// }

// func TestRedisServerError(t *testing.T) {
// 	PushConf.Stat.Redis.Addr = "localhost:6370"

// 	err := initRedis()

// 	assert.Error(t, err)
// }

func TestStatForRedisEngine(t *testing.T) {
	var val int64
	PushConf.Stat.Engine = "redis"
	PushConf.Stat.Redis.Addr = "localhost:6379"
	InitAppStatus()

	StatStorage.Init()
	StatStorage.Reset()

	StatStorage.AddTotalCount(100)
	StatStorage.AddIosSuccess(200)
	StatStorage.AddIosError(300)
	StatStorage.AddAndroidSuccess(400)
	StatStorage.AddAndroidError(500)

	val = StatStorage.GetTotalCount()
	assert.Equal(t, int64(100), val)
	val = StatStorage.GetIosSuccess()
	assert.Equal(t, int64(200), val)
	val = StatStorage.GetIosError()
	assert.Equal(t, int64(300), val)
	val = StatStorage.GetAndroidSuccess()
	assert.Equal(t, int64(400), val)
	val = StatStorage.GetAndroidError()
	assert.Equal(t, int64(500), val)
}

// func TestDefaultEngine(t *testing.T) {
// 	var val int64
// 	PushConf.Stat.Engine = "test"
// 	InitAppStatus()

// 	addTotalCount(1)
// 	addIosSuccess(2)
// 	addIosError(3)
// 	addAndroidSuccess(4)
// 	addAndroidError(5)

// 	val = getTotalCount()
// 	assert.Equal(t, int64(1), val)
// 	val = getIosSuccess()
// 	assert.Equal(t, int64(2), val)
// 	val = getIosError()
// 	assert.Equal(t, int64(3), val)
// 	val = getAndroidSuccess()
// 	assert.Equal(t, int64(4), val)
// 	val = getAndroidError()
// 	assert.Equal(t, int64(5), val)
// }

// func TestStatForBoltDBEngine(t *testing.T) {
// 	var val int64
// 	PushConf.Stat.Engine = "boltdb"
// 	InitAppStatus()

// 	addTotalCount(100)
// 	addIosSuccess(200)
// 	addIosError(300)
// 	addAndroidSuccess(400)
// 	addAndroidError(500)

// 	val = getTotalCount()
// 	assert.Equal(t, int64(100), val)
// 	val = getIosSuccess()
// 	assert.Equal(t, int64(200), val)
// 	val = getIosError()
// 	assert.Equal(t, int64(300), val)
// 	val = getAndroidSuccess()
// 	assert.Equal(t, int64(400), val)
// 	val = getAndroidError()
// 	assert.Equal(t, int64(500), val)
// }
