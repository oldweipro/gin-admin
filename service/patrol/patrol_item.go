package patrol

import (
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/patrol"
	patrolReq "github.com/oldweipro/gin-admin/model/patrol/request"
	"gorm.io/gorm"
)

type PatrolItemService struct {
}

// CreatePatrolItem 创建PatrolItem记录
// Author [piexlmax](https://github.com/piexlmax)
func (patrolItemService *PatrolItemService) CreatePatrolItem(patrolItem patrol.PatrolItem) (err error) {
	err = global.DB.Create(&patrolItem).Error
	return err
}

// DeletePatrolItem 删除PatrolItem记录
// Author [piexlmax](https://github.com/piexlmax)
func (patrolItemService *PatrolItemService) DeletePatrolItem(patrolItem patrol.PatrolItem) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&patrol.PatrolItem{}).Where("id = ?", patrolItem.ID).Update("deleted_by", patrolItem.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&patrolItem).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeletePatrolItemByIds 批量删除PatrolItem记录
// Author [piexlmax](https://github.com/piexlmax)
func (patrolItemService *PatrolItemService) DeletePatrolItemByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&patrol.PatrolItem{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&patrol.PatrolItem{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdatePatrolItem 更新PatrolItem记录
// Author [piexlmax](https://github.com/piexlmax)
func (patrolItemService *PatrolItemService) UpdatePatrolItem(patrolItem patrol.PatrolItem) (err error) {
	err = global.DB.Save(&patrolItem).Error
	return err
}

// GetPatrolItem 根据id获取PatrolItem记录
// Author [piexlmax](https://github.com/piexlmax)
func (patrolItemService *PatrolItemService) GetPatrolItem(id uint) (patrolItem patrol.PatrolItem, err error) {
	err = global.DB.Where("id = ?", id).First(&patrolItem).Error
	return
}

// GetPatrolItemInfoList 分页获取PatrolItem记录
// Author [piexlmax](https://github.com/piexlmax)
func (patrolItemService *PatrolItemService) GetPatrolItemInfoList(info patrolReq.PatrolItemSearch) (list []patrol.PatrolItem, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&patrol.PatrolItem{})
	var patrolItems []patrol.PatrolItem
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ItemTitle != "" {
		db = db.Where("item_title LIKE ?", "%"+info.ItemTitle+"%")
	}
	if info.DeptId != nil {
		db = db.Where("dept_id = ?", info.DeptId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&patrolItems).Error
	return patrolItems, total, err
}
