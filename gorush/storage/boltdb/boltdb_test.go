package boltdb

import (
	"github.com/appleboy/gorush/gorush"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedisEngine(t *testing.T) {
	var val int64

	config := gorush.BuildDefaultPushConf()

	boltDB := New(config, gorush.StatusApp{})
	boltDB.initBoltDB()
	boltDB.resetBoltDB()

	boltDB.addTotalCount(10)
	val = boltDB.getTotalCount()
	assert.Equal(t, int64(10), val)
	boltDB.addTotalCount(10)
	val = boltDB.getTotalCount()
	assert.Equal(t, int64(20), val)

	boltDB.addIosSuccess(20)
	val = boltDB.getIosSuccess()
	assert.Equal(t, int64(20), val)

	boltDB.addIosError(30)
	val = boltDB.getIosError()
	assert.Equal(t, int64(30), val)

	boltDB.addAndroidSuccess(40)
	val = boltDB.getAndroidSuccess()
	assert.Equal(t, int64(40), val)

	boltDB.addAndroidError(50)
	val = boltDB.getAndroidError()
	assert.Equal(t, int64(50), val)
}
