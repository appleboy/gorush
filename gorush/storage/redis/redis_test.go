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

	redis.addTotalCount(1)
	val = redis.getTotalCount()
	assert.Equal(t, int64(1), val)

	redis.addIosSuccess(2)
	val = redis.getIosSuccess()
	assert.Equal(t, int64(2), val)

	redis.addIosError(3)
	val = redis.getIosError()
	assert.Equal(t, int64(3), val)

	redis.addAndroidSuccess(4)
	val = redis.getAndroidSuccess()
	assert.Equal(t, int64(4), val)

	redis.addAndroidError(5)
	val = redis.getAndroidError()
	assert.Equal(t, int64(5), val)
}
