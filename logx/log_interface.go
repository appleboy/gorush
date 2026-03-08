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

func (l DefaultQueueLogger) Infof(format string, args ...any) {
	l.accessLogger.Printf(format, args...)
}

func (l DefaultQueueLogger) Errorf(format string, args ...any) {
	l.errorLogger.Printf(format, args...)
}

func (l DefaultQueueLogger) Fatalf(format string, args ...any) {
	l.errorLogger.Fatalf(format, args...)
}

func (l DefaultQueueLogger) Info(args ...any) {
	l.accessLogger.Println(fmt.Sprint(args...))
}

func (l DefaultQueueLogger) Error(args ...any) {
	l.errorLogger.Println(fmt.Sprint(args...))
}

func (l DefaultQueueLogger) Fatal(args ...any) {
	l.errorLogger.Println(fmt.Sprint(args...))
}
