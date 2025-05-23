package config

import (
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
	Conn        *gorm.DB
	DbType      string
	Debug       bool
	AutoMigrate bool
}

func NewDb() *Database {
	return &Database{}
}

func (d *Database) Connect() error {
	var err error
	config := &gorm.Config{}
	app := NewApp()

	if GetEnv("DEBUG") == "true" {
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

	if app.ENV == EnvTest {
		d.Conn, err = gorm.Open(sqlite.Open("test.db"), config)

	} else {
		d.Conn, err = gorm.Open(postgres.Open(GetEnv("DB_DSN")), config)
	}

	if GetEnv("DB_AUTO_MIGRATE") == "true" {
		d.Conn.AutoMigrate(&model.User{})
	}

	return err
}
