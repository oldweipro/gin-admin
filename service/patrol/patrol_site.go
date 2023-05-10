package patrol

import (
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/patrol"
	patrolReq "github.com/oldweipro/gin-admin/model/patrol/request"
	"gorm.io/gorm"
)

type PatrolSiteService struct {
}

// CreatePatrolSite 创建PatrolSite记录
// Author [piexlmax](https://github.com/piexlmax)
func (patrolSiteService *PatrolSiteService) CreatePatrolSite(patrolSite patrol.PatrolSite) (err error) {
	err = global.GVA_DB.Create(&patrolSite).Error
	return err
}

// DeletePatrolSite 删除PatrolSite记录
// Author [piexlmax](https://github.com/piexlmax)
func (patrolSiteService *PatrolSiteService) DeletePatrolSite(patrolSite patrol.PatrolSite) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&patrol.PatrolSite{}).Where("id = ?", patrolSite.ID).Update("deleted_by", patrolSite.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&patrolSite).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeletePatrolSiteByIds 批量删除PatrolSite记录
// Author [piexlmax](https://github.com/piexlmax)
func (patrolSiteService *PatrolSiteService) DeletePatrolSiteByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&patrol.PatrolSite{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&patrol.PatrolSite{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdatePatrolSite 更新PatrolSite记录
// Author [piexlmax](https://github.com/piexlmax)
func (patrolSiteService *PatrolSiteService) UpdatePatrolSite(patrolSite patrol.PatrolSite) (err error) {
	err = global.GVA_DB.Save(&patrolSite).Error
	return err
}

// GetPatrolSite 根据id获取PatrolSite记录
// Author [piexlmax](https://github.com/piexlmax)
func (patrolSiteService *PatrolSiteService) GetPatrolSite(id uint) (patrolSite patrol.PatrolSite, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&patrolSite).Error
	return
}

// GetPatrolSiteInfoList 分页获取PatrolSite记录
// Author [piexlmax](https://github.com/piexlmax)
func (patrolSiteService *PatrolSiteService) GetPatrolSiteInfoList(info patrolReq.PatrolSiteSearch) (list []patrol.PatrolSite, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&patrol.PatrolSite{})
	var patrolSites []patrol.PatrolSite
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.SiteName != "" {
		db = db.Where("site_name LIKE ?", "%"+info.SiteName+"%")
	}
	if info.SitePositioning != "" {
		db = db.Where("site_positioning LIKE ?", "%"+info.SitePositioning+"%")
	}
	if info.DeptId != nil {
		db = db.Where("dept_id = ?", info.DeptId)
	}
	if info.Scope != nil {
		db = db.Where("scope = ?", info.Scope)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&patrolSites).Error
	return patrolSites, total, err
}
