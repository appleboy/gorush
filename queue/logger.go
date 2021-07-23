package queue

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

// Logger interface is used throughout gorush
type Logger interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

type defaultLogger struct{}

func (l defaultLogger) Infof(format string, args ...interface{}) {
	log.Info().Msgf(format, args...)
}

func (l defaultLogger) Errorf(format string, args ...interface{}) {
	log.Error().Msgf(format, args...)
}

func (l defaultLogger) Fatalf(format string, args ...interface{}) {
	log.Fatal().Msgf(format, args...)
}

func (l defaultLogger) Info(args ...interface{}) {
	log.Info().Msg(fmt.Sprint(args...))
}

func (l defaultLogger) Error(args ...interface{}) {
	log.Error().Msg(fmt.Sprint(args...))
}

func (l defaultLogger) Fatal(args ...interface{}) {
	log.Fatal().Msg(fmt.Sprint(args...))
}
