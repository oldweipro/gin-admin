package internal

import (
	"github.com/oldweipro/gin-admin/pkg/app"
	"github.com/oldweipro/gin-admin/pkg/config"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var Gorm = new(_gorm)

type _gorm struct{}

// Config gorm 自定义配置
// Author [oldweipro](https://github.com/oldweipro)
func (g *_gorm) Config(prefix string, singular bool) *gorm.Config {
	var general config.GeneralDB
	switch app.Config.System.DbType {
	case "mysql":
		general = app.Config.Mysql.GeneralDB
	case "pgsql":
		general = app.Config.Pgsql.GeneralDB
	case "oracle":
		general = app.Config.Oracle.GeneralDB
	case "sqlite":
		general = app.Config.Sqlite.GeneralDB
	case "mssql":
		general = app.Config.Mssql.GeneralDB
	default:
		general = app.Config.Mysql.GeneralDB
	}
	return &gorm.Config{
		Logger: logger.New(NewWriter(general, log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      general.LogLevel(),
			Colorful:      true,
		}),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}
