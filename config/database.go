package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/soteras/expenses/internal/auth/model"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Dsn         string
	DbType      string
	Debug       bool
	AutoMigrate bool
}

func init() {
	LoadEnv()
}

func NewDb() *Database {
	fmt.Println(GetEnv("DB_TYPE"))

	return &Database{
		Dsn:         GetEnv("DB_DSN"),
		DbType:      GetEnv("DB_TYPE"),
		Debug:       GetEnv("DB_SHOW_QUERIES") == "true",
		AutoMigrate: GetEnv("DB_AUTO_MIGRATE") == "true",
	}
}

func (d *Database) Connect() (*gorm.DB, error) {
	var err error
	var db *gorm.DB
	config := &gorm.Config{}

	fmt.Println(d.DbType)

	if d.Debug {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Silent,
				IgnoreRecordNotFoundError: true,
				ParameterizedQueries:      true,
				Colorful:                  false,
			},
		)

		config = &gorm.Config{
			Logger: newLogger,
		}
	}

	switch d.DbType {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(d.Dsn), config)

	default:
		db, err = gorm.Open(postgres.Open(d.Dsn), config)
	}

	if d.AutoMigrate {
		db.AutoMigrate(&model.User{})
	}

	return db, err
}
