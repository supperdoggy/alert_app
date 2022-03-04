package cfg

import (
	"encoding/json"
	"flag"
	"io/ioutil"
)

var config appConfig

type appConfig struct {
	DBUrl                string `json:"db_url"`
	DBName               string `json:"db_name"`
	AlertsCollectionName string `json:"alerts_collection_name"`

	Url  string `json:"url"`
	Port string `json:"port"`

	IsProd bool `json:"is_prod"`
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

	if config.DBName == "" ||
		config.AlertsCollectionName == "" ||
		config.Port == "" {
		panic("fill all the config values!!!!")
	}
}

func GetConfig() appConfig {
	return config
}
