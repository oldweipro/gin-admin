// 自动生成模板PatrolItem
package patrol

import (
	"github.com/oldweipro/gin-admin/global"
)

// PatrolItem 结构体
type PatrolItem struct {
	global.GVA_MODEL
	ItemTitle string `json:"itemTitle" form:"itemTitle" gorm:"column:item_title;comment:巡检项（内容）名称;"`
	DeptId    *int   `json:"deptId" form:"deptId" gorm:"column:dept_id;comment:部门ID;"`
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName PatrolItem 表名
func (PatrolItem) TableName() string {
	return "patrol_item"
}
