package transaction

import (
	"fmt"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/transaction"
	openfishReq "github.com/oldweipro/gin-admin/model/transaction/request"
	"gorm.io/gorm"
	"strconv"
)

type ChatTicketService struct {
}

// CreateChatTicket 创建ChatTicket记录
// Author [piexlmax](https://github.com/piexlmax)
func (chatTicketService *ChatTicketService) CreateChatTicket(chatTicket *transaction.ChatTicket) (err error) {
	err = global.GVA_DB.Create(chatTicket).Error
	return err
}

// DeleteChatTicket 删除ChatTicket记录
// Author [piexlmax](https://github.com/piexlmax)
func (chatTicketService *ChatTicketService) DeleteChatTicket(chatTicket transaction.ChatTicket) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&transaction.ChatTicket{}).Where("id = ?", chatTicket.ID).Update("deleted_by", chatTicket.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&chatTicket).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteChatTicketByIds 批量删除ChatTicket记录
// Author [piexlmax](https://github.com/piexlmax)
func (chatTicketService *ChatTicketService) DeleteChatTicketByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&transaction.ChatTicket{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&transaction.ChatTicket{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateChatTicket 更新ChatTicket记录
// Author [piexlmax](https://github.com/piexlmax)
func (chatTicketService *ChatTicketService) UpdateChatTicket(chatTicket transaction.ChatTicket) (err error) {
	err = global.GVA_DB.Save(&chatTicket).Error
	return err
}

// GetChatTicket 根据id获取ChatTicket记录
// Author [piexlmax](https://github.com/piexlmax)
func (chatTicketService *ChatTicketService) GetChatTicket(id uint) (chatTicket transaction.ChatTicket, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&chatTicket).Error
	return
}

// GetChatTicketInfoList 分页获取ChatTicket记录
// Author [piexlmax](https://github.com/piexlmax)
func (chatTicketService *ChatTicketService) GetChatTicketInfoList(info openfishReq.ChatTicketSearch) (list []transaction.ChatTicket, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&transaction.ChatTicket{})
	var chatTickets []transaction.ChatTicket
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Amount != nil {
		db = db.Where("amount = ?", info.Amount)
	}
	if info.TicketName != "" {
		db = db.Where("ticket_name LIKE ?", "%"+info.TicketName+"%")
	}
	if info.BelongTo != nil {
		db = db.Where("belong_to = ?", info.BelongTo)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&chatTickets).Error
	return chatTickets, total, err
}

// HandleValidateChatTicket 验证鱼币兑换码
func (chatTicketService *ChatTicketService) HandleValidateChatTicket(ticketValue string, wallets *transaction.Wallets) (err error) {
	var chatTicket transaction.ChatTicket
	err = global.GVA_DB.Where("ticket_value = ? and (expiration_time = 0 or expiration_time > UNIX_TIMESTAMP())", ticketValue).First(&chatTicket).Error
	if err != nil {
		return err
	}
	// 兑换鱼币
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var srcWalletId uint = 0
		// 更新交易记录
		remark := fmt.Sprintf("验证鱼币兑换码: %s;兑换数量: %s", ticketValue, strconv.Itoa(*chatTicket.Amount))
		transactionHistory := transaction.TransactionHistory{
			UserId:       &wallets.UserId,
			SrcWalletId:  &srcWalletId,
			DestWalletId: &wallets.ID,
			TypeEnum:     "deposit",
			Quantity:     chatTicket.Amount,
			ProductId:    &chatTicket.ID,
			Remark:       remark,
			CreatedBy:    wallets.UserId,
		}
		if err = tx.Create(&transactionHistory).Error; err != nil {
			return err
		}
		balance := *wallets.Balance + *chatTicket.Amount
		// 更新用户钱包的鱼币
		if err = tx.Model(&transaction.Wallets{}).Where("id = ?", wallets.ID).Update("balance", balance).Error; err != nil {
			return err
		}
		// 删除鱼币兑换码
		if err = tx.Model(&transaction.ChatTicket{}).Where("id = ?", chatTicket.ID).Update("belong_to", wallets.UserId).Update("deleted_by", chatTicket.CreatedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&chatTicket).Error; err != nil {
			return err
		}
		// nil提交事务
		return nil
	})
	return
}
