package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	LiveId        []int64 `json:"live_id"`
	FetchInterval int64   `json:"fetch_interval"`
}

func LoadConfig(file string) (Config,error) {
	var config Config
	configFile, err := ioutil.ReadFile(file)
	if err != nil {
		return Config{},err
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return Config{},err
	}
	return config,nil
}
