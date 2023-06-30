// 自动生成模板PatrolSite
package patrol

import (
	"github.com/oldweipro/gin-admin/global"
)

// PatrolSite 结构体
type PatrolSite struct {
	global.Model
	SiteName        string `json:"siteName" form:"siteName" gorm:"column:site_name;comment:点位名称;"`
	SitePositioning string `json:"sitePositioning" form:"sitePositioning" gorm:"column:site_positioning;comment:地点定位;"`
	DeptId          *int   `json:"deptId" form:"deptId" gorm:"column:dept_id;comment:所属部门;"`
	Scope           *int   `json:"scope" form:"scope" gorm:"column:scope;comment:打卡范围: 单位米;"`
	CreatedBy       uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy       uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy       uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName PatrolSite 表名
func (PatrolSite) TableName() string {
	return "patrol_site"
}
