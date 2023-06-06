package database

import (
	"fmt"
	"github.com/mazzama/todo-grpc/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
	"time"
)

func InitPostgres() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user='%s' password='%s' dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.GetConfigString(config.PostgresHost),
		config.GetConfigString(config.PostgresUser),
		config.GetConfigString(config.PostgresPassword),
		config.GetConfigString(config.PostgresDBName),
		config.GetConfigString(config.PostgresPort),
		config.GetConfigString(config.PostgresSSLMode),
		config.GetConfigString(config.AppTimezone),
	)

	var logLevel logger.LogLevel
	switch strings.ToUpper(config.GetConfigString(config.PostgresLogLevel)) {
	case "SILENT":
		logLevel = logger.Silent
	case "INFO":
		logLevel = logger.Info
	case "WARN":
		logLevel = logger.Warn
	default:
		logLevel = logger.Error
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	maxLifetime := config.GetConfigInt(config.PostgresConnectionMaxLifetime)

	maxOpenConns := config.GetConfigInt(config.PostgresConnectionMaxOpen)

	maxIdleConns := config.GetConfigInt(config.PostgresConnectionMaxIdle)

	sqlDB.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Second)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetMaxIdleConns(maxIdleConns)

	return db, nil
}
