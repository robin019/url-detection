package config

import (
	"github.com/kelseyhightower/envconfig"
)

type serverConfig struct {
	Host string `required:"true" envconfig:"SERVER_HOST"`
	Port string `required:"true" envconfig:"SERVER_PORT"`
}

var (
	Server *serverConfig
)

func init() {
	Server = new(serverConfig)

	//func MustProcess panics if an error occurs. e.g., no such environment variable
	envconfig.MustProcess("", Server)
}
