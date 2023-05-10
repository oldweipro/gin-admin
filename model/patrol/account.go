// 自动生成模板Account
package patrol

import (
	"github.com/oldweipro/gin-admin/global"
)

// Account 结构体
type Account struct {
	global.GVA_MODEL
	AccountName  string `json:"accountName" form:"accountName" gorm:"column:account_name;comment:9377游戏的用户名和密码;"`
	LoginStatus  *int   `json:"loginStatus" gorm:"default:0;column:login_status;comment:登陆状态;"`
	CurrentCalls *int   `json:"currentCalls" gorm:"default:0;column:current_calls;comment:当前调用次数，每天12点清0;"`
	CreatedBy    uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy    uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy    uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName Account 表名
func (Account) TableName() string {
	return "account"
}
