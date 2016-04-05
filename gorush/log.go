package gopush

import (
	"github.com/Sirupsen/logrus"
	"log"
	"os"
)

func InitLog() {

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
	err = SetLogLevel(LogAccess, PushConf.Log.AccessLevel)
	if err != nil {
		log.Fatal(err)
	}
	err = SetLogLevel(LogError, PushConf.Log.ErrorLevel)
	if err != nil {
		log.Fatal(err)
	}
	err = SetLogOut(LogAccess, PushConf.Log.AccessLog)
	if err != nil {
		log.Fatal(err)
	}
	err = SetLogOut(LogError, PushConf.Log.ErrorLog)
	if err != nil {
		log.Fatal(err)
	}
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
