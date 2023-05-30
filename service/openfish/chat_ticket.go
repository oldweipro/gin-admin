package openfish

import (
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/openfish"
	openfishReq "github.com/oldweipro/gin-admin/model/openfish/request"
	"gorm.io/gorm"
)

type ChatTicketService struct {
}

// CreateChatTicket 创建ChatTicket记录
// Author [piexlmax](https://github.com/piexlmax)
func (chatTicketService *ChatTicketService) CreateChatTicket(chatTicket *openfish.ChatTicket) (err error) {
	err = global.GVA_DB.Create(chatTicket).Error
	return err
}

// DeleteChatTicket 删除ChatTicket记录
// Author [piexlmax](https://github.com/piexlmax)
func (chatTicketService *ChatTicketService) DeleteChatTicket(chatTicket openfish.ChatTicket) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&openfish.ChatTicket{}).Where("id = ?", chatTicket.ID).Update("deleted_by", chatTicket.DeletedBy).Error; err != nil {
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
		if err := tx.Model(&openfish.ChatTicket{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&openfish.ChatTicket{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateChatTicket 更新ChatTicket记录
// Author [piexlmax](https://github.com/piexlmax)
func (chatTicketService *ChatTicketService) UpdateChatTicket(chatTicket openfish.ChatTicket) (err error) {
	err = global.GVA_DB.Save(&chatTicket).Error
	return err
}

// GetChatTicket 根据id获取ChatTicket记录
// Author [piexlmax](https://github.com/piexlmax)
func (chatTicketService *ChatTicketService) GetChatTicket(id uint) (chatTicket openfish.ChatTicket, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&chatTicket).Error
	return
}

// GetChatTicketInfoList 分页获取ChatTicket记录
// Author [piexlmax](https://github.com/piexlmax)
func (chatTicketService *ChatTicketService) GetChatTicketInfoList(info openfishReq.ChatTicketSearch) (list []openfish.ChatTicket, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&openfish.ChatTicket{})
	var chatTickets []openfish.ChatTicket
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
