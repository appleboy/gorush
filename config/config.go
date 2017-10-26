package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

var defaultConf = []byte(`
core:
  enabled: true # enabale httpd server
  address: "" # ip address to bind (default: any)
  port: "8088" # ignore this port number if auto_tls is enabled (listen 443).
  worker_num: 0 # default worker number is runtime.NumCPU()
  queue_num: 0 # default queue number is 8192
  max_notification: 100
  sync: false # set true if you need get error message from fail push notification in API response.
  mode: "release"
  ssl: false
  cert_path: "cert.pem"
  key_path: "key.pem"
  http_proxy: "" # only working for FCM server
  pid:
    enabled: false
    path: "gorush.pid"
    override: true
  auto_tls:
    enabled: false # Automatically install TLS certificates from Let's Encrypt.
    folder: ".cache" # folder for storing TLS certificates
    host: "" # which domains the Let's Encrypt will attempt

grpc:
  enabled: false # enabale gRPC server
  port: 9000

api:
  push_uri: "/api/push"
  stat_go_uri: "/api/stat/go"
  stat_app_uri: "/api/stat/app"
  config_uri: "/api/config"
  sys_stat_uri: "/sys/stats"
  metric_uri: "/metrics"
  health_uri: "/healthz"

android:
  enabled: true
  apikey: "YOUR_API_KEY"
  max_retry: 0 # resend fail notification, default value zero is disabled

ios:
  enabled: false
  key_path: "key.pem"
  password: "" # certificate password, default as empty string.
  production: false
  max_retry: 0 # resend fail notification, default value zero is disabled
  key_id: "" # KeyID from developer account (Certificates, Identifiers & Profiles -> Keys)
  team_id: "" # TeamID from developer account (View Account -> Membership)

log:
  format: "string" # string or json
  access_log: "stdout" # stdout: output to console, or define log path like "log/access_log"
  access_level: "debug"
  error_log: "stderr" # stderr: output to console, or define log path like "log/error_log"
  error_level: "error"
  hide_token: true

stat:
  engine: "memory" # support memory, redis, boltdb, buntdb or leveldb
  redis:
    addr: "localhost:6379"
    password: ""
    db: 0
  boltdb:
    path: "bolt.db"
    bucket: "gorush"
  buntdb:
    path: "bunt.db"
  leveldb:
    path: "level.db"
`)

// ConfYaml is config structure.
type ConfYaml struct {
	Core    SectionCore    `yaml:"core"`
	API     SectionAPI     `yaml:"api"`
	Android SectionAndroid `yaml:"android"`
	Ios     SectionIos     `yaml:"ios"`
	Log     SectionLog     `yaml:"log"`
	Stat    SectionStat    `yaml:"stat"`
	GRPC    SectionGRPC    `yaml:"grpc"`
}

// SectionCore is sub section of config.
type SectionCore struct {
	Enabled         bool           `yaml:"enabled"`
	Address         string         `yaml:"address"`
	Port            string         `yaml:"port"`
	MaxNotification int64          `yaml:"max_notification"`
	WorkerNum       int64          `yaml:"worker_num"`
	QueueNum        int64          `yaml:"queue_num"`
	Mode            string         `yaml:"mode"`
	Sync            bool           `yaml:"sync"`
	SSL             bool           `yaml:"ssl"`
	CertPath        string         `yaml:"cert_path"`
	KeyPath         string         `yaml:"key_path"`
	HTTPProxy       string         `yaml:"http_proxy"`
	PID             SectionPID     `yaml:"pid"`
	AutoTLS         SectionAutoTLS `yaml:"auto_tls"`
}

// SectionAutoTLS support Let's Encrypt setting.
type SectionAutoTLS struct {
	Enabled bool   `yaml:"enabled"`
	Folder  string `yaml:"folder"`
	Host    string `yaml:"host"`
}

// SectionAPI is sub section of config.
type SectionAPI struct {
	PushURI    string `yaml:"push_uri"`
	StatGoURI  string `yaml:"stat_go_uri"`
	StatAppURI string `yaml:"stat_app_uri"`
	ConfigURI  string `yaml:"config_uri"`
	SysStatURI string `yaml:"sys_stat_uri"`
	MetricURI  string `yaml:"metric_uri"`
	HealthURI  string `yaml:"health_uri"`
}

// SectionAndroid is sub section of config.
type SectionAndroid struct {
	Enabled  bool   `yaml:"enabled"`
	APIKey   string `yaml:"apikey"`
	MaxRetry int    `yaml:"max_retry"`
}

// SectionIos is sub section of config.
type SectionIos struct {
	Enabled    bool   `yaml:"enabled"`
	KeyPath    string `yaml:"key_path"`
	Password   string `yaml:"password"`
	Production bool   `yaml:"production"`
	MaxRetry   int    `yaml:"max_retry"`
	KeyID      string `yaml:"key_id"`
	TeamID     string `yaml:"team_id"`
}

// SectionLog is sub section of config.
type SectionLog struct {
	Format      string `yaml:"format"`
	AccessLog   string `yaml:"access_log"`
	AccessLevel string `yaml:"access_level"`
	ErrorLog    string `yaml:"error_log"`
	ErrorLevel  string `yaml:"error_level"`
	HideToken   bool   `yaml:"hide_token"`
}

// SectionStat is sub section of config.
type SectionStat struct {
	Engine  string         `yaml:"engine"`
	Redis   SectionRedis   `yaml:"redis"`
	BoltDB  SectionBoltDB  `yaml:"boltdb"`
	BuntDB  SectionBuntDB  `yaml:"buntdb"`
	LevelDB SectionLevelDB `yaml:"leveldb"`
}

// SectionRedis is sub section of config.
type SectionRedis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// SectionBoltDB is sub section of config.
type SectionBoltDB struct {
	Path   string `yaml:"path"`
	Bucket string `yaml:"bucket"`
}

// SectionBuntDB is sub section of config.
type SectionBuntDB struct {
	Path string `yaml:"path"`
}

// SectionLevelDB is sub section of config.
type SectionLevelDB struct {
	Path string `yaml:"path"`
}

// SectionPID is sub section of config.
type SectionPID struct {
	Enabled  bool   `yaml:"enabled"`
	Path     string `yaml:"path"`
	Override bool   `yaml:"override"`
}

// SectionGRPC is sub section of config.
type SectionGRPC struct {
	Enabled bool   `yaml:"enabled"`
	Port    string `yaml:"port"`
}

// LoadConf load config from file and read in environment variables that match
func LoadConf(confPath string) (ConfYaml, error) {
	var conf ConfYaml

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()         // read in environment variables that match
	viper.SetEnvPrefix("gorush") // will be uppercased automatically
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if confPath != "" {
		content, err := ioutil.ReadFile(confPath)

		if err != nil {
			return conf, err
		}

		viper.ReadConfig(bytes.NewBuffer(content))
	} else {
		// Search config in home directory with name ".gorush" (without extension).
		viper.AddConfigPath("/etc/gorush/")
		viper.AddConfigPath("$HOME/.gorush")
		viper.AddConfigPath(".")
		viper.SetConfigName("config")

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		} else {
			// load default config
			viper.ReadConfig(bytes.NewBuffer(defaultConf))
		}
	}

	// Core
	conf.Core.Address = viper.GetString("core.address")
	conf.Core.Port = viper.GetString("core.port")
	conf.Core.Enabled = viper.GetBool("core.enabled")
	conf.Core.WorkerNum = int64(viper.GetInt("core.worker_num"))
	conf.Core.QueueNum = int64(viper.GetInt("core.queue_num"))
	conf.Core.Mode = viper.GetString("core.mode")
	conf.Core.Sync = viper.GetBool("core.sync")
	conf.Core.SSL = viper.GetBool("core.ssl")
	conf.Core.CertPath = viper.GetString("core.cert_path")
	conf.Core.KeyPath = viper.GetString("core.key_path")
	conf.Core.MaxNotification = int64(viper.GetInt("core.max_notification"))
	conf.Core.HTTPProxy = viper.GetString("core.http_proxy")
	conf.Core.PID.Enabled = viper.GetBool("core.pid.enabled")
	conf.Core.PID.Path = viper.GetString("core.pid.path")
	conf.Core.PID.Override = viper.GetBool("core.pid.override")
	conf.Core.AutoTLS.Enabled = viper.GetBool("core.auto_tls.enabled")
	conf.Core.AutoTLS.Folder = viper.GetString("core.auto_tls.folder")
	conf.Core.AutoTLS.Host = viper.GetString("core.auto_tls.host")

	// Api
	conf.API.PushURI = viper.GetString("api.push_uri")
	conf.API.StatGoURI = viper.GetString("api.stat_go_uri")
	conf.API.StatAppURI = viper.GetString("api.stat_app_uri")
	conf.API.ConfigURI = viper.GetString("api.config_uri")
	conf.API.SysStatURI = viper.GetString("api.sys_stat_uri")
	conf.API.MetricURI = viper.GetString("api.metric_uri")
	conf.API.HealthURI = viper.GetString("api.health_uri")

	// Android
	conf.Android.Enabled = viper.GetBool("android.enabled")
	conf.Android.APIKey = viper.GetString("android.apikey")
	conf.Android.MaxRetry = viper.GetInt("android.max_retry")

	// iOS
	conf.Ios.Enabled = viper.GetBool("ios.enabled")
	conf.Ios.KeyPath = viper.GetString("ios.key_path")
	conf.Ios.Password = viper.GetString("ios.password")
	conf.Ios.Production = viper.GetBool("ios.production")
	conf.Ios.MaxRetry = viper.GetInt("ios.max_retry")
	conf.Ios.KeyID = viper.GetString("ios.key_id")
	conf.Ios.TeamID = viper.GetString("ios.team_id")

	// log
	conf.Log.Format = viper.GetString("log.format")
	conf.Log.AccessLog = viper.GetString("log.access_log")
	conf.Log.AccessLevel = viper.GetString("log.access_level")
	conf.Log.ErrorLog = viper.GetString("log.error_log")
	conf.Log.ErrorLevel = viper.GetString("log.error_level")
	conf.Log.HideToken = viper.GetBool("log.hide_token")

	// Stat Engine
	conf.Stat.Engine = viper.GetString("stat.engine")
	conf.Stat.Redis.Addr = viper.GetString("stat.redis.addr")
	conf.Stat.Redis.Password = viper.GetString("stat.redis.password")
	conf.Stat.Redis.DB = viper.GetInt("stat.redis.db")
	conf.Stat.BoltDB.Path = viper.GetString("stat.boltdb.path")
	conf.Stat.BoltDB.Bucket = viper.GetString("stat.boltdb.bucket")
	conf.Stat.BuntDB.Path = viper.GetString("stat.buntdb.path")
	conf.Stat.LevelDB.Path = viper.GetString("stat.leveldb.path")

	// gRPC Server
	conf.GRPC.Enabled = viper.GetBool("grpc.enabled")
	conf.GRPC.Port = viper.GetString("grpc.port")

	if conf.Core.WorkerNum == int64(0) {
		conf.Core.WorkerNum = int64(runtime.NumCPU())
	}

	if conf.Core.QueueNum == int64(0) {
		conf.Core.QueueNum = int64(8192)
	}

	return conf, nil
}
