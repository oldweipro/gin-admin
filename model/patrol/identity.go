// 自动生成模板Identity
package patrol

import (
	"github.com/oldweipro/gin-admin/global"
)

// Identity 结构体
type Identity struct {
	global.GVA_MODEL
	IdCard    string `json:"id_card" form:"id_card" gorm:"column:id_card;comment:身份证;"`
	RealName  string `json:"real_name" form:"real_name" gorm:"column:real_name;comment:真实姓名;"`
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName Identity 表名
func (Identity) TableName() string {
	return "identity"
}
