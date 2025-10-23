package sqls

import (
	"gorm.io/gorm"
)

type GormModel struct {
	Id int64 `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
}

type DbConfig struct {
	Url                    string `yaml:"url"`
	MaxIdleConns           int    `yaml:"maxIdleConns"`
	MaxOpenConns           int    `yaml:"maxOpenConns"`
	ConnMaxIdleTimeSeconds int    `yaml:"connMaxIdleTimeSeconds"`
	ConnMaxLifetimeSeconds int    `yaml:"connMaxLifetimeSeconds"`
}

var (
	_db *gorm.DB
)

func DB() *gorm.DB {
	return _db
}

func SetDB(gormDB *gorm.DB) {
	_db = gormDB
}
