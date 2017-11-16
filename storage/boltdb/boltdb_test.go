package boltdb

import (
	"testing"

	c "github.com/jaraxasoftware/gorush/config"
	"github.com/stretchr/testify/assert"
)

func TestBoltDBEngine(t *testing.T) {
	var val int64

	config, _ := c.LoadConf("")

	boltDB := New(config)
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

	boltDB.AddWebSuccess(60)
	val = boltDB.GetWebSuccess()
	assert.Equal(t, int64(60), val)

	boltDB.AddWebError(70)
	val = boltDB.GetWebError()
	assert.Equal(t, int64(70), val)

	// test reset db
	boltDB.Reset()
	val = boltDB.GetAndroidError()
	assert.Equal(t, int64(0), val)
}
