package system

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/oldweipro/gin-admin/pkg/app"
	"github.com/oldweipro/gin-admin/pkg/utils"
	ast2 "github.com/oldweipro/gin-admin/pkg/utils/ast"
	"github.com/pkg/errors"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	common "github.com/oldweipro/gin-admin/model/common/request"
	model "github.com/oldweipro/gin-admin/model/system"
	request "github.com/oldweipro/gin-admin/model/system/request"
	"go.uber.org/zap"
)

var AutocodeHistory = new(autoCodeHistory)

type autoCodeHistory struct{}

// Create 创建代码生成器历史记录
// Author [oldweipro](https://github.com/oldweipro)
func (s *autoCodeHistory) Create(ctx context.Context, info request.SysAutoHistoryCreate) error {
	create := info.Create()
	err := app.DBClient.WithContext(ctx).Create(&create).Error
	if err != nil {
		return errors.Wrap(err, "创建失败!")
	}
	return nil
}

// First 根据id获取代码生成器历史的数据
// Author [oldweipro](https://github.com/oldweipro)
func (s *autoCodeHistory) First(ctx context.Context, info common.GetById) (string, error) {
	var meta string
	err := app.DBClient.WithContext(ctx).Model(model.SysAutoCodeHistory{}).Where("id = ?", info.ID).Pluck("request", &meta).Error
	if err != nil {
		return "", errors.Wrap(err, "获取失败!")
	}
	return meta, nil
}

// Repeat 检测重复
// Author [oldweipro](https://github.com/oldweipro)
func (s *autoCodeHistory) Repeat(businessDB, structName, Package string) bool {
	var count int64
	app.DBClient.Model(&model.SysAutoCodeHistory{}).Where("business_db = ? and struct_name = ? and package = ? and flag = 0", businessDB, structName, Package).Count(&count)
	return count > 0
}

// RollBack 回滚
// Author [oldweipro](https://github.com/oldweipro)
func (s *autoCodeHistory) RollBack(ctx context.Context, info request.SysAutoHistoryRollBack) error {
	var history model.SysAutoCodeHistory
	err := app.DBClient.Where("id = ?", info.ID).First(&history).Error
	if err != nil {
		return err
	}
	if history.ExportTemplateID != 0 {
		err = app.DBClient.Delete(&model.SysExportTemplate{}, "id = ?", history.ExportTemplateID).Error
		if err != nil {
			return err
		}
	}
	if info.DeleteApi {
		ids := info.ApiIds(history)
		err = ApiServiceApp.DeleteApisByIds(ids)
		if err != nil {
			app.Logger.Error("ClearTag DeleteApiByIds:", zap.Error(err))
		}
	} // 清除API表
	if info.DeleteMenu {
		err = BaseMenuServiceApp.DeleteBaseMenu(int(history.MenuID))
		if err != nil {
			return errors.Wrap(err, "删除菜单失败!")
		}
	} // 清除菜单表
	if info.DeleteTable {
		err = s.DropTable(history.BusinessDB, history.Table)
		if err != nil {
			return errors.Wrap(err, "删除表失败!")
		}
	} // 删除表
	templates := make(map[string]string, len(history.Templates))
	for key, template := range history.Templates {
		{
			server := filepath.Join(app.Config.AutoCode.Root, app.Config.AutoCode.Server)
			keys := strings.Split(key, "/")
			key = filepath.Join(keys...)
			key = strings.TrimPrefix(key, server)
		} // key
		{
			web := filepath.Join(app.Config.AutoCode.Root, app.Config.AutoCode.WebRoot())
			server := filepath.Join(app.Config.AutoCode.Root, app.Config.AutoCode.Server)
			slices := strings.Split(template, "/")
			template = filepath.Join(slices...)
			ext := path.Ext(template)
			switch ext {
			case ".js", ".vue":
				template = filepath.Join(web, template)
			case ".go":
				template = filepath.Join(server, template)
			}
		} // value
		templates[key] = template
	}
	history.Templates = templates
	for key, value := range history.Injections {
		var injection ast2.Ast
		switch key {
		case ast2.TypePackageApiEnter, ast2.TypePackageRouterEnter, ast2.TypePackageServiceEnter:

		case ast2.TypePackageApiModuleEnter, ast2.TypePackageRouterModuleEnter, ast2.TypePackageServiceModuleEnter:
			var entity ast2.PackageModuleEnter
			_ = json.Unmarshal([]byte(value), &entity)
			injection = &entity
		case ast2.TypePackageInitializeGorm:
			var entity ast2.PackageInitializeGorm
			_ = json.Unmarshal([]byte(value), &entity)
			injection = &entity
		case ast2.TypePackageInitializeRouter:
			var entity ast2.PackageInitializeRouter
			_ = json.Unmarshal([]byte(value), &entity)
			injection = &entity
		}
		if injection == nil {
			continue
		}
		file, _ := injection.Parse("", nil)
		if file != nil {
			_ = injection.Rollback(file)
			err = injection.Format("", nil, file)
			if err != nil {
				return err
			}
			fmt.Printf("[filepath:%s]回滚注入代码成功!\n", key)
		}
	} // 清除注入代码
	removeBasePath := filepath.Join(app.Config.AutoCode.Root, "rm_file", strconv.FormatInt(int64(time.Now().Nanosecond()), 10))
	for _, value := range history.Templates {
		if !filepath.IsAbs(value) {
			continue
		}
		removePath := filepath.Join(removeBasePath, strings.TrimPrefix(value, app.Config.AutoCode.Root))
		err = utils.FileMove(value, removePath)
		if err != nil {
			return errors.Wrapf(err, "[src:%s][dst:%s]文件移动失败!", value, removePath)
		}
	} // 移动文件
	err = app.DBClient.WithContext(ctx).Model(&model.SysAutoCodeHistory{}).Where("id = ?", info.ID).Update("flag", 1).Error
	if err != nil {
		return errors.Wrap(err, "更新失败!")
	}
	return nil
}

// Delete 删除历史数据
// Author [oldweipro](https://github.com/oldweipro)
func (s *autoCodeHistory) Delete(ctx context.Context, info common.GetById) error {
	err := app.DBClient.WithContext(ctx).Where("id = ?", info.Uint()).Delete(&model.SysAutoCodeHistory{}).Error
	if err != nil {
		return errors.Wrap(err, "删除失败!")
	}
	return nil
}

// GetList 获取系统历史数据
// Author [oldweipro](https://github.com/oldweipro)
func (s *autoCodeHistory) GetList(ctx context.Context, info common.PageInfo) (list []model.SysAutoCodeHistory, total int64, err error) {
	var entities []model.SysAutoCodeHistory
	db := app.DBClient.WithContext(ctx).Model(&model.SysAutoCodeHistory{})
	err = db.Count(&total).Error
	if err != nil {
		return nil, total, err
	}
	err = db.Scopes(info.Paginate()).Order("updated_at desc").Find(&entities).Error
	return entities, total, err
}

// DropTable 获取指定数据库和指定数据表的所有字段名,类型值等
// @author: [oldweipro](https://github.com/oldweipro)
func (s *autoCodeHistory) DropTable(BusinessDb, tableName string) error {
	if BusinessDb != "" {
		return app.MustGetGlobalDBByDBName(BusinessDb).Exec("DROP TABLE " + tableName).Error
	} else {
		return app.DBClient.Exec("DROP TABLE " + tableName).Error
	}
}
