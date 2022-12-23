package config

import (
	"encoding/json"
	"io/ioutil"
)

type AppConfig struct {
	Env      string      `json:"env"`
	Service  ServiceConf `json:"service"`
	Database DbConf      `json:"database"`
}

type DbConf struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type ServiceConf struct {
	Port string `json:"port"`
}

type RedisConf struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

func InitConfig() *AppConfig {
	file, err := ioutil.ReadFile("./cred/appconfig.{env}.json")
	if err != nil {
		panic(err)
	}

	confData := AppConfig{}

	err = json.Unmarshal([]byte(file), &confData)
	if err != nil {
		panic(err)
	}

	print(confData.Database.Host)
	return &confData
}
