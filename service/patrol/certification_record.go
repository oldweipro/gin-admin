package patrol

import (
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/patrol"
	patrolReq "github.com/oldweipro/gin-admin/model/patrol/request"
	"gorm.io/gorm"
)

type CertificationRecordService struct {
}

// CreateCertificationRecord 创建CertificationRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (certificationRecordService *CertificationRecordService) CreateCertificationRecord(certificationRecord patrol.CertificationRecord) (err error) {
	err = global.DB.Create(&certificationRecord).Error
	return err
}

// DeleteCertificationRecord 删除CertificationRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (certificationRecordService *CertificationRecordService) DeleteCertificationRecord(certificationRecord patrol.CertificationRecord) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&patrol.CertificationRecord{}).Where("id = ?", certificationRecord.ID).Update("deleted_by", certificationRecord.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&certificationRecord).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteCertificationRecordByIds 批量删除CertificationRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (certificationRecordService *CertificationRecordService) DeleteCertificationRecordByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&patrol.CertificationRecord{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&patrol.CertificationRecord{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateCertificationRecord 更新CertificationRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (certificationRecordService *CertificationRecordService) UpdateCertificationRecord(certificationRecord patrol.CertificationRecord) (err error) {
	err = global.DB.Save(&certificationRecord).Error
	return err
}

// GetCertificationRecord 根据id获取CertificationRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (certificationRecordService *CertificationRecordService) GetCertificationRecord(id uint) (certificationRecord patrol.CertificationRecord, err error) {
	err = global.DB.Where("id = ?", id).First(&certificationRecord).Error
	return
}

// GetCertificationRecordInfoList 分页获取CertificationRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (certificationRecordService *CertificationRecordService) GetCertificationRecordInfoList(info patrolReq.CertificationRecordSearch) (list []patrol.CertificationRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&patrol.CertificationRecord{})
	var certificationRecords []patrol.CertificationRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.CertificationIdCard != "" {
		db = db.Where("certification_id_card LIKE ?", "%"+info.CertificationIdCard+"%")
	}
	if info.CertificationRealName != "" {
		db = db.Where("certification_real_name LIKE ?", "%"+info.CertificationRealName+"%")
	}
	if info.CertificationResult != "" {
		db = db.Where("certification_result LIKE ?", "%"+info.CertificationResult+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&certificationRecords).Error
	return certificationRecords, total, err
}
