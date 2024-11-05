package system

import (
	"context"

	sysModel "github.com/oldweipro/gin-admin/model/system"
	"github.com/oldweipro/gin-admin/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderMenuAuthority = initOrderMenu + initOrderAuthority

type initMenuAuthority struct{}

// auto run
func init() {
	system.RegisterInit(initOrderMenuAuthority, &initMenuAuthority{})
}

func (i *initMenuAuthority) MigrateTable(ctx context.Context) (context.Context, error) {
	return ctx, nil // do nothing
}

func (i *initMenuAuthority) TableCreated(ctx context.Context) bool {
	return false // always replace
}

func (i initMenuAuthority) InitializerName() string {
	return "sys_menu_authorities"
}

func (i *initMenuAuthority) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	authorities, ok := ctx.Value(initAuthority{}.InitializerName()).([]sysModel.SysAuthority)
	if !ok {
		return ctx, errors.Wrap(system.ErrMissingDependentContext, "创建 [菜单-权限] 关联失败, 未找到权限表初始化数据")
	}
	menus, ok := ctx.Value(initMenu{}.InitializerName()).([]sysModel.SysBaseMenu)
	if !ok {
		return next, errors.Wrap(errors.New(""), "创建 [菜单-权限] 关联失败, 未找到菜单表初始化数据")
	}
	next = ctx
	// 1
	if err = db.Model(&authorities[0]).Association("SysBaseMenus").Replace(menus); err != nil {
		return next, err
	}

	// 2
	menu2 := menus[:2]
	menu2 = append(menu2, menus[3])
	menu2 = append(menu2, menus[4])
	menu2 = append(menu2, menus[5])
	if err = db.Model(&authorities[1]).Association("SysBaseMenus").Replace(menu2); err != nil {
		return next, err
	}

	// 3
	if err = db.Model(&authorities[2]).Association("SysBaseMenus").Replace(menus[:5]); err != nil {
		return next, err
	}
	if err = db.Model(&authorities[2]).Association("SysBaseMenus").Append(menus[9:12]); err != nil {
		return next, err
	}
	return next, nil
}

func (i *initMenuAuthority) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	auth := &sysModel.SysAuthority{}
	if ret := db.Model(auth).
		Where("authority_id = ?", 3).Preload("SysBaseMenus").Find(auth); ret != nil {
		if ret.Error != nil {
			return false
		}
		return len(auth.SysBaseMenus) > 0
	}
	return false
}