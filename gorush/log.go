package gopush

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"os"
	// "time"
)

var (
	green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow  = string([]byte{27, 91, 57, 55, 59, 52, 51, 109})
	red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset   = string([]byte{27, 91, 48, 109})
)

type LogReq struct {
	URI         string `json:"uri"`
	Method      string `json:"method"`
	IP          string `json:"ip"`
	ContentType string `json:"content_type"`
	Agent       string `json:"agent"`
}

type LogPushEntry struct {
	Type     string `json:"type"`
	Platform string `json:"platform"`
	Token    string `json:"token"`
	Message  string `json:"message"`
	Error    string `json:"error"`

	// Android
	To                    string `json:"to,omitempty"`
	CollapseKey           string `json:"collapse_key,omitempty"`
	DelayWhileIdle        bool   `json:"delay_while_idle,omitempty"`
	TimeToLive            uint   `json:"time_to_live,omitempty"`
	RestrictedPackageName string `json:"restricted_package_name,omitempty"`
	DryRun                bool   `json:"dry_run,omitempty"`

	// iOS
	ApnsID   string `json:"apns_id,omitempty"`
	Topic    string `json:"topic,omitempty"`
	Badge    int    `json:"badge,omitempty"`
	Sound    string `json:"sound,omitempty"`
	Category string `json:"category,omitempty"`
}

func InitLog() error {

	var err error

	// init logger
	LogAccess = logrus.New()
	LogError = logrus.New()

	LogAccess.Formatter = &logrus.TextFormatter{
		TimestampFormat: "2006/01/02 - 15:04:05",
		ForceColors:     true,
		FullTimestamp:   true,
	}

	LogError.Formatter = &logrus.TextFormatter{
		TimestampFormat: "2006/01/02 - 15:04:05",
		ForceColors:     true,
		FullTimestamp:   true,
	}

	// set logger
	if err := SetLogLevel(LogAccess, PushConf.Log.AccessLevel); err != nil {
		return errors.New("Set access log level error: " + err.Error())
	}

	if err := SetLogLevel(LogError, PushConf.Log.ErrorLevel); err != nil {
		return errors.New("Set error log level error: " + err.Error())
	}

	if err = SetLogOut(LogAccess, PushConf.Log.AccessLog); err != nil {
		return errors.New("Set access log path error: " + err.Error())
	}

	if err = SetLogOut(LogError, PushConf.Log.ErrorLog); err != nil {
		return errors.New("Set error log path error: " + err.Error())
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

func LogRequest(uri string, method string, ip string, contentType string, agent string) {
	var output string
	log := &LogReq{
		URI:         uri,
		Method:      method,
		IP:          ip,
		ContentType: contentType,
		Agent:       agent,
	}

	if PushConf.Log.Format == "json" {
		logJson, _ := json.Marshal(log)

		output = string(logJson)
	} else {
		// format is string
		output = fmt.Sprintf("|%s header %s| %s %s %s %s %s", magenta, reset, log.Method, log.URI, log.IP, log.ContentType, log.Agent)
	}

	LogAccess.Info(output)
}

func colorForPlatForm(platform int) string {
	switch platform {
	case PlatFormIos:
		return blue
	case PlatFormAndroid:
		return cyan
	default:
		return reset
	}
}

func typeForPlatForm(platform int) string {
	switch platform {
	case PlatFormIos:
		return "ios"
	case PlatFormAndroid:
		return "android"
	default:
		return ""
	}
}

func LogPush(status, token string, req RequestPushNotification, errPush error) {
	var plat, platColor, output string

	platColor = colorForPlatForm(req.Platform)
	plat = typeForPlatForm(req.Platform)

	errMsg := ""
	if errPush != nil {
		errMsg = errPush.Error()
	}

	log := &LogPushEntry{
		Type:     status,
		Platform: plat,
		Token:    token,
		Message:  req.Message,
		Error:    errMsg,
	}

	if PushConf.Log.Format == "json" {
		logJson, _ := json.Marshal(log)

		output = string(logJson)
	} else {
		switch status {
		case StatusSucceededPush:
			output = fmt.Sprintf("|%s %s %s| %s%s%s [%s] %s", green, log.Type, reset, platColor, log.Platform, reset, log.Token, log.Message)
		case StatusFailedPush:
			output = fmt.Sprintf("|%s %s %s| %s%s%s [%s] | %s | Error Message: %s", red, log.Type, reset, platColor, log.Platform, reset, log.Token, log.Message, log.Error)
		}
	}

	switch status {
	case StatusSucceededPush:
		LogAccess.Info(string(output))
	case StatusFailedPush:
		LogError.Error(string(output))
	}
}

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		LogRequest(c.Request.URL.Path, c.Request.Method, c.ClientIP(), c.ContentType(), c.Request.Header.Get("User-Agent"))
		c.Next()
	}
}
