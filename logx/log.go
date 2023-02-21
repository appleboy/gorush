package logx

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/appleboy/gorush/core"

	"github.com/mattn/go-isatty"
	"github.com/sirupsen/logrus"
)

var (
	green  = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	yellow = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	red    = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue   = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	reset  = string([]byte{27, 91, 48, 109})
)

// LogPushEntry is push response log
type LogPushEntry struct {
	ID       string `json:"notif_id,omitempty"`
	Type     string `json:"type"`
	Platform string `json:"platform"`
	Token    string `json:"token"`
	Message  string `json:"message"`
	Error    string `json:"error"`
}

var isTerm bool

//nolint
func init() {
	isTerm = isatty.IsTerminal(os.Stdout.Fd())
}

var (
	// LogAccess is log server request log
	LogAccess = logrus.New()
	// LogError is log server error log
	LogError = logrus.New()
)

// InitLog use for initial log module
func InitLog(accessLevel, accessLog, errorLevel, errorLog string) error {
	var err error

	if !isTerm {
		LogAccess.SetFormatter(&logrus.JSONFormatter{})
		LogError.SetFormatter(&logrus.JSONFormatter{})
	} else {
		LogAccess.Formatter = &logrus.TextFormatter{
			TimestampFormat: "2006/01/02 - 15:04:05",
			FullTimestamp:   true,
		}

		LogError.Formatter = &logrus.TextFormatter{
			TimestampFormat: "2006/01/02 - 15:04:05",
			FullTimestamp:   true,
		}
	}

	// set logger
	if err = SetLogLevel(LogAccess, accessLevel); err != nil {
		return errors.New("Set access log level error: " + err.Error())
	}

	if err = SetLogLevel(LogError, errorLevel); err != nil {
		return errors.New("Set error log level error: " + err.Error())
	}

	if err = SetLogOut(LogAccess, accessLog); err != nil {
		return errors.New("Set access log path error: " + err.Error())
	}

	if err = SetLogOut(LogError, errorLog); err != nil {
		return errors.New("Set error log path error: " + err.Error())
	}

	return nil
}

// SetLogOut provide log stdout and stderr output
func SetLogOut(log *logrus.Logger, outString string) error {
	switch outString {
	case "stdout":
		log.Out = os.Stdout
	case "stderr":
		log.Out = os.Stderr
	default:
		f, err := os.OpenFile(outString, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o600)
		if err != nil {
			return err
		}

		log.Out = f
	}

	return nil
}

// SetLogLevel is define log level what you want
// log level: panic, fatal, error, warn, info and debug
func SetLogLevel(log *logrus.Logger, levelString string) error {
	level, err := logrus.ParseLevel(levelString)
	if err != nil {
		return err
	}

	log.Level = level

	return nil
}

func colorForPlatForm(platform int) string {
	switch platform {
	case core.PlatFormIos:
		return blue
	case core.PlatFormAndroid:
		return yellow
	case core.PlatFormHuawei:
		return green
	default:
		return reset
	}
}

func typeForPlatForm(platform int) string {
	switch platform {
	case core.PlatFormIos:
		return "ios"
	case core.PlatFormAndroid:
		return "android"
	case core.PlatFormHuawei:
		return "huawei"
	default:
		return ""
	}
}

func hideToken(token string, markLen int) string {
	if token == "" {
		return ""
	}

	if len(token) < markLen*2 {
		return strings.Repeat("*", len(token))
	}

	start := token[len(token)-markLen:]
	end := token[0:markLen]

	result := strings.ReplaceAll(token, start, strings.Repeat("*", markLen))
	result = strings.ReplaceAll(result, end, strings.Repeat("*", markLen))

	return result
}

// GetLogPushEntry get push data into log structure
func GetLogPushEntry(input *InputLog) LogPushEntry {
	var errMsg string

	plat := typeForPlatForm(input.Platform)

	if input.Error != nil {
		errMsg = input.Error.Error()
	}

	token := input.Token
	if input.HideToken {
		token = hideToken(input.Token, 10)
	}

	message := input.Message
	if input.HideMessage {
		message = "(message redacted)"
	}

	return LogPushEntry{
		ID:       input.ID,
		Type:     input.Status,
		Platform: plat,
		Token:    token,
		Message:  message,
		Error:    errMsg,
	}
}

// InputLog log request
type InputLog struct {
	ID          string
	Status      string
	Token       string
	Message     string
	Platform    int
	Error       error
	HideToken   bool
	HideMessage bool
	Format      string
}

// LogPush record user push request and server response.
func LogPush(input *InputLog) LogPushEntry {
	var platColor, resetColor, output string

	if isTerm {
		platColor = colorForPlatForm(input.Platform)
		resetColor = reset
	}

	log := GetLogPushEntry(input)

	if input.Format == "json" {
		logJSON, _ := json.Marshal(log)

		output = string(logJSON)
	} else {
		var typeColor string
		switch input.Status {
		case core.SucceededPush:
			if isTerm {
				typeColor = green
			}

			output = fmt.Sprintf("|%s %s %s| %s%s%s [%s] %s",
				typeColor, log.Type, resetColor,
				platColor, log.Platform, resetColor,
				log.Token,
				log.Message,
			)
		case core.FailedPush:
			if isTerm {
				typeColor = red
			}

			output = fmt.Sprintf("|%s %s %s| %s%s%s [%s] | %s | Error Message: %s",
				typeColor, log.Type, resetColor,
				platColor, log.Platform, resetColor,
				log.Token,
				log.Message,
				log.Error,
			)
		}
	}

	switch input.Status {
	case core.SucceededPush:
		LogAccess.Info(output)
	case core.FailedPush:
		LogError.Error(output)
	}

	return log
}
