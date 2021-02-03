package config

import (
	"github.com/kelseyhightower/envconfig"
)

type logConfig struct {
	File string `required:"true" envconfig:"LOG_FILE"`
}

type serverConfig struct {
	Port string `required:"true" envconfig:"SERVER_PORT"`
}

type databaseConfig struct {
	User     string `required:"true" envconfig:"DATABASE_USER"`
	Password string `required:"true" envconfig:"DATABASE_PASSWORD"`
	Host     string `required:"true" envconfig:"DATABASE_HOST"`
	Port     string `required:"true" envconfig:"DATABASE_PORT"`
	DBNAME   string `required:"true" envconfig:"DATABASE_NAME"`
}

var (
	Log      *logConfig
	Server   *serverConfig
	Database *databaseConfig
)

func init() {
	Log = new(logConfig)
	Server = new(serverConfig)
	Database = new(databaseConfig)

	//func MustProcess panics if an error occurs. e.g., no such environment variable
	//for readability, prefixes are not used
	envconfig.MustProcess("", Log)
	envconfig.MustProcess("", Server)
	envconfig.MustProcess("", Database)
}
