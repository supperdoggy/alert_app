package cfg

import (
	"encoding/json"
	"flag"
	"io/ioutil"
)

var config appConfig

type appConfig struct {
	NotificationServiceURL string `json:"notification_service_url"`
	NotificationServicePort string`json:"notification_service_port"`

	DBServiceURL string `json:"db_url"`
	DBServicePort string `json:"db_port"`

}

func init() {
	path := flag.String("config", "", "for config path")
	flag.Parse()
	if path == nil || *path == "" {
		panic("no config path specified")
	}

	data, err := ioutil.ReadFile(*path)
	if err != nil {
		panic(err.Error())
	}

	if err := json.Unmarshal(data, &config); err != nil {
		panic(err.Error())
	}

	if config.DBServiceURL == "" ||
		config.DBServicePort == "" ||
		config.NotificationServiceURL == "" ||
		config.NotificationServicePort == "" {
		panic("fill all the config values!!!!")
	}
}

func GetConfig() appConfig {
	return config
}
