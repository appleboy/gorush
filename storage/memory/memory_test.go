package memory

import (
	"github.com/appleboy/gorush/gorush"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemoryEngine(t *testing.T) {
	var val int64

	memory := New(gorush.StatusApp{})

	memory.addTotalCount(1)
	val = memory.getTotalCount()
	assert.Equal(t, int64(1), val)

	memory.addIosSuccess(2)
	val = memory.getIosSuccess()
	assert.Equal(t, int64(2), val)

	memory.addIosError(3)
	val = memory.getIosError()
	assert.Equal(t, int64(3), val)

	memory.addAndroidSuccess(4)
	val = memory.getAndroidSuccess()
	assert.Equal(t, int64(4), val)

	memory.addAndroidError(5)
	val = memory.getAndroidError()
	assert.Equal(t, int64(5), val)
}
