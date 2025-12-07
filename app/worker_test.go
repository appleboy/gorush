package app

import (
	"testing"

	"github.com/appleboy/gorush/config"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewQueueWorker_LocalQueue(t *testing.T) {
	cfg, _ := config.LoadConf("")
	cfg.Queue.Engine = "local"

	w, err := NewQueueWorker(cfg)
	require.NoError(t, err)
	assert.NotNil(t, w)
}

func TestNewQueueWorker_UnsupportedEngine(t *testing.T) {
	cfg, _ := config.LoadConf("")
	cfg.Queue.Engine = "unsupported"

	w, err := NewQueueWorker(cfg)
	assert.Error(t, err)
	assert.Nil(t, w)
	assert.Contains(t, err.Error(), "unsupported queue engine")
}

func TestNewQueuePool(t *testing.T) {
	cfg, _ := config.LoadConf("")
	cfg.Queue.Engine = "local"

	w, err := NewQueueWorker(cfg)
	require.NoError(t, err)

	q := NewQueuePool(cfg, w)
	assert.NotNil(t, q)

	// Clean up
	q.Release()
}
