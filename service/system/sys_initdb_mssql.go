package system

import (
	"context"
	"errors"
	"github.com/gofrs/uuid/v5"
	"github.com/gookit/color"
	"github.com/oldweipro/gin-admin/model/system/request"
	"github.com/oldweipro/gin-admin/pkg/app"
	"github.com/oldweipro/gin-admin/pkg/config"
	"github.com/oldweipro/gin-admin/pkg/utils"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"path/filepath"
)

type MssqlInitHandler struct{}

func NewMssqlInitHandler() *MssqlInitHandler {
	return &MssqlInitHandler{}
}

// WriteConfig mssql回写配置
func (h MssqlInitHandler) WriteConfig(ctx context.Context) error {
	c, ok := ctx.Value("config").(config.Mssql)
	if !ok {
		return errors.New("mssql config invalid")
	}
	app.Config.System.DbType = "mssql"
	app.Config.Mssql = c
	app.Config.JWT.SigningKey = uuid.Must(uuid.NewV4()).String()
	cs := utils.StructToMap(app.Config)
	for k, v := range cs {
		app.Viper.Set(k, v)
	}
	app.ActiveDBName = &c.Dbname
	return app.Viper.WriteConfig()
}

// EnsureDB 创建数据库并初始化 mssql
func (h MssqlInitHandler) EnsureDB(ctx context.Context, conf *request.InitDB) (next context.Context, err error) {
	if s, ok := ctx.Value("dbType").(string); !ok || s != "mssql" {
		return ctx, ErrDBTypeMismatch
	}

	c := conf.ToMssqlConfig()
	next = context.WithValue(ctx, "config", c)
	if c.Dbname == "" {
		return ctx, nil
	} // 如果没有数据库名, 则跳出初始化数据

	dsn := conf.MssqlEmptyDsn()

	mssqlConfig := sqlserver.Config{
		DSN:               dsn, // DSN data source name
		DefaultStringSize: 191, // string 类型字段的默认长度
	}

	var db *gorm.DB

	if db, err = gorm.Open(sqlserver.New(mssqlConfig), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return nil, err
	}

	app.Config.AutoCode.Root, _ = filepath.Abs("..")
	next = context.WithValue(next, "db", db)
	return next, err
}

func (h MssqlInitHandler) InitTables(ctx context.Context, inits initSlice) error {
	return createTables(ctx, inits)
}

func (h MssqlInitHandler) InitData(ctx context.Context, inits initSlice) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) { c() }(cancel)
	for _, init := range inits {
		if init.DataInserted(next) {
			color.Info.Printf(InitDataExist, Mssql, init.InitializerName())
			continue
		}
		if n, err := init.InitializeData(next); err != nil {
			color.Info.Printf(InitDataFailed, Mssql, init.InitializerName(), err)
			return err
		} else {
			next = n
			color.Info.Printf(InitDataSuccess, Mssql, init.InitializerName())
		}
	}
	color.Info.Printf(InitSuccess, Mssql)
	return nil
}
