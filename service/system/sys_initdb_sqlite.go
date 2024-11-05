package system

import (
	"context"
	"errors"
	"github.com/glebarez/sqlite"
	"github.com/gofrs/uuid/v5"
	"github.com/gookit/color"
	"github.com/oldweipro/gin-admin/pkg/app"
	"github.com/oldweipro/gin-admin/pkg/utils"
	"gorm.io/gorm"
	"path/filepath"

	"github.com/oldweipro/gin-admin/model/system/request"
	"github.com/oldweipro/gin-admin/pkg/config"
)

type SqliteInitHandler struct{}

func NewSqliteInitHandler() *SqliteInitHandler {
	return &SqliteInitHandler{}
}

// WriteConfig mysql回写配置
func (h SqliteInitHandler) WriteConfig(ctx context.Context) error {
	c, ok := ctx.Value("config").(config.Sqlite)
	if !ok {
		return errors.New("sqlite config invalid")
	}
	app.Config.System.DbType = "sqlite"
	app.Config.Sqlite = c
	app.Config.JWT.SigningKey = uuid.Must(uuid.NewV4()).String()
	cs := utils.StructToMap(app.Config)
	for k, v := range cs {
		app.Viper.Set(k, v)
	}
	app.ActiveDBName = &c.Dbname
	return app.Viper.WriteConfig()
}

// EnsureDB 创建数据库并初始化 sqlite
func (h SqliteInitHandler) EnsureDB(ctx context.Context, conf *request.InitDB) (next context.Context, err error) {
	if s, ok := ctx.Value("dbType").(string); !ok || s != "sqlite" {
		return ctx, ErrDBTypeMismatch
	}

	c := conf.ToSqliteConfig()
	next = context.WithValue(ctx, "config", c)
	if c.Dbname == "" {
		return ctx, nil
	} // 如果没有数据库名, 则跳出初始化数据

	dsn := conf.SqliteEmptyDsn()

	var db *gorm.DB
	if db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	}); err != nil {
		return ctx, err
	}
	app.Config.AutoCode.Root, _ = filepath.Abs("..")
	next = context.WithValue(next, "db", db)
	return next, err
}

func (h SqliteInitHandler) InitTables(ctx context.Context, inits initSlice) error {
	return createTables(ctx, inits)
}

func (h SqliteInitHandler) InitData(ctx context.Context, inits initSlice) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) { c() }(cancel)
	for _, init := range inits {
		if init.DataInserted(next) {
			color.Info.Printf(InitDataExist, Sqlite, init.InitializerName())
			continue
		}
		if n, err := init.InitializeData(next); err != nil {
			color.Info.Printf(InitDataFailed, Sqlite, init.InitializerName(), err)
			return err
		} else {
			next = n
			color.Info.Printf(InitDataSuccess, Sqlite, init.InitializerName())
		}
	}
	color.Info.Printf(InitSuccess, Sqlite)
	return nil
}
