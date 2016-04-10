package gopush

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type ConfYaml struct {
	Core    SectionCore    `yaml:"core"`
	Api     SectionApi     `yaml:"api"`
	Android SectionAndroid `yaml:"android"`
	Ios     SectionIos     `yaml:"ios"`
	Log     SectionLog     `yaml:"log"`
}

type SectionCore struct {
	Port            string `yaml:"port"`
	MaxNotification int    `yaml:"max_notification"`
	Mode            string `yaml:"mode"`
	SSL             bool   `yaml:"ssl"`
	CertPath        string `yaml:"cert_path"`
	KeyPath         string `yaml:"key_path"`
}

type SectionApi struct {
	PushUri   string `yaml:"push_uri"`
	StatGoUri string `yaml:"stat_go_uri"`
}

type SectionAndroid struct {
	Enabled bool   `yaml:"enabled"`
	ApiKey  string `yaml:"apikey"`
}

type SectionIos struct {
	Enabled     bool   `yaml:"enabled"`
	PemCertPath string `yaml:"pem_cert_path"`
	PemKeyPath  string `yaml:"pem_key_path"`
	Production  bool   `yaml:"production"`
}

type SectionLog struct {
	Format      string `yaml:"format"`
	AccessLog   string `yaml:"access_log"`
	AccessLevel string `yaml:"access_level"`
	ErrorLog    string `yaml:"error_log"`
	ErrorLevel  string `yaml:"error_level"`
}

func BuildDefaultPushConf() ConfYaml {
	var conf ConfYaml

	// Core
	conf.Core.Port = "8088"
	conf.Core.Mode = "release"
	conf.Core.SSL = false
	conf.Core.CertPath = "cert.pem"
	conf.Core.KeyPath = "key.pem"
	conf.Core.MaxNotification = 100

	// Api
	conf.Api.PushUri = "/api/push"
	conf.Api.StatGoUri = "/api/status"

	// Android
	conf.Android.Enabled = false
	conf.Android.ApiKey = ""

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

func LoadConfYaml(confPath string) (ConfYaml, error) {
	var config ConfYaml

	configFile, err := ioutil.ReadFile(confPath)

	if err != nil {
		log.Printf("Unable to read config file, path: '%s'", confPath)

		return config, err
	}

	err = yaml.Unmarshal([]byte(configFile), &config)

	if err != nil {
		log.Printf("Unable to Unmarshal config file '%s'", confPath)

		return config, err
	}

	return config, nil
}
