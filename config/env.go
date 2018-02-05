package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type envConfig struct {
	VkGroupFrom         []string `envconfig:"VK_GROUP_FROM"`
	VkGroupTo           string   `envconfig:"VK_GROUP_TO"`
	RefreshTimePerGroup int      `envconfig:"REFRESH_TIME_PER_GROUP"`
}

var c = new(envConfig)

func init() {
	err := envconfig.Process("", c)
	if err != nil {
		log.Fatalln("Enviroment error:", err.Error())
	}
}

func GetConfig() *envConfig {
	return c
}
