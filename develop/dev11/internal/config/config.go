package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Host string
	Port string
}

func ReadConfigFile(filename string) (Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}
	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
