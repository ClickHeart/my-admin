package global

import (
	"my-admin/pkg/config"
	"my-admin/pkg/db"
	"my-admin/pkg/mlog"

	"gorm.io/gorm"
)

var (
	Cfg *config.Config
	Db  *gorm.DB
)

func init() {
	Cfg = config.Init()
	mlog.Init(&Cfg.Log)
	Db = db.PgsqlInit(&Cfg.Data.Pgsql)
}
