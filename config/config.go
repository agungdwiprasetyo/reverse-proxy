package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/agungdwiprasetyo/reverse-proxy/helper"
)

// Config app
type Config struct {
	GatewayPort int `json:"gatewayPort"`
	Proxy       []struct {
		Root string `json:"root"`
		Host string `json:"host"`
	} `json:"proxy"`
	Key struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"key"`
}

// Init config
func Init(appPath string) *Config {
	b, err := ioutil.ReadFile(fmt.Sprintf("%s/config.json", appPath))
	if err != nil {
		log.Fatal(helper.StringRed(err))
	}

	var config Config
	if err := json.Unmarshal(b, &config); err != nil {
		log.Fatal(helper.StringRed(err))
	}

	return &config
}
