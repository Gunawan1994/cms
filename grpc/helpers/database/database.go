package database

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"gorm.io/gorm/schema"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	DbHost string
	DbUser string
	DbPass string
	DbName string
	DbPort string
}

type Database struct {
	db *gorm.DB
}

func (d *Database) GetDB() *gorm.DB {
	return d.db
}

func NewDatabase(driver string, cfg *Config) *Database {
	var db *gorm.DB
	var err error
	var dialect gorm.Dialector

	switch driver {
	case "postgres", "pgsql":
		dialect = postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", cfg.DbHost, cfg.DbUser, cfg.DbPass, cfg.DbName, cfg.DbPort))
	default:
		slog.Warn("unknown database driver")
		os.Exit(1)
	}

	for {
		configGorm := &gorm.Config{
			Logger:          logger.Default.LogMode(logger.Info),
			NowFunc:         time.Now().UTC,
			CreateBatchSize: 1000,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		}
		db, err = gorm.Open(dialect, configGorm)
		if err != nil {
			slog.Error(fmt.Sprintf("failed to connect to %s database", driver), "error", err.Error())
			slog.Info(fmt.Sprintf("retrying to connect to %s database in 5 seconds...", driver))
			time.Sleep(5 * time.Second)
			continue
		}
		slog.Info(fmt.Sprintf("successfully connected to %s database", driver))
		break
	}

	sqlDB, err := db.DB()
	if err != nil {
		slog.Error("failed to configure connection pool", "error", err.Error())
		os.Exit(1)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(30 * time.Minute)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return &Database{db: db}
}
