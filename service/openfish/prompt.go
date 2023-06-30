package openfish

import (
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/openfish"
	openfishReq "github.com/oldweipro/gin-admin/model/openfish/request"
)

type PromptService struct {
}

// CreatePrompt 创建Prompt记录
func (promptService *PromptService) CreatePrompt(prompt *openfish.Prompt) (err error) {
	err = global.DB.Create(prompt).Error
	return err
}

// DeletePrompt 删除Prompt记录
func (promptService *PromptService) DeletePrompt(prompt openfish.Prompt) (err error) {
	err = global.DB.Delete(&prompt).Error
	return err
}

// DeletePromptByIds 批量删除Prompt记录
func (promptService *PromptService) DeletePromptByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]openfish.Prompt{}, "id in ?", ids.Ids).Error
	return err
}

// UpdatePrompt 更新Prompt记录
func (promptService *PromptService) UpdatePrompt(prompt openfish.Prompt) (err error) {
	err = global.DB.Save(&prompt).Error
	return err
}

// GetPrompt 根据id获取Prompt记录
func (promptService *PromptService) GetPrompt(id uint) (prompt openfish.Prompt, err error) {
	err = global.DB.Where("id = ?", id).First(&prompt).Error
	return
}

// GetPromptInfoList 分页获取Prompt记录
func (promptService *PromptService) GetPromptInfoList(info openfishReq.PromptSearch) (list []openfish.Prompt, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&openfish.Prompt{})
	var prompts []openfish.Prompt
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Description != "" {
		db = db.Where("description LIKE ?", "%"+info.Description+"%")
	}
	if info.Content != "" {
		db = db.Where("content LIKE ?", "%"+info.Content+"%")
	}
	if info.ShortcutKey != "" {
		db = db.Where("shortcut_key LIKE ?", "%"+info.ShortcutKey+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&prompts).Error
	return prompts, total, err
}
