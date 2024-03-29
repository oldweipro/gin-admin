package transaction

import (
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/transaction"
	openfishReq "github.com/oldweipro/gin-admin/model/transaction/request"
	"gorm.io/gorm"
)

type HistoryService struct {
}

// CreateTransactionHistory 创建TransactionHistory记录
func (historyService *HistoryService) CreateTransactionHistory(transactionHistory *transaction.TransactionHistory) (err error) {
	err = global.DB.Create(transactionHistory).Error
	return err
}

// DeleteTransactionHistory 删除TransactionHistory记录
func (historyService *HistoryService) DeleteTransactionHistory(transactionHistory transaction.TransactionHistory) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&transaction.TransactionHistory{}).Where("id = ?", transactionHistory.ID).Update("deleted_by", transactionHistory.DeletedBy).Error; err != nil {
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
func (historyService *HistoryService) DeleteTransactionHistoryByIds(ids request.IdsReq, deletedBy uint) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&transaction.TransactionHistory{}).Where("id in ?", ids.Ids).Update("deleted_by", deletedBy).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&transaction.TransactionHistory{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateTransactionHistory 更新TransactionHistory记录
func (historyService *HistoryService) UpdateTransactionHistory(transactionHistory transaction.TransactionHistory) (err error) {
	err = global.DB.Save(&transactionHistory).Error
	return err
}

// GetTransactionHistory 根据id获取TransactionHistory记录
func (historyService *HistoryService) GetTransactionHistory(id uint) (transactionHistory transaction.TransactionHistory, err error) {
	err = global.DB.Where("id = ?", id).First(&transactionHistory).Error
	return
}

// GetTransactionHistoryInfoList 分页获取TransactionHistory记录
func (historyService *HistoryService) GetTransactionHistoryInfoList(info openfishReq.TransactionHistorySearch) (list []transaction.TransactionHistory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&transaction.TransactionHistory{})
	var transactionHistorys []transaction.TransactionHistory
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

func (historyService *HistoryService) GetTodayTransactionHistoryByCurrentUser(userId uint) (count int64, err error) {
	err = global.DB.Model(&transaction.TransactionHistory{}).Where("TO_DAYS(created_at)=TO_DAYS(NOW()) AND type_enum='checkin' AND user_id = ?", userId).Count(&count).Error
	return
}
