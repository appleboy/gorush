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

func TestEmptyConfig(t *testing.T) {
	conf, err := LoadConf("testdata/empty.yml")
	if err != nil {
		panic("failed to load config.yml from file")
	}

	assert.Equal(t, uint(100), conf.Ios.MaxConcurrentPushes)
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
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Core.Address)
	assert.Equal(suite.T(), "8088", suite.ConfGorushDefault.Core.Port)
	assert.Equal(suite.T(), int64(30), suite.ConfGorushDefault.Core.ShutdownTimeout)
	assert.Equal(suite.T(), true, suite.ConfGorushDefault.Core.Enabled)
	assert.Equal(suite.T(), int64(runtime.NumCPU()), suite.ConfGorushDefault.Core.WorkerNum)
	assert.Equal(suite.T(), int64(8192), suite.ConfGorushDefault.Core.QueueNum)
	assert.Equal(suite.T(), "release", suite.ConfGorushDefault.Core.Mode)
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Core.Sync)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Core.FeedbackURL)
	assert.Equal(suite.T(), 0, len(suite.ConfGorushDefault.Core.FeedbackHeader))
	assert.Equal(suite.T(), int64(10), suite.ConfGorushDefault.Core.FeedbackTimeout)
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Core.SSL)
	assert.Equal(suite.T(), "cert.pem", suite.ConfGorushDefault.Core.CertPath)
	assert.Equal(suite.T(), "key.pem", suite.ConfGorushDefault.Core.KeyPath)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Core.KeyBase64)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Core.CertBase64)
	assert.Equal(suite.T(), int64(100), suite.ConfGorushDefault.Core.MaxNotification)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Core.HTTPProxy)
	// Pid
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Core.PID.Enabled)
	assert.Equal(suite.T(), "gorush.pid", suite.ConfGorushDefault.Core.PID.Path)
	assert.Equal(suite.T(), true, suite.ConfGorushDefault.Core.PID.Override)
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Core.AutoTLS.Enabled)
	assert.Equal(suite.T(), ".cache", suite.ConfGorushDefault.Core.AutoTLS.Folder)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Core.AutoTLS.Host)

	// Api
	assert.Equal(suite.T(), "/api/push", suite.ConfGorushDefault.API.PushURI)
	assert.Equal(suite.T(), "/api/stat/go", suite.ConfGorushDefault.API.StatGoURI)
	assert.Equal(suite.T(), "/api/stat/app", suite.ConfGorushDefault.API.StatAppURI)
	assert.Equal(suite.T(), "/api/config", suite.ConfGorushDefault.API.ConfigURI)
	assert.Equal(suite.T(), "/sys/stats", suite.ConfGorushDefault.API.SysStatURI)
	assert.Equal(suite.T(), "/metrics", suite.ConfGorushDefault.API.MetricURI)
	assert.Equal(suite.T(), "/healthz", suite.ConfGorushDefault.API.HealthURI)

	// Android
	assert.Equal(suite.T(), true, suite.ConfGorushDefault.Android.Enabled)
	assert.Equal(suite.T(), "YOUR_API_KEY", suite.ConfGorushDefault.Android.APIKey)
	assert.Equal(suite.T(), 0, suite.ConfGorushDefault.Android.MaxRetry)

	// iOS
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Ios.Enabled)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Ios.KeyPath)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Ios.KeyBase64)
	assert.Equal(suite.T(), "pem", suite.ConfGorushDefault.Ios.KeyType)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Ios.Password)
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Ios.Production)
	assert.Equal(suite.T(), uint(100), suite.ConfGorushDefault.Ios.MaxConcurrentPushes)
	assert.Equal(suite.T(), 0, suite.ConfGorushDefault.Ios.MaxRetry)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Ios.KeyID)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Ios.TeamID)

	// queue
	assert.Equal(suite.T(), "local", suite.ConfGorushDefault.Queue.Engine)
	assert.Equal(suite.T(), "127.0.0.1:4150", suite.ConfGorushDefault.Queue.NSQ.Addr)
	assert.Equal(suite.T(), "gorush", suite.ConfGorushDefault.Queue.NSQ.Topic)
	assert.Equal(suite.T(), "gorush", suite.ConfGorushDefault.Queue.NSQ.Channel)

	assert.Equal(suite.T(), "127.0.0.1:4222", suite.ConfGorushDefault.Queue.NATS.Addr)
	assert.Equal(suite.T(), "gorush", suite.ConfGorushDefault.Queue.NATS.Subj)
	assert.Equal(suite.T(), "gorush", suite.ConfGorushDefault.Queue.NATS.Queue)

	assert.Equal(suite.T(), "127.0.0.1:6379", suite.ConfGorushDefault.Queue.Redis.Addr)
	assert.Equal(suite.T(), "gorush", suite.ConfGorushDefault.Queue.Redis.StreamName)
	assert.Equal(suite.T(), "gorush", suite.ConfGorushDefault.Queue.Redis.Group)
	assert.Equal(suite.T(), "gorush", suite.ConfGorushDefault.Queue.Redis.Consumer)

	// log
	assert.Equal(suite.T(), "string", suite.ConfGorushDefault.Log.Format)
	assert.Equal(suite.T(), "stdout", suite.ConfGorushDefault.Log.AccessLog)
	assert.Equal(suite.T(), "debug", suite.ConfGorushDefault.Log.AccessLevel)
	assert.Equal(suite.T(), "stderr", suite.ConfGorushDefault.Log.ErrorLog)
	assert.Equal(suite.T(), "error", suite.ConfGorushDefault.Log.ErrorLevel)
	assert.Equal(suite.T(), true, suite.ConfGorushDefault.Log.HideToken)
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Log.HideMessages)

	assert.Equal(suite.T(), "memory", suite.ConfGorushDefault.Stat.Engine)
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Stat.Redis.Cluster)
	assert.Equal(suite.T(), "localhost:6379", suite.ConfGorushDefault.Stat.Redis.Addr)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Stat.Redis.Password)
	assert.Equal(suite.T(), 0, suite.ConfGorushDefault.Stat.Redis.DB)

	assert.Equal(suite.T(), "bolt.db", suite.ConfGorushDefault.Stat.BoltDB.Path)
	assert.Equal(suite.T(), "gorush", suite.ConfGorushDefault.Stat.BoltDB.Bucket)

	assert.Equal(suite.T(), "bunt.db", suite.ConfGorushDefault.Stat.BuntDB.Path)
	assert.Equal(suite.T(), "level.db", suite.ConfGorushDefault.Stat.LevelDB.Path)
	assert.Equal(suite.T(), "badger.db", suite.ConfGorushDefault.Stat.BadgerDB.Path)

	// gRPC
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.GRPC.Enabled)
	assert.Equal(suite.T(), "9000", suite.ConfGorushDefault.GRPC.Port)
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
	assert.Equal(suite.T(), 1, len(suite.ConfGorush.Core.FeedbackHeader))
	assert.Equal(suite.T(), "x-gorush-token:4e989115e09680f44a645519fed6a976", suite.ConfGorush.Core.FeedbackHeader[0])
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
	assert.Equal(suite.T(), "/api/push", suite.ConfGorush.API.PushURI)
	assert.Equal(suite.T(), "/api/stat/go", suite.ConfGorush.API.StatGoURI)
	assert.Equal(suite.T(), "/api/stat/app", suite.ConfGorush.API.StatAppURI)
	assert.Equal(suite.T(), "/api/config", suite.ConfGorush.API.ConfigURI)
	assert.Equal(suite.T(), "/sys/stats", suite.ConfGorush.API.SysStatURI)
	assert.Equal(suite.T(), "/metrics", suite.ConfGorush.API.MetricURI)
	assert.Equal(suite.T(), "/healthz", suite.ConfGorush.API.HealthURI)

	// Android
	assert.Equal(suite.T(), true, suite.ConfGorush.Android.Enabled)
	assert.Equal(suite.T(), "YOUR_API_KEY", suite.ConfGorush.Android.APIKey)
	assert.Equal(suite.T(), 0, suite.ConfGorush.Android.MaxRetry)

	// iOS
	assert.Equal(suite.T(), false, suite.ConfGorush.Ios.Enabled)
	assert.Equal(suite.T(), "key.pem", suite.ConfGorush.Ios.KeyPath)
	assert.Equal(suite.T(), "", suite.ConfGorush.Ios.KeyBase64)
	assert.Equal(suite.T(), "pem", suite.ConfGorush.Ios.KeyType)
	assert.Equal(suite.T(), "", suite.ConfGorush.Ios.Password)
	assert.Equal(suite.T(), false, suite.ConfGorush.Ios.Production)
	assert.Equal(suite.T(), uint(100), suite.ConfGorush.Ios.MaxConcurrentPushes)
	assert.Equal(suite.T(), 0, suite.ConfGorush.Ios.MaxRetry)
	assert.Equal(suite.T(), "", suite.ConfGorush.Ios.KeyID)
	assert.Equal(suite.T(), "", suite.ConfGorush.Ios.TeamID)

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
	os.Setenv("GORUSH_IOS_KEY_ID", "ABC123DEFG")
	os.Setenv("GORUSH_IOS_TEAM_ID", "DEF123GHIJ")
	os.Setenv("GORUSH_API_HEALTH_URI", "/healthz")
	os.Setenv("GORUSH_CORE_FEEDBACK_HOOK_URL", "http://example.com")
	os.Setenv("GORUSH_CORE_FEEDBACK_HEADER", "x-api-key:1234567890 x-auth-key:0987654321")
	ConfGorush, err := LoadConf("testdata/config.yml")
	if err != nil {
		panic("failed to load config.yml from file")
	}
	assert.Equal(t, "9001", ConfGorush.Core.Port)
	assert.Equal(t, int64(200), ConfGorush.Core.MaxNotification)
	assert.True(t, ConfGorush.GRPC.Enabled)
	assert.Equal(t, "ABC123DEFG", ConfGorush.Ios.KeyID)
	assert.Equal(t, "DEF123GHIJ", ConfGorush.Ios.TeamID)
	assert.Equal(t, "/healthz", ConfGorush.API.HealthURI)
	assert.Equal(t, "http://example.com", ConfGorush.Core.FeedbackURL)
	assert.Equal(t, "x-api-key:1234567890", ConfGorush.Core.FeedbackHeader[0])
	assert.Equal(t, "x-auth-key:0987654321", ConfGorush.Core.FeedbackHeader[1])
}

func TestLoadWrongDefaultYAMLConfig(t *testing.T) {
	defaultConf = []byte(`a`)
	_, err := LoadConf()
	assert.Error(t, err)
}
