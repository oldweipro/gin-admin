package openfish

import (
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/openfish"
	openfishReq "github.com/oldweipro/gin-admin/model/openfish/request"
	"gorm.io/gorm"
	"time"
)

type ConversationService struct {
}

// CreateConversation 创建Conversation记录
// Author [piexlmax](https://github.com/piexlmax)
func (conversationService *ConversationService) CreateConversation(conversation *openfish.Conversation) (err error) {
	err = global.GVA_DB.Create(conversation).Error
	return err
}

// CreateConversationRecord 创建ConversationRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (conversationService *ConversationService) CreateConversationRecord(conversationRecord *openfish.ConversationRecord) (err error) {
	err = global.GVA_DB.Create(conversationRecord).Error
	return err
}

// DeleteConversation 删除Conversation记录
// Author [piexlmax](https://github.com/piexlmax)
func (conversationService *ConversationService) DeleteConversation(conversation openfish.Conversation) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&openfish.Conversation{}).Where("id = ?", conversation.ID).Update("deleted_by", conversation.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&conversation).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteConversationByIds 批量删除Conversation记录
// Author [piexlmax](https://github.com/piexlmax)
func (conversationService *ConversationService) DeleteConversationByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&openfish.Conversation{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&openfish.Conversation{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateConversation 更新Conversation记录
// Author [piexlmax](https://github.com/piexlmax)
func (conversationService *ConversationService) UpdateConversation(conversation openfish.Conversation) (err error) {
	err = global.GVA_DB.Save(&conversation).Error
	return err
}

// UpdateConversationTime 更新Conversation时间
// Author [piexlmax](https://github.com/piexlmax)
func (conversationService *ConversationService) UpdateConversationTime(id uint) (err error) {
	err = global.GVA_DB.Model(&openfish.Conversation{}).Where("id = ?", id).Update("updated_at", time.Now()).Error
	return err
}

// GetConversation 根据id获取Conversation记录
// Author [piexlmax](https://github.com/piexlmax)
func (conversationService *ConversationService) GetConversation(id uint) (conversation openfish.Conversation, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&conversation).Error
	return
}

// GetConversationInfoList 分页获取Conversation记录
// Author [piexlmax](https://github.com/piexlmax)
func (conversationService *ConversationService) GetConversationInfoList(info openfishReq.ConversationSearch) (list []openfish.Conversation, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&openfish.Conversation{})
	var conversations []openfish.Conversation
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.CreatedBy != 0 {
		db = db.Where("created_by = ?", info.CreatedBy)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&conversations).Error
	return conversations, total, err
}

// GetConversationRecordListWithTokenByConversationId 根据pid查询会话信息列表
// error: error, status code: 400, message: This model's maximum context length is 4097 tokens. However, your messages resulted in 6301 tokens. Please reduce the length of the messages.
func (conversationService *ConversationService) GetConversationRecordListWithTokenByConversationId(conversationId uint, tokenCount int) ([]openfish.ConversationRecord, error) {
	var conversationRecords []openfish.ConversationRecord
	query := `SELECT *
		FROM conversation_record
		WHERE id IN (
		  SELECT id
		  FROM (
			SELECT id, created_by,created_at,
				   @sum := @sum + CHAR_LENGTH(content) AS sum
			FROM conversation_record, 
				 (SELECT @sum := 0) AS vars 
				WHERE conversation_id = ?
			ORDER BY created_at desc
		  ) AS t
		  WHERE sum <= ?
		) ORDER BY created_at ASC`
	err := global.GVA_DB.Raw(query, conversationId, 2596-tokenCount).Scan(&conversationRecords).Error
	return conversationRecords, err
}

// GetConversationRecordListByUserId 根据用户ID查询会话列表及会话信息列表
func (conversationService *ConversationService) GetConversationRecordListByUserId(userId uint) ([]openfish.ConversationRecord, error) {
	var conversationRecords []openfish.ConversationRecord
	err := global.GVA_DB.Model(&openfish.ConversationRecord{}).Where("conversation_id IN (SELECT id FROM conversation WHERE created_by = ?)", userId).Order("created_at asc").Find(&conversationRecords).Error
	return conversationRecords, err
}

// GetConversationRecordListByConversationId 根据conversationId查询会话信息列表
func (conversationService *ConversationService) GetConversationRecordListByConversationId(conversationId uint) ([]openfish.ConversationRecord, error) {
	var conversationRecords []openfish.ConversationRecord
	err := global.GVA_DB.Model(&openfish.ConversationRecord{}).Where("conversation_id = ?", conversationId).Order("created_at asc").Find(&conversationRecords).Error
	return conversationRecords, err
}

// GetConversationListByUserId 根据用户ID查询会话列表
func (conversationService *ConversationService) GetConversationListByUserId(userId uint) ([]openfish.Conversation, error) {
	var conversations []openfish.Conversation
	err := global.GVA_DB.Model(&openfish.Conversation{}).Where("created_by = ?", userId).Order("updated_at desc").First(&conversations).Error
	return conversations, err
}
