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
	GatewayPort int `json:"gateway_port"`
	Services    []struct {
		Root string `json:"root"`
		Host string `json:"host"`
	} `json:"services"`
	Key struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"key"`
}

// GlobalConfig var
var GlobalConfig Config

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

	GlobalConfig = config

	return &config
}
