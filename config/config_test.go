package config

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"testing"
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
	defer os.Remove(filename)

	// parse JSON format error
	_, err := LoadConfYaml(filename)

	assert.NotNil(t, err)
}

type ConfigTestSuite struct {
	suite.Suite
	ConfGorushDefault ConfYaml
	ConfGorush        ConfYaml
}

func (suite *ConfigTestSuite) SetupTest() {
	suite.ConfGorushDefault = BuildDefaultPushConf()
	var err error
	suite.ConfGorush, err = LoadConfYaml("config.yml")
	if err != nil {
		panic("failed to load config.yml")
	}
}

func (suite *ConfigTestSuite) TestValidateConfDefault() {
	// Core
	assert.Equal(suite.T(), "8088", suite.ConfGorushDefault.Core.Port)
	assert.Equal(suite.T(), int64(runtime.NumCPU()), suite.ConfGorushDefault.Core.WorkerNum)
	assert.Equal(suite.T(), int64(8192), suite.ConfGorushDefault.Core.QueueNum)
	assert.Equal(suite.T(), "release", suite.ConfGorushDefault.Core.Mode)
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Core.SSL)
	assert.Equal(suite.T(), "cert.pem", suite.ConfGorushDefault.Core.CertPath)
	assert.Equal(suite.T(), "key.pem", suite.ConfGorushDefault.Core.KeyPath)
	assert.Equal(suite.T(), int64(100), suite.ConfGorushDefault.Core.MaxNotification)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Core.HTTPProxy)
	// Pid
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Core.PID.Enabled)
	assert.Equal(suite.T(), "gorush.pid", suite.ConfGorushDefault.Core.PID.Path)
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Core.PID.Override)

	// Api
	assert.Equal(suite.T(), "/api/push", suite.ConfGorushDefault.API.PushURI)
	assert.Equal(suite.T(), "/api/stat/go", suite.ConfGorushDefault.API.StatGoURI)
	assert.Equal(suite.T(), "/api/stat/app", suite.ConfGorushDefault.API.StatAppURI)
	assert.Equal(suite.T(), "/api/config", suite.ConfGorushDefault.API.ConfigURI)
	assert.Equal(suite.T(), "/sys/stats", suite.ConfGorushDefault.API.SysStatURI)

	// Android
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Android.Enabled)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Android.APIKey)

	// iOS
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Ios.Enabled)
	assert.Equal(suite.T(), "key.pem", suite.ConfGorushDefault.Ios.KeyPath)
	assert.Equal(suite.T(), "", suite.ConfGorushDefault.Ios.Password)
	assert.Equal(suite.T(), false, suite.ConfGorushDefault.Ios.Production)

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

	assert.Equal(suite.T(), "gorush.db", suite.ConfGorushDefault.Stat.BoltDB.Path)
	assert.Equal(suite.T(), "gorush", suite.ConfGorushDefault.Stat.BoltDB.Bucket)

	assert.Equal(suite.T(), "gorush.db", suite.ConfGorushDefault.Stat.BuntDB.Path)
	assert.Equal(suite.T(), "gorush.db", suite.ConfGorushDefault.Stat.LevelDB.Path)
}

func (suite *ConfigTestSuite) TestValidateConf() {
	// Core
	assert.Equal(suite.T(), "8088", suite.ConfGorush.Core.Port)
	assert.Equal(suite.T(), int64(8), suite.ConfGorush.Core.WorkerNum)
	assert.Equal(suite.T(), int64(8192), suite.ConfGorush.Core.QueueNum)
	assert.Equal(suite.T(), "release", suite.ConfGorush.Core.Mode)
	assert.Equal(suite.T(), false, suite.ConfGorush.Core.SSL)
	assert.Equal(suite.T(), "cert.pem", suite.ConfGorush.Core.CertPath)
	assert.Equal(suite.T(), "key.pem", suite.ConfGorush.Core.KeyPath)
	assert.Equal(suite.T(), int64(100), suite.ConfGorush.Core.MaxNotification)
	assert.Equal(suite.T(), "", suite.ConfGorush.Core.HTTPProxy)
	// Pid
	assert.Equal(suite.T(), false, suite.ConfGorush.Core.PID.Enabled)
	assert.Equal(suite.T(), "gorush.pid", suite.ConfGorush.Core.PID.Path)
	assert.Equal(suite.T(), true, suite.ConfGorush.Core.PID.Override)

	// Api
	assert.Equal(suite.T(), "/api/push", suite.ConfGorush.API.PushURI)
	assert.Equal(suite.T(), "/api/stat/go", suite.ConfGorush.API.StatGoURI)
	assert.Equal(suite.T(), "/api/stat/app", suite.ConfGorush.API.StatAppURI)
	assert.Equal(suite.T(), "/api/config", suite.ConfGorush.API.ConfigURI)
	assert.Equal(suite.T(), "/sys/stats", suite.ConfGorush.API.SysStatURI)

	// Android
	assert.Equal(suite.T(), true, suite.ConfGorush.Android.Enabled)
	assert.Equal(suite.T(), "YOUR_API_KEY", suite.ConfGorush.Android.APIKey)

	// iOS
	assert.Equal(suite.T(), false, suite.ConfGorush.Ios.Enabled)
	assert.Equal(suite.T(), "key.pem", suite.ConfGorush.Ios.KeyPath)
	assert.Equal(suite.T(), "", suite.ConfGorush.Ios.Password)
	assert.Equal(suite.T(), false, suite.ConfGorush.Ios.Production)

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

	assert.Equal(suite.T(), "gorush.db", suite.ConfGorush.Stat.BoltDB.Path)
	assert.Equal(suite.T(), "gorush", suite.ConfGorush.Stat.BoltDB.Bucket)

	assert.Equal(suite.T(), "gorush.db", suite.ConfGorush.Stat.BuntDB.Path)
	assert.Equal(suite.T(), "gorush.db", suite.ConfGorush.Stat.LevelDB.Path)
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
