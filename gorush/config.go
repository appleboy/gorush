package gorush

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"runtime"
)

// ConfYaml is config structure.
type ConfYaml struct {
	Core    SectionCore    `yaml:"core"`
	API     SectionAPI     `yaml:"api"`
	Android SectionAndroid `yaml:"android"`
	Ios     SectionIos     `yaml:"ios"`
	Log     SectionLog     `yaml:"log"`
}

// SectionCore is sub seciont of config.
type SectionCore struct {
	Port            string `yaml:"port"`
	MaxNotification int    `yaml:"max_notification"`
	WorkerNum       int    `yaml:"worker_num"`
	QueueNum        int    `yaml:"queue_num"`
	Mode            string `yaml:"mode"`
	SSL             bool   `yaml:"ssl"`
	CertPath        string `yaml:"cert_path"`
	KeyPath         string `yaml:"key_path"`
}

// SectionAPI is sub seciont of config.
type SectionAPI struct {
	PushURI   string `yaml:"push_uri"`
	StatGoURI string `yaml:"stat_go_uri"`
}

// SectionAndroid is sub seciont of config.
type SectionAndroid struct {
	Enabled bool   `yaml:"enabled"`
	APIKey  string `yaml:"apikey"`
}

// SectionIos is sub seciont of config.
type SectionIos struct {
	Enabled     bool   `yaml:"enabled"`
	PemCertPath string `yaml:"pem_cert_path"`
	PemKeyPath  string `yaml:"pem_key_path"`
	Production  bool   `yaml:"production"`
}

// SectionLog is sub seciont of config.
type SectionLog struct {
	Format      string `yaml:"format"`
	AccessLog   string `yaml:"access_log"`
	AccessLevel string `yaml:"access_level"`
	ErrorLog    string `yaml:"error_log"`
	ErrorLevel  string `yaml:"error_level"`
}

// BuildDefaultPushConf is default config setting.
func BuildDefaultPushConf() ConfYaml {
	var conf ConfYaml

	// Core
	conf.Core.Port = "8088"
	conf.Core.WorkerNum = runtime.NumCPU()
	conf.Core.QueueNum = 8192
	conf.Core.Mode = "release"
	conf.Core.SSL = false
	conf.Core.CertPath = "cert.pem"
	conf.Core.KeyPath = "key.pem"
	conf.Core.MaxNotification = 100

	// Api
	conf.API.PushURI = "/api/push"
	conf.API.StatGoURI = "/api/status"

	// Android
	conf.Android.Enabled = false
	conf.Android.APIKey = ""

	// iOS
	conf.Ios.Enabled = false
	conf.Ios.PemCertPath = "cert.pem"
	conf.Ios.PemKeyPath = "key.pem"
	conf.Ios.Production = false

	// log
	conf.Log.Format = "string"
	conf.Log.AccessLog = "stdout"
	conf.Log.AccessLevel = "debug"
	conf.Log.ErrorLog = "stderr"
	conf.Log.ErrorLevel = "error"

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
