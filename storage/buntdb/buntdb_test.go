package buntdb

import (
	"os"
	"testing"

	"github.com/appleboy/gorush/config"
	"github.com/stretchr/testify/assert"
)

func TestBuntDBEngine(t *testing.T) {
	var val int64

	cfg, _ := config.LoadConf()

	if _, err := os.Stat(cfg.Stat.BuntDB.Path); os.IsNotExist(err) {
		err := os.RemoveAll(cfg.Stat.BuntDB.Path)
		assert.Nil(t, err)
	}

	buntDB := New(cfg)
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

	buntDB.Reset()
	val = buntDB.GetAndroidError()
	assert.Equal(t, int64(0), val)

	assert.NoError(t, buntDB.Close())
}
