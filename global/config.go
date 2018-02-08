package global

import (
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/lukashenka/vkposter/config"
)

var c = new(config.EnvConfig)

func init() {
	err := envconfig.Process("", c)
	if err != nil {
		log.Fatalln("Enviroment error:", err.Error())
	}
}

func GetConfig() *config.EnvConfig {
	return c
}
