package webserver

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Config ...
type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	LogLevel string `json:"log_level"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Host:     "localhost",
		Port:     ":8080",
		Password: "12345678",
		LogLevel: "debug",
	}
}

// ConfigInit ...
func (config *Config) ConfigInit(path string) {
	b, err := ioutil.ReadFile("../../configs/" + path)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(b, &config)
	if err != nil {
		log.Println(err)
	}
}
