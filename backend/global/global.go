package global

import (
	"naive-backend/util/config"
)

var (
	Cfg *config.Config
)

func init() {
	Cfg = new(config.Config)
}
