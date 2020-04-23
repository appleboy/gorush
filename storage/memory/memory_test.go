package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemoryEngine(t *testing.T) {
	var val int64

	memory := New()

	assert.Nil(t, memory.Init())

	memory.AddTotalCount(1)
	val = memory.GetTotalCount()
	assert.Equal(t, int64(1), val)

	memory.AddTotalCount(100)
	val = memory.GetTotalCount()
	assert.Equal(t, int64(101), val)

	memory.AddIosSuccess(2)
	val = memory.GetIosSuccess()
	assert.Equal(t, int64(2), val)

	memory.AddIosError(3)
	val = memory.GetIosError()
	assert.Equal(t, int64(3), val)

	memory.AddAndroidSuccess(4)
	val = memory.GetAndroidSuccess()
	assert.Equal(t, int64(4), val)

	memory.AddAndroidError(5)
	val = memory.GetAndroidError()
	assert.Equal(t, int64(5), val)

	// test reset db
	memory.Reset()
	val = memory.GetTotalCount()
	assert.Equal(t, int64(0), val)

	assert.NoError(t, memory.Close())
}
