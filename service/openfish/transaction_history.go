package openfish

import (
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/openfish"
	openfishReq "github.com/oldweipro/gin-admin/model/openfish/request"
	"gorm.io/gorm"
)

type TransactionHistoryService struct {
}

// CreateTransactionHistory 创建TransactionHistory记录
// Author [piexlmax](https://github.com/piexlmax)
func (transactionHistoryService *TransactionHistoryService) CreateTransactionHistory(transactionHistory *openfish.TransactionHistory) (err error) {
	err = global.GVA_DB.Create(transactionHistory).Error
	return err
}

// DeleteTransactionHistory 删除TransactionHistory记录
// Author [piexlmax](https://github.com/piexlmax)
func (transactionHistoryService *TransactionHistoryService) DeleteTransactionHistory(transactionHistory openfish.TransactionHistory) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&openfish.TransactionHistory{}).Where("id = ?", transactionHistory.ID).Update("deleted_by", transactionHistory.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&transactionHistory).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteTransactionHistoryByIds 批量删除TransactionHistory记录
// Author [piexlmax](https://github.com/piexlmax)
func (transactionHistoryService *TransactionHistoryService) DeleteTransactionHistoryByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&openfish.TransactionHistory{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&openfish.TransactionHistory{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateTransactionHistory 更新TransactionHistory记录
// Author [piexlmax](https://github.com/piexlmax)
func (transactionHistoryService *TransactionHistoryService) UpdateTransactionHistory(transactionHistory openfish.TransactionHistory) (err error) {
	err = global.GVA_DB.Save(&transactionHistory).Error
	return err
}

// GetTransactionHistory 根据id获取TransactionHistory记录
// Author [piexlmax](https://github.com/piexlmax)
func (transactionHistoryService *TransactionHistoryService) GetTransactionHistory(id uint) (transactionHistory openfish.TransactionHistory, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&transactionHistory).Error
	return
}

// GetTransactionHistoryInfoList 分页获取TransactionHistory记录
// Author [piexlmax](https://github.com/piexlmax)
func (transactionHistoryService *TransactionHistoryService) GetTransactionHistoryInfoList(info openfishReq.TransactionHistorySearch) (list []openfish.TransactionHistory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&openfish.TransactionHistory{})
	var transactionHistorys []openfish.TransactionHistory
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.UserId != nil {
		db = db.Where("user_id = ?", info.UserId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&transactionHistorys).Error
	return transactionHistorys, total, err
}
