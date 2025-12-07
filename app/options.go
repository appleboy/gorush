package app

import (
	"flag"

	"github.com/appleboy/gorush/config"
)

// Options holds all CLI flag values.
type Options struct {
	ShowVersion bool
	Ping        bool
	ConfigFile  string

	// Notification options (for CLI mode)
	Token   string
	Message string
	Title   string
	Topic   string

	// Config overrides
	Conf config.ConfYaml
}

// NewOptions creates a new Options instance with default values.
func NewOptions() *Options {
	return &Options{}
}

// BindFlags binds CLI flags to the Options struct.
// Call this before flag.Parse().
func (o *Options) BindFlags() {
	// Version flags
	flag.BoolVar(&o.ShowVersion, "version", false, "Print version information.")
	flag.BoolVar(&o.ShowVersion, "V", false, "Print version information.")

	// Config file
	flag.StringVar(&o.ConfigFile, "c", "", "Configuration file path.")
	flag.StringVar(&o.ConfigFile, "config", "", "Configuration file path.")

	// PID file
	flag.StringVar(&o.Conf.Core.PID.Path, "pid", "", "PID file path.")

	// iOS options
	flag.StringVar(&o.Conf.Ios.KeyPath, "i", "", "iOS certificate key file path")
	flag.StringVar(&o.Conf.Ios.KeyPath, "key", "", "iOS certificate key file path")
	flag.StringVar(&o.Conf.Ios.KeyID, "key-id", "", "iOS Key ID for P8 token")
	flag.StringVar(&o.Conf.Ios.TeamID, "team-id", "", "iOS Team ID for P8 token")
	flag.StringVar(&o.Conf.Ios.Password, "P", "", "iOS certificate password for gorush")
	flag.StringVar(&o.Conf.Ios.Password, "password", "", "iOS certificate password for gorush")
	flag.BoolVar(&o.Conf.Ios.Enabled, "ios", false, "send ios notification")
	flag.BoolVar(&o.Conf.Ios.Production, "production", false, "production mode in iOS")

	// Android options
	flag.StringVar(&o.Conf.Android.KeyPath, "fcm-key", "", "FCM key path configuration for gorush")
	flag.BoolVar(&o.Conf.Android.Enabled, "android", false, "send android notification")

	// Huawei options
	flag.StringVar(&o.Conf.Huawei.AppSecret, "hk", "", "Huawei api key configuration for gorush")
	flag.StringVar(
		&o.Conf.Huawei.AppSecret,
		"hmskey",
		"",
		"Huawei api key configuration for gorush",
	)
	flag.StringVar(&o.Conf.Huawei.AppID, "hid", "", "HMS app id configuration for gorush")
	flag.StringVar(&o.Conf.Huawei.AppID, "hmsid", "", "HMS app id configuration for gorush")
	flag.BoolVar(&o.Conf.Huawei.Enabled, "huawei", false, "send huawei notification")

	// Server options
	flag.StringVar(&o.Conf.Core.Address, "A", "", "address to bind")
	flag.StringVar(&o.Conf.Core.Address, "address", "", "address to bind")
	flag.StringVar(&o.Conf.Core.Port, "p", "", "port number for gorush")
	flag.StringVar(&o.Conf.Core.Port, "port", "", "port number for gorush")
	flag.StringVar(&o.Conf.Core.HTTPProxy, "proxy", "", "http proxy url")

	// Storage options
	flag.StringVar(&o.Conf.Stat.Engine, "e", "", "store engine")
	flag.StringVar(&o.Conf.Stat.Engine, "engine", "", "store engine")
	flag.StringVar(&o.Conf.Stat.Redis.Addr, "redis-addr", "", "redis addr")

	// Notification options (CLI mode)
	flag.StringVar(&o.Token, "t", "", "token string")
	flag.StringVar(&o.Token, "token", "", "token string")
	flag.StringVar(&o.Message, "m", "", "notification message")
	flag.StringVar(&o.Message, "message", "", "notification message")
	flag.StringVar(&o.Title, "title", "", "notification title")
	flag.StringVar(&o.Topic, "topic", "", "apns topic in iOS")

	// Health check
	flag.BoolVar(&o.Ping, "ping", false, "ping server")
}

// CLISendOptions returns CLI send options for notification sending.
func (o *Options) CLISendOptions() CLISendOptions {
	return CLISendOptions{
		Token:   o.Token,
		Message: o.Message,
		Title:   o.Title,
		Topic:   o.Topic,
	}
}

// IsCLIMode returns true if running in CLI notification mode.
func (o *Options) IsCLIMode() bool {
	return o.Conf.Android.Enabled || o.Conf.Ios.Enabled || o.Conf.Huawei.Enabled
}
