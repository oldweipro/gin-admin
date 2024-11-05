package initialize

import (
	"github.com/oldweipro/gin-admin/pkg/app"
	"os"

	"github.com/oldweipro/gin-admin/model/example"
	"github.com/oldweipro/gin-admin/model/system"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewGorm() *gorm.DB {
	switch app.Config.System.DbType {
	case "mysql":
		app.ActiveDBName = &app.Config.Mysql.Dbname
		return Mysql()
	case "pgsql":
		app.ActiveDBName = &app.Config.Pgsql.Dbname
		return PgSql()
	case "oracle":
		app.ActiveDBName = &app.Config.Oracle.Dbname
		return Oracle()
	case "mssql":
		app.ActiveDBName = &app.Config.Mssql.Dbname
		return Mssql()
	case "sqlite":
		app.ActiveDBName = &app.Config.Sqlite.Dbname
		return Sqlite()
	default:
		app.ActiveDBName = &app.Config.Mysql.Dbname
		return Mysql()
	}
}

func RegisterTables() {
	db := app.DBClient
	err := db.AutoMigrate(

		system.SysApi{},
		system.SysIgnoreApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysAutoCodeHistory{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
		system.SysAutoCodePackage{},
		system.SysExportTemplate{},
		system.Condition{},
		system.JoinTemplate{},

		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{},
	)
	if err != nil {
		app.Logger.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}

	err = bizModel()

	if err != nil {
		app.Logger.Error("register biz_table failed", zap.Error(err))
		os.Exit(0)
	}
	app.Logger.Info("register table success")
}
