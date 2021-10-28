package boltdb

import (
	"testing"

	"github.com/appleboy/gorush/config"
	"github.com/stretchr/testify/assert"
)

func TestBoltDBEngine(t *testing.T) {
	var val int64

	cfg, _ := config.LoadConf()

	boltDB := New(cfg)
	err := boltDB.Init()
	assert.Nil(t, err)
	boltDB.Reset()

	boltDB.AddTotalCount(10)
	val = boltDB.GetTotalCount()
	assert.Equal(t, int64(10), val)
	boltDB.AddTotalCount(10)
	val = boltDB.GetTotalCount()
	assert.Equal(t, int64(20), val)

	boltDB.AddIosSuccess(20)
	val = boltDB.GetIosSuccess()
	assert.Equal(t, int64(20), val)

	boltDB.AddIosError(30)
	val = boltDB.GetIosError()
	assert.Equal(t, int64(30), val)

	boltDB.AddAndroidSuccess(40)
	val = boltDB.GetAndroidSuccess()
	assert.Equal(t, int64(40), val)

	boltDB.AddAndroidError(50)
	val = boltDB.GetAndroidError()
	assert.Equal(t, int64(50), val)

	// test reset db
	boltDB.Reset()
	val = boltDB.GetAndroidError()
	assert.Equal(t, int64(0), val)

	assert.NoError(t, boltDB.Close())
}
