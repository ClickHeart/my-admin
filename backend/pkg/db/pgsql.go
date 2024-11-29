package db

import (
	"fmt"
	"my-admin/pkg/config"

	"github.com/donnie4w/go-logger/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PgsqlInit(cfg *config.Pgsql) *gorm.DB {
	if !cfg.Enable {
		return nil
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", cfg.Host, cfg.User, cfg.Password, cfg.Db, cfg.Port)
	logger.Debugf("dsn: %s", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatalf("connect to postgres error: %v", err)
		panic(err)
	}
	return db
}
