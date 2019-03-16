package gorush

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStorageDriverExist(t *testing.T) {
	PushConf.Stat.Engine = "Test"
	err := InitAppStatus()
	assert.Error(t, err)
}

func TestStatForMemoryEngine(t *testing.T) {
	// wait android push notification response.
	time.Sleep(5 * time.Second)

	var val int64
	PushConf.Stat.Engine = "memory"
	err := InitAppStatus()
	assert.Nil(t, err)

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

func TestRedisServerSuccess(t *testing.T) {
	PushConf.Stat.Engine = "redis"
	PushConf.Stat.Redis.Addr = "redis:6379"

	err := InitAppStatus()

	assert.NoError(t, err)
}

func TestRedisServerError(t *testing.T) {
	PushConf.Stat.Engine = "redis"
	PushConf.Stat.Redis.Addr = "redis:6370"

	err := InitAppStatus()

	assert.Error(t, err)
}

func TestStatForRedisEngine(t *testing.T) {
	var val int64
	PushConf.Stat.Engine = "redis"
	PushConf.Stat.Redis.Addr = "redis:6379"
	err := InitAppStatus()
	assert.Nil(t, err)

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

func TestDefaultEngine(t *testing.T) {
	var val int64
	// defaul engine as memory
	err := InitAppStatus()
	assert.Nil(t, err)

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

func TestStatForBoltDBEngine(t *testing.T) {
	var val int64
	PushConf.Stat.Engine = "boltdb"
	InitAppStatus()

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

// func TestStatForBuntDBEngine(t *testing.T) {
// 	var val int64
// 	PushConf.Stat.Engine = "buntdb"
// 	err := InitAppStatus()
// 	assert.Nil(t, err)

// 	StatStorage.Reset()

// 	StatStorage.AddTotalCount(100)
// 	StatStorage.AddIosSuccess(200)
// 	StatStorage.AddIosError(300)
// 	StatStorage.AddAndroidSuccess(400)
// 	StatStorage.AddAndroidError(500)

// 	val = StatStorage.GetTotalCount()
// 	assert.Equal(t, int64(100), val)
// 	val = StatStorage.GetIosSuccess()
// 	assert.Equal(t, int64(200), val)
// 	val = StatStorage.GetIosError()
// 	assert.Equal(t, int64(300), val)
// 	val = StatStorage.GetAndroidSuccess()
// 	assert.Equal(t, int64(400), val)
// 	val = StatStorage.GetAndroidError()
// 	assert.Equal(t, int64(500), val)
// }

// func TestStatForLevelDBEngine(t *testing.T) {
// 	var val int64
// 	PushConf.Stat.Engine = "leveldb"
// 	err := InitAppStatus()
// 	assert.Nil(t, err)

// 	StatStorage.Reset()

// 	StatStorage.AddTotalCount(100)
// 	StatStorage.AddIosSuccess(200)
// 	StatStorage.AddIosError(300)
// 	StatStorage.AddAndroidSuccess(400)
// 	StatStorage.AddAndroidError(500)

// 	val = StatStorage.GetTotalCount()
// 	assert.Equal(t, int64(100), val)
// 	val = StatStorage.GetIosSuccess()
// 	assert.Equal(t, int64(200), val)
// 	val = StatStorage.GetIosError()
// 	assert.Equal(t, int64(300), val)
// 	val = StatStorage.GetAndroidSuccess()
// 	assert.Equal(t, int64(400), val)
// 	val = StatStorage.GetAndroidError()
// 	assert.Equal(t, int64(500), val)
// }

// func TestStatForBadgerEngine(t *testing.T) {
// 	var val int64
// 	PushConf.Stat.Engine = "badger"
// 	err := InitAppStatus()
// 	assert.Nil(t, err)

// 	StatStorage.Reset()

// 	StatStorage.AddTotalCount(100)
// 	StatStorage.AddIosSuccess(200)
// 	StatStorage.AddIosError(300)
// 	StatStorage.AddAndroidSuccess(400)
// 	StatStorage.AddAndroidError(500)

// 	val = StatStorage.GetTotalCount()
// 	assert.Equal(t, int64(100), val)
// 	val = StatStorage.GetIosSuccess()
// 	assert.Equal(t, int64(200), val)
// 	val = StatStorage.GetIosError()
// 	assert.Equal(t, int64(300), val)
// 	val = StatStorage.GetAndroidSuccess()
// 	assert.Equal(t, int64(400), val)
// 	val = StatStorage.GetAndroidError()
// 	assert.Equal(t, int64(500), val)
// }
