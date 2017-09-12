package redis

import (
	"testing"

	c "github.com/axiomzen/gorush/config"
	"github.com/stretchr/testify/assert"
)

func TestRedisServerError(t *testing.T) {
	config := c.BuildDefaultPushConf()
	config.Stat.Redis.Addr = "redis:6370"

	redis := New(config)
	err := redis.Init()

	assert.Error(t, err)
}

func TestRedisEngine(t *testing.T) {
	var val int64

	config := c.BuildDefaultPushConf()
	config.Stat.Redis.Addr = "redis:6379"

	redis := New(config)
	err := redis.Init()
	assert.Nil(t, err)
	redis.Reset()

	redis.AddTotalCount(10)
	val = redis.GetTotalCount()
	assert.Equal(t, int64(10), val)
	redis.AddTotalCount(10)
	val = redis.GetTotalCount()
	assert.Equal(t, int64(20), val)

	redis.AddIosSuccess(20)
	val = redis.GetIosSuccess()
	assert.Equal(t, int64(20), val)

	redis.AddIosError(30)
	val = redis.GetIosError()
	assert.Equal(t, int64(30), val)

	redis.AddAndroidSuccess(40)
	val = redis.GetAndroidSuccess()
	assert.Equal(t, int64(40), val)

	redis.AddAndroidError(50)
	val = redis.GetAndroidError()
	assert.Equal(t, int64(50), val)

	// test reset db
	redis.Reset()
	val = redis.GetAndroidError()
	assert.Equal(t, int64(0), val)
}
