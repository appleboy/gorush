package app

import (
	"testing"

	"github.com/appleboy/gorush/config"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMergeConfig_IOSOptions(t *testing.T) {
	cfg, _ := config.LoadConf("")
	opts := NewOptions()
	opts.Conf.Ios.KeyPath = "/path/to/key"
	opts.Conf.Ios.KeyID = "key-id-123"
	opts.Conf.Ios.TeamID = "team-id-456"
	opts.Conf.Ios.Password = "secret"
	opts.Conf.Ios.Production = true

	err := MergeConfig(cfg, opts)
	require.NoError(t, err)

	assert.Equal(t, "/path/to/key", cfg.Ios.KeyPath)
	assert.Equal(t, "key-id-123", cfg.Ios.KeyID)
	assert.Equal(t, "team-id-456", cfg.Ios.TeamID)
	assert.Equal(t, "secret", cfg.Ios.Password)
	assert.True(t, cfg.Ios.Production)
}

func TestMergeConfig_AndroidOptions(t *testing.T) {
	cfg, _ := config.LoadConf("")
	opts := NewOptions()
	opts.Conf.Android.KeyPath = "/path/to/fcm/key.json"

	err := MergeConfig(cfg, opts)
	require.NoError(t, err)

	assert.Equal(t, "/path/to/fcm/key.json", cfg.Android.KeyPath)
}

func TestMergeConfig_HuaweiOptions(t *testing.T) {
	cfg, _ := config.LoadConf("")
	opts := NewOptions()
	opts.Conf.Huawei.AppSecret = "hms-secret"
	opts.Conf.Huawei.AppID = "hms-id"

	err := MergeConfig(cfg, opts)
	require.NoError(t, err)

	assert.Equal(t, "hms-secret", cfg.Huawei.AppSecret)
	assert.Equal(t, "hms-id", cfg.Huawei.AppID)
}

func TestMergeConfig_StorageOptions(t *testing.T) {
	cfg, _ := config.LoadConf("")
	opts := NewOptions()
	opts.Conf.Stat.Engine = "redis"
	opts.Conf.Stat.Redis.Addr = "localhost:6379"

	err := MergeConfig(cfg, opts)
	require.NoError(t, err)

	assert.Equal(t, "redis", cfg.Stat.Engine)
	assert.Equal(t, "localhost:6379", cfg.Stat.Redis.Addr)
}

func TestMergeConfig_ServerOptions(t *testing.T) {
	cfg, _ := config.LoadConf("")
	opts := NewOptions()
	opts.Conf.Core.Port = "9000"
	opts.Conf.Core.Address = "127.0.0.1"
	opts.Conf.Core.HTTPProxy = "http://proxy:8080"

	err := MergeConfig(cfg, opts)
	require.NoError(t, err)

	assert.Equal(t, "9000", cfg.Core.Port)
	assert.Equal(t, "127.0.0.1", cfg.Core.Address)
	assert.Equal(t, "http://proxy:8080", cfg.Core.HTTPProxy)
}

func TestMergeConfig_InvalidPort(t *testing.T) {
	cfg, _ := config.LoadConf("")
	opts := NewOptions()
	opts.Conf.Core.Port = "invalid"

	err := MergeConfig(cfg, opts)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid port")
}

func TestMergeConfig_NoOverrideEmpty(t *testing.T) {
	cfg, _ := config.LoadConf("")
	originalPort := cfg.Core.Port

	opts := NewOptions()
	// Empty values should not override

	err := MergeConfig(cfg, opts)
	require.NoError(t, err)

	assert.Equal(t, originalPort, cfg.Core.Port)
}

func TestValidateAndMerge(t *testing.T) {
	opts := NewOptions()
	// Use empty config file to get defaults

	cfg, err := ValidateAndMerge(opts)
	require.NoError(t, err)
	assert.NotNil(t, cfg)
}
