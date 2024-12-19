package config

import (
	"os"

	"github.com/jeanphorn/log4go"
	"github.com/tkanos/gonfig"
)

type Conf struct {
	DB            string
	IsDebug       bool
	IsQueue       bool
	IsConcurrent  bool
	Secret        string
	GoogleSmtpKey string
	URL           string
	Quota         int
	KafkaUrl      string
	KafkaTopic    string
}

var conf Conf

func Init() {
	err := gonfig.GetConf("config.json", &conf)
	if err != nil {
		log4go.Error(err)
		os.Exit(500)
	}
}

func GetConf() Conf {
	return conf
}
