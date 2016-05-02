package redis

import (
	"github.com/appleboy/gorush/gorush"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedisServerError(t *testing.T) {
	config := gorush.BuildDefaultPushConf()
	config.Stat.Redis.Addr = "localhost:6370"

	redis := New(config, gorush.StatusApp{})
	err := redis.initRedis()

	assert.Error(t, err)
}

func TestRedisEngine(t *testing.T) {
	var val int64

	config := gorush.BuildDefaultPushConf()

	redis := New(config, gorush.StatusApp{})
	redis.initRedis()
	redis.resetRedis()

	redis.addTotalCount(10)
	val = redis.getTotalCount()
	assert.Equal(t, int64(10), val)
	redis.addTotalCount(10)
	val = redis.getTotalCount()
	assert.Equal(t, int64(20), val)

	redis.addIosSuccess(20)
	val = redis.getIosSuccess()
	assert.Equal(t, int64(20), val)

	redis.addIosError(30)
	val = redis.getIosError()
	assert.Equal(t, int64(30), val)

	redis.addAndroidSuccess(40)
	val = redis.getAndroidSuccess()
	assert.Equal(t, int64(40), val)

	redis.addAndroidError(50)
	val = redis.getAndroidError()
	assert.Equal(t, int64(50), val)
}
