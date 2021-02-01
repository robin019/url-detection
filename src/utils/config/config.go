package config

import (
	"github.com/kelseyhightower/envconfig"
)

type LogConfig struct {
	File string `required:"true" envconfig:"LOG_FILE"`
}

type serverConfig struct {
	Host string `required:"true" envconfig:"SERVER_HOST"`
	Port string `required:"true" envconfig:"SERVER_PORT"`
}

var (
	Log    *LogConfig
	Server *serverConfig
)

func init() {
	Log = new(LogConfig)
	Server = new(serverConfig)

	//func MustProcess panics if an error occurs. e.g., no such environment variable
	envconfig.MustProcess("", Log) //for readability, prefixes are not used
	envconfig.MustProcess("", Server)
}
