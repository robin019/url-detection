package sql

import (
	"fmt"

	"github.com/robin019/url-detection/src/utils/logger"

	"gorm.io/driver/postgres"

	"github.com/robin019/url-detection/src/utils/config"

	"gorm.io/gorm"
)

var (
	db      *gorm.DB
	dialect = fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.DBNAME,
		config.Database.Password,
	)
)

func init() {
	connect(dialect)
}

// DB() returns postgres instance
func DB() *gorm.DB {
	return db
}

func connect(dialect string) {
	conn, err := gorm.Open(postgres.Open(dialect), &gorm.Config{})

	if err != nil {
		logger.ApiLog().Fatal(err.Error())
	}
	db = conn
}
