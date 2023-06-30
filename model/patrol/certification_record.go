// 自动生成模板CertificationRecord
package patrol

import (
	"github.com/oldweipro/gin-admin/global"
)

// CertificationRecord 结构体
type CertificationRecord struct {
	global.Model
	CertificationIdCard   string `json:"certificationIdCard" form:"certification_id_card" gorm:"column:certification_id_card;comment:认证身份ID;"`
	CertificationRealName string `json:"certificationRealName" form:"certification_real_name" gorm:"column:certification_real_name;comment:认证真实姓名;"`
	CertificationResult   string `json:"certificationResult" form:"certification_result" gorm:"column:certification_result;comment:认证结果;"`
	CertificationCode     *int   `json:"certificationCode" form:"certification_code" gorm:"column:certification_code;default:0;comment:认证结果状态码;"`
	CertificationMsg      string `json:"certificationMsg" form:"certification_msg" gorm:"column:certification_msg;comment:;"`
	CreatedBy             uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy             uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy             uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName CertificationRecord 表名
func (CertificationRecord) TableName() string {
	return "certification_record"
}
