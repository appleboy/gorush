package config

import (
	"io/ioutil"
	"runtime"

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
}

// SectionCore is sub section of config.
type SectionCore struct {
	Port            string     `yaml:"port"`
	MaxNotification int64      `yaml:"max_notification"`
	WorkerNum       int64      `yaml:"worker_num"`
	QueueNum        int64      `yaml:"queue_num"`
	Mode            string     `yaml:"mode"`
	SSL             bool       `yaml:"ssl"`
	CertPath        string     `yaml:"cert_path"`
	KeyPath         string     `yaml:"key_path"`
	HTTPProxy       string     `yaml:"http_proxy"`
	PID             SectionPID `yaml:"pid"`
}

// SectionAPI is sub section of config.
type SectionAPI struct {
	PushURI    string `yaml:"push_uri"`
	StatGoURI  string `yaml:"stat_go_uri"`
	StatAppURI string `yaml:"stat_app_uri"`
	ConfigURI  string `yaml:"config_uri"`
	SysStatURI string `yaml:"sys_stat_uri"`
}

// SectionAndroid is sub section of config.
type SectionAndroid struct {
	Enabled bool   `yaml:"enabled"`
	APIKey  string `yaml:"apikey"`
}

// SectionIos is sub section of config.
type SectionIos struct {
	Enabled    bool   `yaml:"enabled"`
	KeyPath    string `yaml:"key_path"`
	Password   string `yaml:"password"`
	Production bool   `yaml:"production"`
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

// BuildDefaultPushConf is default config setting.
func BuildDefaultPushConf() ConfYaml {
	var conf ConfYaml

	// Core
	conf.Core.Port = "8088"
	conf.Core.WorkerNum = int64(runtime.NumCPU())
	conf.Core.QueueNum = int64(8192)
	conf.Core.Mode = "release"
	conf.Core.SSL = false
	conf.Core.CertPath = "cert.pem"
	conf.Core.KeyPath = "key.pem"
	conf.Core.MaxNotification = int64(100)
	conf.Core.HTTPProxy = ""
	conf.Core.PID.Enabled = false
	conf.Core.PID.Path = "gorush.pid"
	conf.Core.PID.Override = false

	// Api
	conf.API.PushURI = "/api/push"
	conf.API.StatGoURI = "/api/stat/go"
	conf.API.StatAppURI = "/api/stat/app"
	conf.API.ConfigURI = "/api/config"
	conf.API.SysStatURI = "/sys/stats"

	// Android
	conf.Android.Enabled = false
	conf.Android.APIKey = ""

	// iOS
	conf.Ios.Enabled = false
	conf.Ios.KeyPath = "key.pem"
	conf.Ios.Password = ""
	conf.Ios.Production = false

	// log
	conf.Log.Format = "string"
	conf.Log.AccessLog = "stdout"
	conf.Log.AccessLevel = "debug"
	conf.Log.ErrorLog = "stderr"
	conf.Log.ErrorLevel = "error"
	conf.Log.HideToken = true

	conf.Stat.Engine = "memory"
	conf.Stat.Redis.Addr = "localhost:6379"
	conf.Stat.Redis.Password = ""
	conf.Stat.Redis.DB = 0

	conf.Stat.BoltDB.Path = "bolt.db"
	conf.Stat.BoltDB.Bucket = "gorush"

	conf.Stat.BuntDB.Path = "bunt.db"
	conf.Stat.LevelDB.Path = "level.db"

	return conf
}

// LoadConfYaml provide load yml config.
func LoadConfYaml(confPath string) (ConfYaml, error) {
	var config ConfYaml

	configFile, err := ioutil.ReadFile(confPath)

	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal([]byte(configFile), &config)

	if err != nil {
		return config, err
	}

	return config, nil
}
