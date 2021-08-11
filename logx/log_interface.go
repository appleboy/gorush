package logx

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// QueueLogger for simple logger.
func QueueLogger() defaultLogger {
	return defaultLogger{
		accessLogger: LogAccess,
		errorLogger:  LogError,
	}
}

type defaultLogger struct {
	accessLogger *logrus.Logger
	errorLogger  *logrus.Logger
}

func (l defaultLogger) Infof(format string, args ...interface{}) {
	l.accessLogger.Printf(format, args...)
}

func (l defaultLogger) Errorf(format string, args ...interface{}) {
	l.errorLogger.Printf(format, args...)
}

func (l defaultLogger) Fatalf(format string, args ...interface{}) {
	l.errorLogger.Fatalf(format, args...)
}

func (l defaultLogger) Info(args ...interface{}) {
	l.accessLogger.Println(fmt.Sprint(args...))
}

func (l defaultLogger) Error(args ...interface{}) {
	l.errorLogger.Println(fmt.Sprint(args...))
}

func (l defaultLogger) Fatal(args ...interface{}) {
	l.errorLogger.Println(fmt.Sprint(args...))
}
