package buntdb

import (
	"os"
	"testing"

	c "github.com/jaraxasoftware/gorush/config"
	"github.com/stretchr/testify/assert"
)

func TestBuntDBEngine(t *testing.T) {
	var val int64

	config, _ := c.LoadConf("")

	if _, err := os.Stat(config.Stat.BuntDB.Path); os.IsNotExist(err) {
		err := os.RemoveAll(config.Stat.BuntDB.Path)
		assert.Nil(t, err)
	}

	buntDB := New(config)
	err := buntDB.Init()
	assert.Nil(t, err)
	buntDB.Reset()

	buntDB.AddTotalCount(10)
	val = buntDB.GetTotalCount()
	assert.Equal(t, int64(10), val)
	buntDB.AddTotalCount(10)
	val = buntDB.GetTotalCount()
	assert.Equal(t, int64(20), val)

	buntDB.AddIosSuccess(20)
	val = buntDB.GetIosSuccess()
	assert.Equal(t, int64(20), val)

	buntDB.AddIosError(30)
	val = buntDB.GetIosError()
	assert.Equal(t, int64(30), val)

	buntDB.AddAndroidSuccess(40)
	val = buntDB.GetAndroidSuccess()
	assert.Equal(t, int64(40), val)

	buntDB.AddAndroidError(50)
	val = buntDB.GetAndroidError()
	assert.Equal(t, int64(50), val)

	buntDB.AddWebSuccess(60)
	val = buntDB.GetWebSuccess()
	assert.Equal(t, int64(60), val)

	buntDB.AddWebError(70)
	val = buntDB.GetWebError()
	assert.Equal(t, int64(70), val)

	buntDB.Reset()
	val = buntDB.GetAndroidError()
	assert.Equal(t, int64(0), val)
}
