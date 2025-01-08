package global

import (
	"my-admin/pkg/config"
	"my-admin/pkg/db"
	"my-admin/pkg/log"

	"gorm.io/gorm"
)

var (
	Cfg *config.Config
	DB  *gorm.DB
)

func init() {
	Cfg = config.Init()
	log.Init(&Cfg.Log)
	DB = db.PgsqlInit(&Cfg.Data.Pgsql)
}
