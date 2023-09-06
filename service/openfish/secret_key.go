package openfish

import (
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/openfish"
	openfishReq "github.com/oldweipro/gin-admin/model/openfish/request"
	"gorm.io/gorm"
	"strings"
)

type SecretKeyService struct {
}

// CreateSecretKey 创建SecretKey记录
func (secretKeyService *SecretKeyService) CreateSecretKey(secretKey *openfish.SecretKey) (err error) {
	err = global.DB.Create(secretKey).Error
	return err
}

// DeleteSecretKey 删除SecretKey记录
func (secretKeyService *SecretKeyService) DeleteSecretKey(secretKey openfish.SecretKey) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&openfish.SecretKey{}).Where("id = ?", secretKey.ID).Update("deleted_by", secretKey.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&secretKey).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteSecretKeyByIds 批量删除SecretKey记录
func (secretKeyService *SecretKeyService) DeleteSecretKeyByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&openfish.SecretKey{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&openfish.SecretKey{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateSecretKey 更新SecretKey记录
func (secretKeyService *SecretKeyService) UpdateSecretKey(secretKey openfish.SecretKey) (err error) {
	err = global.DB.Save(&secretKey).Error
	return err
}

// GetSecretKey 根据id获取SecretKey记录
func (secretKeyService *SecretKeyService) GetSecretKey(id, userId uint) (secretKey openfish.SecretKey, err error) {
	err = global.DB.Where("id = ? and created_by = ?", id, userId).First(&secretKey).Error
	return
}

// GetSecretKeyBySk 根据sk获取SecretKey记录
func (secretKeyService *SecretKeyService) GetSecretKeyBySk(sk string) (secretKey openfish.SecretKey, err error) {
	err = global.DB.Where("sk = ? and status = 1", sk).First(&secretKey).Error
	return
}

// GetSecretKeyInfoList 分页获取SecretKey记录
func (secretKeyService *SecretKeyService) GetSecretKeyInfoList(info openfishReq.SecretKeySearch) (list []openfish.SecretKey, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&openfish.SecretKey{})
	var secretKeys []openfish.SecretKey
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Sk != "" {
		db = db.Where("sk = ?", info.Sk)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&secretKeys).Error
	return secretKeys, total, err
}

// GetSecretKeyInfoLessList 分页获取SecretKey记录
func (secretKeyService *SecretKeyService) GetSecretKeyInfoLessList(info openfishReq.SecretKeySearch) (list []openfish.SecretKey, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&openfish.SecretKey{})
	var secretKeys []openfish.SecretKey
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Sk != "" {
		db = db.Where("sk = ?", info.Sk)
	}
	db = db.Where("created_by = ?", info.CreatedBy)
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&secretKeys).Error
	var maskedKeys []openfish.SecretKey
	for _, key := range secretKeys {
		// sk密钥脱敏
		key.Sk = secretKeyService.MaskString(key.Sk, 6, 36)
		maskedKeys = append(maskedKeys, key)
	}
	return maskedKeys, total, err
}

func (secretKeyService *SecretKeyService) MaskString(input string, start int, length int) string {
	if start < 0 || start >= len(input) || length <= 0 {
		return input
	}

	end := start + length
	if end > len(input) {
		end = len(input)
	}

	prefix := input[:start]
	suffix := input[end:]

	maskedPart := strings.Repeat("*", length)
	return prefix + maskedPart + suffix
}

func (secretKeyService *SecretKeyService) GetSecretKeyList(id uint) (secretKey int64, err error) {
	err = global.DB.Model(&openfish.SecretKey{}).Where("created_by = ?", id).Count(&secretKey).Error
	return
}
