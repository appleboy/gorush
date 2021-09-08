package leveldb

import (
	"os"
	"testing"

	"github.com/appleboy/gorush/config"
	"github.com/stretchr/testify/assert"
)

func TestLevelDBEngine(t *testing.T) {
	var val int64

	cfg, _ := config.LoadConf()

	if _, err := os.Stat(cfg.Stat.LevelDB.Path); os.IsNotExist(err) {
		err = os.RemoveAll(cfg.Stat.LevelDB.Path)
		assert.Nil(t, err)
	}

	levelDB := New(cfg)
	err := levelDB.Init()
	assert.Nil(t, err)
	levelDB.Reset()

	levelDB.AddTotalCount(10)
	val = levelDB.GetTotalCount()
	assert.Equal(t, int64(10), val)
	levelDB.AddTotalCount(10)
	val = levelDB.GetTotalCount()
	assert.Equal(t, int64(20), val)

	levelDB.AddIosSuccess(20)
	val = levelDB.GetIosSuccess()
	assert.Equal(t, int64(20), val)

	levelDB.AddIosError(30)
	val = levelDB.GetIosError()
	assert.Equal(t, int64(30), val)

	levelDB.AddAndroidSuccess(40)
	val = levelDB.GetAndroidSuccess()
	assert.Equal(t, int64(40), val)

	levelDB.AddAndroidError(50)
	val = levelDB.GetAndroidError()
	assert.Equal(t, int64(50), val)

	levelDB.Reset()
	val = levelDB.GetAndroidError()
	assert.Equal(t, int64(0), val)

	assert.NoError(t, levelDB.Close())
}
