package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Address string
	MaxIdle int
	MaxOpen int
	Log     bool
}

func NewOpen(cfg Config) (*gorm.DB, error) {
	config := &gorm.Config{}
	if cfg.Log { //info
		config.Logger = logger.Default.LogMode(logger.Info)
	} else { //silent
		config.Logger = logger.Default.LogMode(logger.Silent)
	}
	conn, err := gorm.Open(mysql.Open(cfg.Address), config)
	if err != nil {
		return nil, err
	}
	sqlDB, err := conn.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(cfg.MaxIdle)
	sqlDB.SetMaxOpenConns(cfg.MaxOpen)
	return conn, nil
}
