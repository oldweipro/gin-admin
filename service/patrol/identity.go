package patrol

import (
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/patrol"
	patrolReq "github.com/oldweipro/gin-admin/model/patrol/request"
	"gorm.io/gorm"
)

type IdentityService struct {
}

// CreateIdentity 创建Identity记录
// Author [piexlmax](https://github.com/piexlmax)
func (identityService *IdentityService) CreateIdentity(identity patrol.Identity) (err error) {
	err = global.GVA_DB.Create(&identity).Error
	return err
}

// DeleteIdentity 删除Identity记录
// Author [piexlmax](https://github.com/piexlmax)
func (identityService *IdentityService) DeleteIdentity(identity patrol.Identity) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&patrol.Identity{}).Where("id = ?", identity.ID).Update("deleted_by", identity.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&identity).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteIdentityByIds 批量删除Identity记录
// Author [piexlmax](https://github.com/piexlmax)
func (identityService *IdentityService) DeleteIdentityByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&patrol.Identity{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&patrol.Identity{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateIdentity 更新Identity记录
// Author [piexlmax](https://github.com/piexlmax)
func (identityService *IdentityService) UpdateIdentity(identity patrol.Identity) (err error) {
	err = global.GVA_DB.Save(&identity).Error
	return err
}

// GetIdentity 根据id获取Identity记录
// Author [piexlmax](https://github.com/piexlmax)
func (identityService *IdentityService) GetIdentity(id uint) (identity patrol.Identity, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&identity).Error
	return
}

// GetIdentityInfoList 分页获取Identity记录
// Author [piexlmax](https://github.com/piexlmax)
func (identityService *IdentityService) GetIdentityInfoList(info patrolReq.IdentitySearch) (list []patrol.Identity, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&patrol.Identity{})
	var identitys []patrol.Identity
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.IdCard != "" {
		db = db.Where("id_card = ?", info.IdCard)
	}
	if info.RealName != "" {
		db = db.Where("real_name = ?", info.RealName)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&identitys).Error
	return identitys, total, err
}
