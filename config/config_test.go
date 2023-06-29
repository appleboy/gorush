package config

import (
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Test file is missing
func TestMissingFile(t *testing.T) {
	filename := "test"
	_, err := LoadConf(filename)

	assert.NotNil(t, err)
}

type ConfigTestSuite struct {
	suite.Suite
	ConfGorushDefault *ConfYaml
	ConfGorush        *ConfYaml
}

func (suite *ConfigTestSuite) SetupTest() {
	var err error
	suite.ConfGorushDefault, err = LoadConf()
	if err != nil {
		panic("failed to load default config.yml")
	}
	suite.ConfGorush, err = LoadConf("testdata/config.yml")
	if err != nil {
		panic("failed to load config.yml from file")
	}
}

func (suite *ConfigTestSuite) TestValidateConfDefault() {
	// Core
	assert.Equal(suite.T(), "", suite.ConfGorush.Core.Address)
	assert.Equal(suite.T(), "8088", suite.ConfGorush.Core.Port)
	assert.Equal(suite.T(), int64(30), suite.ConfGorush.Core.ShutdownTimeout)
	assert.Equal(suite.T(), true, suite.ConfGorush.Core.Enabled)
	assert.Equal(suite.T(), int64(runtime.NumCPU()), suite.ConfGorush.Core.WorkerNum)
	assert.Equal(suite.T(), int64(8192), suite.ConfGorush.Core.QueueNum)
	assert.Equal(suite.T(), "release", suite.ConfGorush.Core.Mode)
	assert.Equal(suite.T(), false, suite.ConfGorush.Core.Sync)
	assert.Equal(suite.T(), "", suite.ConfGorush.Core.FeedbackURL)
	assert.Equal(suite.T(), int64(10), suite.ConfGorush.Core.FeedbackTimeout)
	assert.Equal(suite.T(), false, suite.ConfGorush.Core.SSL)
	assert.Equal(suite.T(), "cert.pem", suite.ConfGorush.Core.CertPath)
	assert.Equal(suite.T(), "key.pem", suite.ConfGorush.Core.KeyPath)
	assert.Equal(suite.T(), "", suite.ConfGorush.Core.KeyBase64)
	assert.Equal(suite.T(), "", suite.ConfGorush.Core.CertBase64)
	assert.Equal(suite.T(), int64(100), suite.ConfGorush.Core.MaxNotification)
	assert.Equal(suite.T(), "", suite.ConfGorush.Core.HTTPProxy)
	// Pid
	assert.Equal(suite.T(), false, suite.ConfGorush.Core.PID.Enabled)
	assert.Equal(suite.T(), "gorush.pid", suite.ConfGorush.Core.PID.Path)
	assert.Equal(suite.T(), true, suite.ConfGorush.Core.PID.Override)
	assert.Equal(suite.T(), false, suite.ConfGorush.Core.AutoTLS.Enabled)
	assert.Equal(suite.T(), ".cache", suite.ConfGorush.Core.AutoTLS.Folder)
	assert.Equal(suite.T(), "", suite.ConfGorush.Core.AutoTLS.Host)

	// Api
	assert.Equal(suite.T(), "/api/stat/go", suite.ConfGorush.API.StatGoURI)
	assert.Equal(suite.T(), "/api/stat/app", suite.ConfGorush.API.StatAppURI)
	assert.Equal(suite.T(), "/api/config", suite.ConfGorush.API.ConfigURI)
	assert.Equal(suite.T(), "/sys/stats", suite.ConfGorush.API.SysStatURI)
	assert.Equal(suite.T(), "/metrics", suite.ConfGorush.API.MetricURI)
	assert.Equal(suite.T(), "/healthz", suite.ConfGorush.API.HealthURI)

	tenant := suite.ConfGorush.Tenants["tenant_id1"]
	assert.Equal(suite.T(), "/api/push/tenant1", tenant.PushURI)
	// Android
	assert.Equal(suite.T(), true, tenant.Android.Enabled)
	assert.Equal(suite.T(), "YOUR_API_KEY", tenant.Android.APIKey)
	assert.Equal(suite.T(), 0, tenant.Android.MaxRetry)

	// iOS
	assert.Equal(suite.T(), false, tenant.Ios.Enabled)
	assert.Equal(suite.T(), "key.pem", tenant.Ios.KeyPath)
	assert.Equal(suite.T(), "", tenant.Ios.KeyBase64)
	assert.Equal(suite.T(), "pem", tenant.Ios.KeyType)
	assert.Equal(suite.T(), "", tenant.Ios.Password)
	assert.Equal(suite.T(), false, tenant.Ios.Production)
	assert.Equal(suite.T(), uint(100), tenant.Ios.MaxConcurrentPushes)
	assert.Equal(suite.T(), 0, tenant.Ios.MaxRetry)
	assert.Equal(suite.T(), "", tenant.Ios.KeyID)
	assert.Equal(suite.T(), "", tenant.Ios.TeamID)

	// queue
	assert.Equal(suite.T(), "local", suite.ConfGorush.Queue.Engine)
	assert.Equal(suite.T(), "127.0.0.1:4150", suite.ConfGorush.Queue.NSQ.Addr)
	assert.Equal(suite.T(), "gorush", suite.ConfGorush.Queue.NSQ.Topic)
	assert.Equal(suite.T(), "gorush", suite.ConfGorush.Queue.NSQ.Channel)

	assert.Equal(suite.T(), "127.0.0.1:4222", suite.ConfGorush.Queue.NATS.Addr)
	assert.Equal(suite.T(), "gorush", suite.ConfGorush.Queue.NATS.Subj)
	assert.Equal(suite.T(), "gorush", suite.ConfGorush.Queue.NATS.Queue)

	assert.Equal(suite.T(), "127.0.0.1:6379", suite.ConfGorush.Queue.Redis.Addr)
	assert.Equal(suite.T(), "gorush", suite.ConfGorush.Queue.Redis.StreamName)
	assert.Equal(suite.T(), "gorush", suite.ConfGorush.Queue.Redis.Group)
	assert.Equal(suite.T(), "gorush", suite.ConfGorush.Queue.Redis.Consumer)

	// log
	assert.Equal(suite.T(), "string", suite.ConfGorush.Log.Format)
	assert.Equal(suite.T(), "stdout", suite.ConfGorush.Log.AccessLog)
	assert.Equal(suite.T(), "debug", suite.ConfGorush.Log.AccessLevel)
	assert.Equal(suite.T(), "stderr", suite.ConfGorush.Log.ErrorLog)
	assert.Equal(suite.T(), "error", suite.ConfGorush.Log.ErrorLevel)
	assert.Equal(suite.T(), true, suite.ConfGorush.Log.HideToken)
	assert.Equal(suite.T(), false, suite.ConfGorush.Log.HideMessages)

	assert.Equal(suite.T(), "memory", suite.ConfGorush.Stat.Engine)
	assert.Equal(suite.T(), false, suite.ConfGorush.Stat.Redis.Cluster)
	assert.Equal(suite.T(), "localhost:6379", suite.ConfGorush.Stat.Redis.Addr)
	assert.Equal(suite.T(), "", suite.ConfGorush.Stat.Redis.Password)
	assert.Equal(suite.T(), 0, suite.ConfGorush.Stat.Redis.DB)

	assert.Equal(suite.T(), "bolt.db", suite.ConfGorush.Stat.BoltDB.Path)
	assert.Equal(suite.T(), "gorush", suite.ConfGorush.Stat.BoltDB.Bucket)

	assert.Equal(suite.T(), "bunt.db", suite.ConfGorush.Stat.BuntDB.Path)
	assert.Equal(suite.T(), "level.db", suite.ConfGorush.Stat.LevelDB.Path)
	assert.Equal(suite.T(), "badger.db", suite.ConfGorush.Stat.BadgerDB.Path)

	// gRPC
	assert.Equal(suite.T(), false, suite.ConfGorush.GRPC.Enabled)
	assert.Equal(suite.T(), "9000", suite.ConfGorush.GRPC.Port)
}

func (suite *ConfigTestSuite) TestValidateConf() {
	// Core
	assert.Equal(suite.T(), "8088", suite.ConfGorush.Core.Port)
	assert.Equal(suite.T(), int64(30), suite.ConfGorush.Core.ShutdownTimeout)
	assert.Equal(suite.T(), true, suite.ConfGorush.Core.Enabled)
	assert.Equal(suite.T(), int64(runtime.NumCPU()), suite.ConfGorush.Core.WorkerNum)
	assert.Equal(suite.T(), int64(8192), suite.ConfGorush.Core.QueueNum)
	assert.Equal(suite.T(), "release", suite.ConfGorush.Core.Mode)
	assert.Equal(suite.T(), false, suite.ConfGorush.Core.Sync)
	assert.Equal(suite.T(), "", suite.ConfGorush.Core.FeedbackURL)
	assert.Equal(suite.T(), int64(10), suite.ConfGorush.Core.FeedbackTimeout)
	assert.Equal(suite.T(), false, suite.ConfGorush.Core.SSL)
	assert.Equal(suite.T(), "cert.pem", suite.ConfGorush.Core.CertPath)
	assert.Equal(suite.T(), "key.pem", suite.ConfGorush.Core.KeyPath)
	assert.Equal(suite.T(), "", suite.ConfGorush.Core.CertBase64)
	assert.Equal(suite.T(), "", suite.ConfGorush.Core.KeyBase64)
	assert.Equal(suite.T(), int64(100), suite.ConfGorush.Core.MaxNotification)
	assert.Equal(suite.T(), "", suite.ConfGorush.Core.HTTPProxy)
	// Pid
	assert.Equal(suite.T(), false, suite.ConfGorush.Core.PID.Enabled)
	assert.Equal(suite.T(), "gorush.pid", suite.ConfGorush.Core.PID.Path)
	assert.Equal(suite.T(), true, suite.ConfGorush.Core.PID.Override)
	assert.Equal(suite.T(), false, suite.ConfGorush.Core.AutoTLS.Enabled)
	assert.Equal(suite.T(), ".cache", suite.ConfGorush.Core.AutoTLS.Folder)
	assert.Equal(suite.T(), "", suite.ConfGorush.Core.AutoTLS.Host)

	// Api
	assert.Equal(suite.T(), "/api/stat/go", suite.ConfGorush.API.StatGoURI)
	assert.Equal(suite.T(), "/api/stat/app", suite.ConfGorush.API.StatAppURI)
	assert.Equal(suite.T(), "/api/config", suite.ConfGorush.API.ConfigURI)
	assert.Equal(suite.T(), "/sys/stats", suite.ConfGorush.API.SysStatURI)
	assert.Equal(suite.T(), "/metrics", suite.ConfGorush.API.MetricURI)
	assert.Equal(suite.T(), "/healthz", suite.ConfGorush.API.HealthURI)

	// Android
	assert.Equal(suite.T(), true, suite.ConfGorush.Tenants["tenant_id1"].Android.Enabled)
	assert.Equal(suite.T(), "YOUR_API_KEY", suite.ConfGorush.Tenants["tenant_id1"].Android.APIKey)
	assert.Equal(suite.T(), 0, suite.ConfGorush.Tenants["tenant_id1"].Android.MaxRetry)

	// iOS
	assert.Equal(suite.T(), false, suite.ConfGorush.Tenants["tenant_id1"].Ios.Enabled)
	assert.Equal(suite.T(), "key.pem", suite.ConfGorush.Tenants["tenant_id1"].Ios.KeyPath)
	assert.Equal(suite.T(), "", suite.ConfGorush.Tenants["tenant_id1"].Ios.KeyBase64)
	assert.Equal(suite.T(), "pem", suite.ConfGorush.Tenants["tenant_id1"].Ios.KeyType)
	assert.Equal(suite.T(), "", suite.ConfGorush.Tenants["tenant_id1"].Ios.Password)
	assert.Equal(suite.T(), false, suite.ConfGorush.Tenants["tenant_id1"].Ios.Production)
	assert.Equal(suite.T(), uint(100), suite.ConfGorush.Tenants["tenant_id1"].Ios.MaxConcurrentPushes)
	assert.Equal(suite.T(), 0, suite.ConfGorush.Tenants["tenant_id1"].Ios.MaxRetry)
	assert.Equal(suite.T(), "", suite.ConfGorush.Tenants["tenant_id1"].Ios.KeyID)
	assert.Equal(suite.T(), "", suite.ConfGorush.Tenants["tenant_id1"].Ios.TeamID)

	// log
	assert.Equal(suite.T(), "string", suite.ConfGorush.Log.Format)
	assert.Equal(suite.T(), "stdout", suite.ConfGorush.Log.AccessLog)
	assert.Equal(suite.T(), "debug", suite.ConfGorush.Log.AccessLevel)
	assert.Equal(suite.T(), "stderr", suite.ConfGorush.Log.ErrorLog)
	assert.Equal(suite.T(), "error", suite.ConfGorush.Log.ErrorLevel)
	assert.Equal(suite.T(), true, suite.ConfGorush.Log.HideToken)

	assert.Equal(suite.T(), "memory", suite.ConfGorush.Stat.Engine)
	assert.Equal(suite.T(), false, suite.ConfGorush.Stat.Redis.Cluster)
	assert.Equal(suite.T(), "localhost:6379", suite.ConfGorush.Stat.Redis.Addr)
	assert.Equal(suite.T(), "", suite.ConfGorush.Stat.Redis.Password)
	assert.Equal(suite.T(), 0, suite.ConfGorush.Stat.Redis.DB)

	assert.Equal(suite.T(), "bolt.db", suite.ConfGorush.Stat.BoltDB.Path)
	assert.Equal(suite.T(), "gorush", suite.ConfGorush.Stat.BoltDB.Bucket)

	assert.Equal(suite.T(), "bunt.db", suite.ConfGorush.Stat.BuntDB.Path)
	assert.Equal(suite.T(), "level.db", suite.ConfGorush.Stat.LevelDB.Path)
	assert.Equal(suite.T(), "badger.db", suite.ConfGorush.Stat.BadgerDB.Path)

	// gRPC
	assert.Equal(suite.T(), false, suite.ConfGorush.GRPC.Enabled)
	assert.Equal(suite.T(), "9000", suite.ConfGorush.GRPC.Port)
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

func TestLoadConfigFromEnv(t *testing.T) {
	os.Setenv("GORUSH_CORE_PORT", "9001")
	os.Setenv("GORUSH_GRPC_ENABLED", "true")
	os.Setenv("GORUSH_CORE_MAX_NOTIFICATION", "200")
	os.Setenv("GORUSH_API_HEALTH_URI", "/healthz")
	ConfGorush, err := LoadConf("testdata/config.yml")
	if err != nil {
		panic("failed to load config.yml from file")
	}
	assert.Equal(t, "9001", ConfGorush.Core.Port)
	assert.Equal(t, int64(200), ConfGorush.Core.MaxNotification)
	assert.True(t, ConfGorush.GRPC.Enabled)
	assert.Equal(t, "/healthz", ConfGorush.API.HealthURI)
}
