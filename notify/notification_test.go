package notify

import (
	"testing"

	"github.com/appleboy/gorush/config"

	"github.com/stretchr/testify/assert"
)

const (
	testHuaweiAppID     = "app-id"
	testHuaweiAppSecret = "app-secret"
)

func TestCorrectConf(t *testing.T) {
	cfg, _ := config.LoadConf()

	cfg.Android.Enabled = true
	cfg.Android.Credential = "xxxxx"

	cfg.Ios.Enabled = true
	cfg.Ios.KeyPath = testKeyPath

	err := CheckPushConf(cfg)

	assert.NoError(t, err)
}

func TestSetProxyURL(t *testing.T) {
	err := SetProxy("87.236.233.92:8080")
	assert.Error(t, err)
	assert.Equal(t, "parse \"87.236.233.92:8080\": invalid URI for request", err.Error())

	err = SetProxy("a.html")
	assert.Error(t, err)

	err = SetProxy("http://87.236.233.92:8080")
	assert.NoError(t, err)
}

// Tests for refactored helper functions

func TestCheckIOSConf(t *testing.T) {
	cfg, _ := config.LoadConf()

	// iOS disabled - should pass
	cfg.Ios.Enabled = false
	err := checkIOSConf(cfg)
	assert.NoError(t, err)

	// iOS enabled but missing key path and base64
	cfg.Ios.Enabled = true
	cfg.Ios.KeyPath = ""
	cfg.Ios.KeyBase64 = ""
	err = checkIOSConf(cfg)
	assert.Error(t, err)
	assert.Equal(t, "missing iOS certificate key", err.Error())

	// iOS enabled with valid key path
	cfg.Ios.KeyPath = testKeyPath
	err = checkIOSConf(cfg)
	assert.NoError(t, err)

	// iOS enabled with key base64 (no key path)
	cfg.Ios.KeyPath = ""
	cfg.Ios.KeyBase64 = "some-base64-data"
	err = checkIOSConf(cfg)
	assert.NoError(t, err)

	// iOS enabled with non-existent file
	cfg.Ios.KeyPath = "non-existent-file.pem"
	cfg.Ios.KeyBase64 = ""
	err = checkIOSConf(cfg)
	assert.Error(t, err)
	assert.Equal(t, "certificate file does not exist", err.Error())
}

func TestCheckAndroidConf(t *testing.T) {
	cfg, _ := config.LoadConf()

	// Android disabled - should pass
	cfg.Android.Enabled = false
	err := checkAndroidConf(cfg)
	assert.NoError(t, err)

	// Android enabled but no credentials
	cfg.Android.Enabled = true
	cfg.Android.Credential = ""
	cfg.Android.KeyPath = ""
	// Clear environment variable for this test
	t.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "")
	err = checkAndroidConf(cfg)
	assert.Error(t, err)
	assert.Equal(t, "missing fcm credential data", err.Error())

	// Android enabled with credential
	cfg.Android.Credential = "some-credential"
	err = checkAndroidConf(cfg)
	assert.NoError(t, err)

	// Android enabled with key path
	cfg.Android.Credential = ""
	cfg.Android.KeyPath = "/path/to/key.json"
	err = checkAndroidConf(cfg)
	assert.NoError(t, err)
}

func TestCheckHuaweiConf(t *testing.T) {
	cfg, _ := config.LoadConf()

	// Huawei disabled - should pass
	cfg.Huawei.Enabled = false
	err := checkHuaweiConf(cfg)
	assert.NoError(t, err)

	// Huawei enabled but missing app secret
	cfg.Huawei.Enabled = true
	cfg.Huawei.AppSecret = ""
	cfg.Huawei.AppID = testHuaweiAppID
	err = checkHuaweiConf(cfg)
	assert.Error(t, err)
	assert.Equal(t, "missing huawei app secret", err.Error())

	// Huawei enabled but missing app id
	cfg.Huawei.AppSecret = testHuaweiAppSecret
	cfg.Huawei.AppID = ""
	err = checkHuaweiConf(cfg)
	assert.Error(t, err)
	assert.Equal(t, "missing huawei app id", err.Error())

	// Huawei enabled with all credentials
	cfg.Huawei.AppSecret = testHuaweiAppSecret
	cfg.Huawei.AppID = testHuaweiAppID
	err = checkHuaweiConf(cfg)
	assert.NoError(t, err)
}

func TestCheckPushConfNoPlatformEnabled(t *testing.T) {
	cfg, _ := config.LoadConf()

	cfg.Ios.Enabled = false
	cfg.Android.Enabled = false
	cfg.Huawei.Enabled = false

	err := CheckPushConf(cfg)
	assert.Error(t, err)
	assert.Equal(t, "please enable iOS, Android or Huawei config in yml config", err.Error())
}

func TestCheckPushConfAllPlatformsValid(t *testing.T) {
	cfg, _ := config.LoadConf()

	cfg.Ios.Enabled = true
	cfg.Ios.KeyPath = testKeyPath

	cfg.Android.Enabled = true
	cfg.Android.Credential = "some-credential"

	cfg.Huawei.Enabled = true
	cfg.Huawei.AppSecret = testHuaweiAppSecret
	cfg.Huawei.AppID = testHuaweiAppID

	err := CheckPushConf(cfg)
	assert.NoError(t, err)
}
