package transaction

import (
	"github.com/oldweipro/gin-admin/global"
)

// SubscriptionPlan 结构体
type SubscriptionPlan struct {
	global.Model
	Name        string `json:"name" form:"name" gorm:"column:name;comment:表示订阅计划的名称。;"`
	Description string `json:"description" form:"description" gorm:"column:description;comment:表示订阅计划的详细描述信息。;"`
	Price       *uint  `json:"price" form:"price" gorm:"column:price;comment:表示订阅计划的价格。;"`
	Duration    *uint  `json:"duration" form:"duration" gorm:"column:duration;comment:表示订阅计划的时长，用秒表示;"`
	Quantity    *uint  `json:"quantity" form:"quantity" gorm:"column:quantity;comment:表示订阅计划的数量;"`
	MenuId      *uint  `json:"menuId" form:"menuId" gorm:"column:menu_id;comment:关联的权限ID;"`
	Tag         *uint  `json:"tag" form:"tag" gorm:"column:tag;comment:功能模块标识。;"`
	Status      *uint  `json:"status" form:"status" gorm:"column:status;comment:是否开启，1启用，0禁用。;"`
}

// TableName SubscriptionPlan 表名
func (SubscriptionPlan) TableName() string {
	return "subscription_plan"
}
