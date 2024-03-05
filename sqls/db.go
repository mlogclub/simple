package sqls

import (
	"database/sql"
	"log/slog"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type GormModel struct {
	Id int64 `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
}

type DbConfig struct {
	Url                    string `yaml:"Url"`
	MaxIdleConns           int    `yaml:"MaxIdleConns"`
	MaxOpenConns           int    `yaml:"MaxOpenConns"`
	ConnMaxIdleTimeSeconds int    `yaml:"ConnMaxIdleTimeSeconds"`
	ConnMaxLifetimeSeconds int    `yaml:"ConnMaxLifetimeSeconds"`
}

var (
	db    *gorm.DB
	sqlDB *sql.DB
)

func Open(dbConfig DbConfig, config *gorm.Config, models ...interface{}) (err error) {
	if config == nil {
		config = &gorm.Config{}
	}

	if config.NamingStrategy == nil {
		config.NamingStrategy = schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,
		}
	}

	if db, err = gorm.Open(mysql.Open(dbConfig.Url), config); err != nil {
		slog.Error("opens database failed", slog.Any("error", err))
		return
	}

	if sqlDB, err = db.DB(); err == nil {
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
		sqlDB.SetConnMaxIdleTime(time.Duration(dbConfig.ConnMaxIdleTimeSeconds) * time.Second)
		sqlDB.SetConnMaxLifetime(time.Duration(dbConfig.ConnMaxLifetimeSeconds) * time.Second)
	} else {
		slog.Error(err.Error(), slog.Any("error", err))
	}

	if err = db.AutoMigrate(models...); nil != err {
		slog.Error("auto migrate tables failed", slog.Any("error", err))
	}
	return
}

func DB() *gorm.DB {
	return db
}

func Close() {
	if sqlDB == nil {
		return
	}
	if err := sqlDB.Close(); nil != err {
		slog.Error("Disconnect from database failed", slog.Any("error", err))
	}
}
