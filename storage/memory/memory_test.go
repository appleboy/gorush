package memory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemoryEngine(t *testing.T) {
	var val int64

	memory := New()

	assert.Nil(t, memory.Init())

	memory.AddTotalCount(1)
	val = memory.GetTotalCount()
	assert.Equal(t, int64(1), val)

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
}
