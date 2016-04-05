package gopush

import (
	"github.com/Sirupsen/logrus"
	"os"
	"errors"
)

func InitLog() error {

	var err error

	// init logger
	LogAccess = logrus.New()
	LogError = logrus.New()

	LogAccess.Formatter = &logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	}

	LogError.Formatter = &logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	}

	// set logger
	if err := SetLogLevel(LogAccess, PushConf.Log.AccessLevel); err != nil {
		return errors.New("Set access log level error: "+ err.Error())
	}

	if err := SetLogLevel(LogError, PushConf.Log.ErrorLevel); err != nil {
		return errors.New("Set error log level error: "+ err.Error())
	}

	if err = SetLogOut(LogAccess, PushConf.Log.AccessLog); err != nil {
		return errors.New("Set access log path error: "+ err.Error())
	}

	if err = SetLogOut(LogError, PushConf.Log.ErrorLog); err != nil {
		return errors.New("Set error log path error: "+ err.Error())
	}

	return nil
}

func SetLogOut(log *logrus.Logger, outString string) error {
	switch outString {
	case "stdout":
		log.Out = os.Stdout
	case "stderr":
		log.Out = os.Stderr
	default:
		f, err := os.OpenFile(outString, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

		if err != nil {
			return err
		}

		log.Out = f
	}

	return nil
}

func SetLogLevel(log *logrus.Logger, levelString string) error {
	level, err := logrus.ParseLevel(levelString)

	if err != nil {
		return err
	}

	log.Level = level

	return nil
}
