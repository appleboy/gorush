package logx

import (
	"testing"

	"github.com/appleboy/gorush/config"
	"github.com/appleboy/gorush/core"

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
	PushConf, _ := config.LoadConf("")

	// no errors on default config
	assert.Nil(t, InitLog(
		PushConf.Log.AccessLevel,
		PushConf.Log.AccessLog,
		PushConf.Log.ErrorLevel,
		PushConf.Log.ErrorLog,
	))

	PushConf.Log.AccessLevel = "invalid"

	assert.NotNil(t, InitLog(
		PushConf.Log.AccessLevel,
		PushConf.Log.AccessLog,
		PushConf.Log.ErrorLevel,
		PushConf.Log.ErrorLog,
	))
}

func TestAccessLevel(t *testing.T) {
	PushConf, _ := config.LoadConf("")

	PushConf.Log.AccessLevel = "invalid"

	assert.NotNil(t, InitLog(
		PushConf.Log.AccessLevel,
		PushConf.Log.AccessLog,
		PushConf.Log.ErrorLevel,
		PushConf.Log.ErrorLog,
	))
}

func TestErrorLevel(t *testing.T) {
	PushConf, _ := config.LoadConf("")

	PushConf.Log.ErrorLevel = "invalid"

	assert.NotNil(t, InitLog(
		PushConf.Log.AccessLevel,
		PushConf.Log.AccessLog,
		PushConf.Log.ErrorLevel,
		PushConf.Log.ErrorLog,
	))
}

func TestAccessLogPath(t *testing.T) {
	PushConf, _ := config.LoadConf("")

	PushConf.Log.AccessLog = "logs/access.log"

	assert.NotNil(t, InitLog(
		PushConf.Log.AccessLevel,
		PushConf.Log.AccessLog,
		PushConf.Log.ErrorLevel,
		PushConf.Log.ErrorLog,
	))
}

func TestErrorLogPath(t *testing.T) {
	PushConf, _ := config.LoadConf("")

	PushConf.Log.ErrorLog = "logs/error.log"

	assert.NotNil(t, InitLog(
		PushConf.Log.AccessLevel,
		PushConf.Log.AccessLog,
		PushConf.Log.ErrorLevel,
		PushConf.Log.ErrorLog,
	))
}

func TestPlatFormType(t *testing.T) {
	assert.Equal(t, "ios", typeForPlatForm(core.PlatFormIos))
	assert.Equal(t, "android", typeForPlatForm(core.PlatFormAndroid))
	assert.Equal(t, "", typeForPlatForm(10000))
}

func TestPlatFormColor(t *testing.T) {
	assert.Equal(t, blue, colorForPlatForm(core.PlatFormIos))
	assert.Equal(t, yellow, colorForPlatForm(core.PlatFormAndroid))
	assert.Equal(t, reset, colorForPlatForm(1000000))
}

func TestHideToken(t *testing.T) {
	assert.Equal(t, "", hideToken("", 2))
	assert.Equal(t, "**345678**", hideToken("1234567890", 2))
	assert.Equal(t, "*****", hideToken("12345", 10))
}
