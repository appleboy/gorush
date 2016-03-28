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
}

type SectionCore struct {
	Port            string `yaml:"port"`
	NotificationMax int    `yaml:"notification_max"`
	Production      bool   `yaml:"production"`
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

func BuildDefaultPushConf() ConfYaml {
	var conf ConfYaml

	// Core
	conf.Core.Port = "8088"
	conf.Core.NotificationMax = 100
	conf.Core.Production = true
	conf.Core.SSL = false
	conf.Core.CertPath = "cert.pem"
	conf.Core.KeyPath = "key.pem"

	// Api
	conf.Api.PushUri = "/api/push"
	conf.Api.StatGoUri = "/api/status"

	// Android
	conf.Android.ApiKey = ""
	conf.Android.Enabled = false

	// iOS
	conf.Ios.Enabled = false
	conf.Ios.PemCertPath = "cert.pem"
	conf.Ios.PemKeyPath = "key.pem"
	conf.Ios.Production = false

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
