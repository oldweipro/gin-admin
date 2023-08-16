package transaction

import (
	"github.com/oldweipro/gin-admin/global"
	"time"
)

// SubscriptionUser 结构体
type SubscriptionUser struct {
	global.Model
	SubscriptionPlanId uint      `json:"subscriptionPlanId" form:"subscriptionPlanId" gorm:"column:subscription_plan_id;comment:订阅ID。;"`
	UserId             uint      `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID。;"`
	StartTime          time.Time `json:"startTime" form:"startTime" gorm:"column:start_time;comment:开始时间。;"`
	EndTime            time.Time `json:"endTime" form:"endTime" gorm:"column:end_time;comment:结束时间，根据续费修改时间，定时任务刷新结束时间，及时去掉相关权限。;"`
}

// TableName SubscriptionUser 表名
func (SubscriptionUser) TableName() string {
	return "subscription_user"
}
