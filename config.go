package main

import (
	"log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
}

type SectionApi struct {
	PushUri      string `yaml:"push_uri"`
	StatGoUri    string `yaml:"stat_go_uri"`
}

type SectionAndroid struct {
	Enabled  bool   `yaml:"enabled"`
	ApiKey   string `yaml:"apikey"`
}

type SectionIos struct {
	Enabled              bool   `yaml:"enabled"`
	PemCertPath          string `yaml:"pem_cert_path"`
	PemKeyPath           string `yaml:"pem_key_path"`
	Production           bool   `yaml:"production"`
}

func BuildDefaultPushConf() ConfYaml {
	var conf ConfYaml

	// Core
	conf.Core.Port = "8088"
	conf.Core.NotificationMax = 100

	// Api
	conf.Api.PushUri = "/api/push"
	conf.Api.StatGoUri = "/api/status"

	// Android
	conf.Android.ApiKey = ""
	conf.Android.Enabled = true

	// iOS
	conf.Ios.Enabled = true
	conf.Ios.PemCertPath = ""
	conf.Ios.PemKeyPath = ""
	conf.Ios.Production = false

	return conf
}

func LoadConfYaml(confPath string) (ConfYaml, error) {
	var config ConfYaml

	configFile, err := ioutil.ReadFile(confPath)

	if err != nil {
		log.Printf("Unable to read config file '%s'", confPath)

		return config, err
	}

	err = yaml.Unmarshal([]byte(configFile), &config)
	if err != nil {
		log.Printf("Unable to read config file '%v'", err)

		return config, err
	}

	return config, nil
}
