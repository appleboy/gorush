package gorush

import (
	"testing"

	"github.com/appleboy/gorush/config"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestSetLogLevel(t *testing.T) {
	log := logrus.New()

	err := SetLogLevel(log, "debug")
	assert.Nil(t, err)

	err = SetLogLevel(log, "invalid")
	assert.Equal(t, "not a valid logrus Level: \"invalid\"", err.Error())
}

func TestSetLogOut(t *testing.T) {
	log := logrus.New()

	err := SetLogOut(log, "stdout")
	assert.Nil(t, err)

	err = SetLogOut(log, "stderr")
	assert.Nil(t, err)

	err = SetLogOut(log, "log/access.log")
	assert.Nil(t, err)

	// missing create logs folder.
	err = SetLogOut(log, "logs/access.log")
	assert.NotNil(t, err)
}

func TestInitDefaultLog(t *testing.T) {
	PushConf, _ = config.LoadConf("")

	// no errors on default config
	assert.Nil(t, InitLog())

	PushConf.Log.AccessLevel = "invalid"

	assert.NotNil(t, InitLog())
}

func TestAccessLevel(t *testing.T) {
	PushConf, _ = config.LoadConf("")

	PushConf.Log.AccessLevel = "invalid"

	assert.NotNil(t, InitLog())
}

func TestErrorLevel(t *testing.T) {
	PushConf, _ = config.LoadConf("")

	PushConf.Log.ErrorLevel = "invalid"

	assert.NotNil(t, InitLog())
}

func TestAccessLogPath(t *testing.T) {
	PushConf, _ = config.LoadConf("")

	PushConf.Log.AccessLog = "logs/access.log"

	assert.NotNil(t, InitLog())
}

func TestErrorLogPath(t *testing.T) {
	PushConf, _ = config.LoadConf("")

	PushConf.Log.ErrorLog = "logs/error.log"

	assert.NotNil(t, InitLog())
}

func TestPlatformType(t *testing.T) {
	assert.Equal(t, "ios", typeForPlatform(PlatformIos))
	assert.Equal(t, "android", typeForPlatform(PlatformAndroid))
	assert.Equal(t, "", typeForPlatform(10000))
}

func TestPlatformColor(t *testing.T) {
	assert.Equal(t, blue, colorForPlatform(PlatformIos))
	assert.Equal(t, yellow, colorForPlatform(PlatformAndroid))
	assert.Equal(t, reset, colorForPlatform(1000000))
}

func TestHideToken(t *testing.T) {
	assert.Equal(t, "", hideToken("", 2))
	assert.Equal(t, "**345678**", hideToken("1234567890", 2))
	assert.Equal(t, "*****", hideToken("12345", 10))
}
