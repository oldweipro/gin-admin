package openfish

import (
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/openfish"
	openfishReq "github.com/oldweipro/gin-admin/model/openfish/request"
	"gorm.io/gorm"
)

type FeedbackService struct {
}

// CreateFeedback 创建Feedback记录
// Author [piexlmax](https://github.com/piexlmax)
func (feedbackService *FeedbackService) CreateFeedback(feedback *openfish.Feedback) (err error) {
	err = global.GVA_DB.Create(feedback).Error
	return err
}

// DeleteFeedback 删除Feedback记录
// Author [piexlmax](https://github.com/piexlmax)
func (feedbackService *FeedbackService) DeleteFeedback(feedback openfish.Feedback) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&openfish.Feedback{}).Where("id = ?", feedback.ID).Update("deleted_by", feedback.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&feedback).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteFeedbackByIds 批量删除Feedback记录
// Author [piexlmax](https://github.com/piexlmax)
func (feedbackService *FeedbackService) DeleteFeedbackByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&openfish.Feedback{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&openfish.Feedback{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateFeedback 更新Feedback记录
// Author [piexlmax](https://github.com/piexlmax)
func (feedbackService *FeedbackService) UpdateFeedback(feedback openfish.Feedback) (err error) {
	err = global.GVA_DB.Save(&feedback).Error
	return err
}

// GetFeedback 根据id获取Feedback记录
// Author [piexlmax](https://github.com/piexlmax)
func (feedbackService *FeedbackService) GetFeedback(id uint) (feedback openfish.Feedback, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&feedback).Error
	return
}

// GetFeedbackInfoList 分页获取Feedback记录
// Author [piexlmax](https://github.com/piexlmax)
func (feedbackService *FeedbackService) GetFeedbackInfoList(info openfishReq.FeedbackSearch) (list []openfish.Feedback, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&openfish.Feedback{})
	var feedbacks []openfish.Feedback
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Feedback_text != "" {
		db = db.Where("feedback_text LIKE ?", "%"+info.Feedback_text+"%")
	}
	if info.User_id != nil {
		db = db.Where("user_id = ?", info.User_id)
	}
	if info.Parent_id != nil {
		db = db.Where("parent_id = ?", info.Parent_id)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&feedbacks).Error
	return feedbacks, total, err
}
