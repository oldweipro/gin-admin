package initialize

import (
	"os"

	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/example"
	"github.com/oldweipro/gin-admin/model/system"

	"github.com/oldweipro/gin-admin/model/ladder"
	"github.com/oldweipro/gin-admin/model/openfish"
	"github.com/oldweipro/gin-admin/model/patrol"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	case "oracle":
		return GormOracle()
	case "mssql":
		return GormMssql()
	default:
		return GormMysql()
	}
}

func RegisterTables() {
	db := global.GVA_DB
	err := db.AutoMigrate(

		system.SysApi{},
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
		system.SysAutoCode{},
		system.SysChatGptOption{},

		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{},

		patrol.PatrolItem{},
		patrol.PatrolTask{},
		patrol.PatrolSite{},
		patrol.Identity{},
		patrol.Account{},
		patrol.CertificationRecord{},
		patrol.Personnel{}, ladder.ServerNode{}, ladder.Inbounds{}, openfish.Conversation{}, openfish.ChatTicket{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
