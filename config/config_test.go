package config

import (
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Test file is missing
func TestMissingFile(t *testing.T) {
	filename := "test"
	_, err := LoadConfYaml(filename)

	assert.NotNil(t, err)
}

// Test wrong json format
func TestWrongYAMLormat(t *testing.T) {
	content := []byte(`Wrong format`)

	filename := "tempfile"

	if err := ioutil.WriteFile(filename, content, 0644); err != nil {
		log.Fatalf("WriteFile %s: %v", filename, err)
	}

	// clean up
	defer func() {
		err := os.Remove(filename)
		assert.Nil(t, err)
	}()

	// parse JSON format error
	_, err := LoadConfYaml(filename)

	assert.NotNil(t, err)
}

type ConfigTestSuite struct {
	suite.Suite
	ConfGorushDefault ConfYaml
	ConfGorush        ConfYaml
	ConfGorushEnv     ConfYaml
}

func (suite *ConfigTestSuite) SetupTest() {
	suite.ConfGorushDefault = BuildDefaultPushConf()
	var err error
	suite.ConfGorush, err = LoadConfYaml("config.yml")
	if err != nil {
		panic("failed to load config.yml")
	}

	os.Setenv("GORUSH_ENV_TEST_PORT", "9000")
	defer os.Unsetenv("GORUSH_ENV_TEST_PORT")

	suite.ConfGorushEnv, err = LoadConfYaml("config_env_test.yml")
	if err != nil {
		panic("failed to load config_env_test.yml")
	}
}

func (suite *ConfigTestSuite) TestValidateConfDefault() {
	// Core
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Core.Address)
	assert.Equal(suite.T(), "8088", suite.ConfGorushDefault.Core.Port)
	assert.Equal(suite.T(), true, suite.ConfGorushDefault.Core.Enabled)
	assert.Equal(suite.T(), int64(runtime.NumCPU()), suite.ConfGorushDefault.Core.WorkerNum)
	assert.Equal(suite.T(), int64(8192), suite.ConfGorushDefault.Core.QueueNum)
	assert.Equal(suite.T(), "release", suite.ConfGorushDefault.Core.Mode)
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Core.Sync)
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Core.SSL)
	assert.Equal(suite.T(), "cert.pem", suite.ConfGorushDefault.Core.CertPath)
	assert.Equal(suite.T(), "key.pem", suite.ConfGorushDefault.Core.KeyPath)
	assert.Equal(suite.T(), int64(100), suite.ConfGorushDefault.Core.MaxNotification)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Core.HTTPProxy)
	// Pid
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Core.PID.Enabled)
	assert.Equal(suite.T(), "gorush.pid", suite.ConfGorushDefault.Core.PID.Path)
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Core.PID.Override)
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

	// Android
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Android.Enabled)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Android.APIKey)
	assert.Equal(suite.T(), 0, suite.ConfGorushDefault.Android.MaxRetry)

	// iOS
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Ios.Enabled)
	assert.Equal(suite.T(), "key.pem", suite.ConfGorushDefault.Ios.KeyPath)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Ios.Password)
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Ios.Production)
	assert.Equal(suite.T(), 0, suite.ConfGorushDefault.Ios.MaxRetry)

	// log
	assert.Equal(suite.T(), "string", suite.ConfGorushDefault.Log.Format)
	assert.Equal(suite.T(), "stdout", suite.ConfGorushDefault.Log.AccessLog)
	assert.Equal(suite.T(), "debug", suite.ConfGorushDefault.Log.AccessLevel)
	assert.Equal(suite.T(), "stderr", suite.ConfGorushDefault.Log.ErrorLog)
	assert.Equal(suite.T(), "error", suite.ConfGorushDefault.Log.ErrorLevel)
	assert.Equal(suite.T(), true, suite.ConfGorushDefault.Log.HideToken)

	assert.Equal(suite.T(), "memory", suite.ConfGorushDefault.Stat.Engine)
	assert.Equal(suite.T(), "localhost:6379", suite.ConfGorushDefault.Stat.Redis.Addr)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Stat.Redis.Password)
	assert.Equal(suite.T(), 0, suite.ConfGorushDefault.Stat.Redis.DB)

	assert.Equal(suite.T(), "bolt.db", suite.ConfGorushDefault.Stat.BoltDB.Path)
	assert.Equal(suite.T(), "gorush", suite.ConfGorushDefault.Stat.BoltDB.Bucket)

	assert.Equal(suite.T(), "bunt.db", suite.ConfGorushDefault.Stat.BuntDB.Path)
	assert.Equal(suite.T(), "level.db", suite.ConfGorushDefault.Stat.LevelDB.Path)

	// gRPC
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.GRPC.Enabled)
	assert.Equal(suite.T(), "50051", suite.ConfGorushDefault.GRPC.Port)
}

func (suite *ConfigTestSuite) TestValidateConf() {
	// Core
	assert.Equal(suite.T(), "8088", suite.ConfGorush.Core.Port)
	assert.Equal(suite.T(), true, suite.ConfGorush.Core.Enabled)
	assert.Equal(suite.T(), int64(runtime.NumCPU()), suite.ConfGorush.Core.WorkerNum)
	assert.Equal(suite.T(), int64(8192), suite.ConfGorush.Core.QueueNum)
	assert.Equal(suite.T(), "release", suite.ConfGorush.Core.Mode)
	assert.Equal(suite.T(), false, suite.ConfGorush.Core.Sync)
	assert.Equal(suite.T(), false, suite.ConfGorush.Core.SSL)
	assert.Equal(suite.T(), "cert.pem", suite.ConfGorush.Core.CertPath)
	assert.Equal(suite.T(), "key.pem", suite.ConfGorush.Core.KeyPath)
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

	// Android
	assert.Equal(suite.T(), true, suite.ConfGorush.Android.Enabled)
	assert.Equal(suite.T(), "YOUR_API_KEY", suite.ConfGorush.Android.APIKey)
	assert.Equal(suite.T(), 0, suite.ConfGorush.Android.MaxRetry)

	// iOS
	assert.Equal(suite.T(), false, suite.ConfGorush.Ios.Enabled)
	assert.Equal(suite.T(), "key.pem", suite.ConfGorush.Ios.KeyPath)
	assert.Equal(suite.T(), "", suite.ConfGorush.Ios.Password)
	assert.Equal(suite.T(), false, suite.ConfGorush.Ios.Production)
	assert.Equal(suite.T(), 0, suite.ConfGorush.Ios.MaxRetry)

	// log
	assert.Equal(suite.T(), "string", suite.ConfGorush.Log.Format)
	assert.Equal(suite.T(), "stdout", suite.ConfGorush.Log.AccessLog)
	assert.Equal(suite.T(), "debug", suite.ConfGorush.Log.AccessLevel)
	assert.Equal(suite.T(), "stderr", suite.ConfGorush.Log.ErrorLog)
	assert.Equal(suite.T(), "error", suite.ConfGorush.Log.ErrorLevel)
	assert.Equal(suite.T(), true, suite.ConfGorush.Log.HideToken)

	assert.Equal(suite.T(), "memory", suite.ConfGorush.Stat.Engine)
	assert.Equal(suite.T(), "localhost:6379", suite.ConfGorush.Stat.Redis.Addr)
	assert.Equal(suite.T(), "", suite.ConfGorush.Stat.Redis.Password)
	assert.Equal(suite.T(), 0, suite.ConfGorush.Stat.Redis.DB)

	assert.Equal(suite.T(), "bolt.db", suite.ConfGorush.Stat.BoltDB.Path)
	assert.Equal(suite.T(), "gorush", suite.ConfGorush.Stat.BoltDB.Bucket)

	assert.Equal(suite.T(), "bunt.db", suite.ConfGorush.Stat.BuntDB.Path)
	assert.Equal(suite.T(), "level.db", suite.ConfGorush.Stat.LevelDB.Path)

	// gRPC
	assert.Equal(suite.T(), false, suite.ConfGorush.GRPC.Enabled)
	assert.Equal(suite.T(), "50051", suite.ConfGorush.GRPC.Port)
}

func (suite *ConfigTestSuite) TestValidateConfEnvExpansion() {
	assert.Equal(suite.T(), "9000", suite.ConfGorushEnv.Core.Port)
}

func (suite *ConfigTestSuite) TestValidateConfMissingEnvExpansion() {
	assert.Equal(suite.T(), "${GORUSH_ENV_TEST_MISSING}", suite.ConfGorushEnv.Core.Mode)
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
