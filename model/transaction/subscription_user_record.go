package transaction

import (
	"github.com/oldweipro/gin-admin/global"
	"time"
)

// SubscriptionUserRecord 结构体 对订阅交易前后的数据记录
type SubscriptionUserRecord struct {
	global.Model
	SubscriptionUserId *uint     `json:"subscriptionUserId" form:"subscriptionUserId" gorm:"column:subscription_user_id;comment:订阅用户关联ID。;"`
	UserId             *uint     `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID。;"`
	SubscriptionPlanId *uint     `json:"subscriptionPlanId" form:"subscriptionPlanId" gorm:"column:subscription_plan_id;comment:订阅计划ID。;"`
	BeforeEndTime      time.Time `json:"beforeEndTime" form:"beforeEndTime" gorm:"column:before_end_time;comment:操作之前的结束时间。;"`
	AfterEndTime       time.Time `json:"afterEndTime" form:"afterEndTime" gorm:"column:after_end_time;comment:操作之后的结束时间。;"`
	Name               string    `json:"name" form:"name" gorm:"column:name;comment:表示订阅计划的名称。;"`
	Description        string    `json:"description" form:"description" gorm:"column:description;comment:表示订阅计划的详细描述信息。;"`
	Price              *uint     `json:"price" form:"price" gorm:"column:price;comment:表示订阅计划的价格。;"`
	Duration           *uint     `json:"duration" form:"duration" gorm:"column:duration;comment:表示订阅计划的有效期。天数表示;"`
	Quantity           *uint     `json:"quantity" form:"quantity" gorm:"column:quantity;comment:表示订阅计划的数量;"`
	MenuId             *uint     `json:"menuId" form:"menuId" gorm:"column:menu_id;comment:关联的权限ID;"`
}

// TableName SubscriptionUserRecord 表名
func (SubscriptionUserRecord) TableName() string {
	return "subscription_user_record"
}
