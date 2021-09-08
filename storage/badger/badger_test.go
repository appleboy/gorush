package badger

import (
	"testing"

	"github.com/appleboy/gorush/config"
	"github.com/stretchr/testify/assert"
)

func TestBadgerEngine(t *testing.T) {
	var val int64

	cfg, _ := config.LoadConf()

	badger := New(cfg)
	err := badger.Init()
	assert.Nil(t, err)
	badger.Reset()

	badger.AddTotalCount(10)
	val = badger.GetTotalCount()
	assert.Equal(t, int64(10), val)
	badger.AddTotalCount(10)
	val = badger.GetTotalCount()
	assert.Equal(t, int64(20), val)

	badger.AddIosSuccess(20)
	val = badger.GetIosSuccess()
	assert.Equal(t, int64(20), val)

	badger.AddIosError(30)
	val = badger.GetIosError()
	assert.Equal(t, int64(30), val)

	badger.AddAndroidSuccess(40)
	val = badger.GetAndroidSuccess()
	assert.Equal(t, int64(40), val)

	badger.AddAndroidError(50)
	val = badger.GetAndroidError()
	assert.Equal(t, int64(50), val)

	// test reset db
	badger.Reset()
	val = badger.GetAndroidError()
	assert.Equal(t, int64(0), val)

	assert.NoError(t, badger.Close())
}
