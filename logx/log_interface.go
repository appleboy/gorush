package logx

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// QueueLogger for simple logger.
func QueueLogger() DefaultQueueLogger {
	return DefaultQueueLogger{
		accessLogger: LogAccess,
		errorLogger:  LogError,
	}
}

// DefaultQueueLogger for queue custom logger
type DefaultQueueLogger struct {
	accessLogger *logrus.Logger
	errorLogger  *logrus.Logger
}

func (l DefaultQueueLogger) Infof(format string, args ...interface{}) {
	l.accessLogger.Printf(format, args...)
}

func (l DefaultQueueLogger) Errorf(format string, args ...interface{}) {
	l.errorLogger.Printf(format, args...)
}

func (l DefaultQueueLogger) Fatalf(format string, args ...interface{}) {
	l.errorLogger.Fatalf(format, args...)
}

func (l DefaultQueueLogger) Info(args ...interface{}) {
	l.accessLogger.Println(fmt.Sprint(args...))
}

func (l DefaultQueueLogger) Error(args ...interface{}) {
	l.errorLogger.Println(fmt.Sprint(args...))
}

func (l DefaultQueueLogger) Fatal(args ...interface{}) {
	l.errorLogger.Println(fmt.Sprint(args...))
}
