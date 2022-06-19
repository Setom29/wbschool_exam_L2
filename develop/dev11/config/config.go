package config

import (
	"encoding/json"
	"io/ioutil"
)

type Settings struct {
	Host     string
	Port     string
	PgHost   string
	PgPort   string
	PgUser   string
	PgPasswd string
	PgBase   string
}

func Configure() Settings {
	data, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		panic("Could not read configuration file")
	}
	var cfg Settings
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		panic("Could not parse configuration file")
	}
	return cfg
}
