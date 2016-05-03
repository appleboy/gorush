package redis

import (
	c "github.com/appleboy/gorush/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedisServerError(t *testing.T) {
	config := c.BuildDefaultPushConf()
	config.Stat.Redis.Addr = "localhost:6370"

	redis := New(config)
	err := redis.Init()

	assert.Error(t, err)
}

func TestRedisEngine(t *testing.T) {
	var val int64

	config := c.BuildDefaultPushConf()

	redis := New(config)
	redis.Init()
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
}
