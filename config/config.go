package config

import (
	"io/ioutil"
	"runtime"

	"github.com/pazams/envexpand"
	"gopkg.in/yaml.v2"
)

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

// BuildDefaultPushConf is default config setting.
func BuildDefaultPushConf() ConfYaml {
	var conf ConfYaml

	// Core
	conf.Core.Address = ""
	conf.Core.Port = "8088"
	conf.Core.Enabled = true
	conf.Core.WorkerNum = int64(runtime.NumCPU())
	conf.Core.QueueNum = int64(8192)
	conf.Core.Mode = "release"
	conf.Core.Sync = false
	conf.Core.SSL = false
	conf.Core.CertPath = "cert.pem"
	conf.Core.KeyPath = "key.pem"
	conf.Core.MaxNotification = int64(100)
	conf.Core.HTTPProxy = ""
	conf.Core.PID.Enabled = false
	conf.Core.PID.Path = "gorush.pid"
	conf.Core.PID.Override = false
	conf.Core.AutoTLS.Enabled = false
	conf.Core.AutoTLS.Folder = ".cache"
	conf.Core.AutoTLS.Host = ""

	// Api
	conf.API.PushURI = "/api/push"
	conf.API.StatGoURI = "/api/stat/go"
	conf.API.StatAppURI = "/api/stat/app"
	conf.API.ConfigURI = "/api/config"
	conf.API.SysStatURI = "/sys/stats"
	conf.API.MetricURI = "/metrics"

	// Android
	conf.Android.Enabled = false
	conf.Android.APIKey = ""
	conf.Android.MaxRetry = 0

	// iOS
	conf.Ios.Enabled = false
	conf.Ios.KeyPath = "key.pem"
	conf.Ios.Password = ""
	conf.Ios.Production = false
	conf.Ios.MaxRetry = 0

	// log
	conf.Log.Format = "string"
	conf.Log.AccessLog = "stdout"
	conf.Log.AccessLevel = "debug"
	conf.Log.ErrorLog = "stderr"
	conf.Log.ErrorLevel = "error"
	conf.Log.HideToken = true

	// Stat Engine
	conf.Stat.Engine = "memory"
	conf.Stat.Redis.Addr = "localhost:6379"
	conf.Stat.Redis.Password = ""
	conf.Stat.Redis.DB = 0
	conf.Stat.BoltDB.Path = "bolt.db"
	conf.Stat.BoltDB.Bucket = "gorush"
	conf.Stat.BuntDB.Path = "bunt.db"
	conf.Stat.LevelDB.Path = "level.db"

	// gRPC Server
	conf.GRPC.Enabled = false
	conf.GRPC.Port = "50051"
	return conf
}

// LoadConfYaml provide load yml config.
func LoadConfYaml(confPath string) (ConfYaml, error) {
	var config ConfYaml

	configFile, err := ioutil.ReadFile(confPath)
	configFileExpanded := envexpand.Expand(configFile)

	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(configFileExpanded, &config)

	if err != nil {
		return config, err
	}

	if config.Core.WorkerNum == int64(0) {
		config.Core.WorkerNum = int64(runtime.NumCPU())
	}

	if config.Core.QueueNum == int64(0) {
		config.Core.QueueNum = int64(8192)
	}

	return config, nil
}
