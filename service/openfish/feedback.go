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
	err = global.DB.Create(feedback).Error
	return err
}

// DeleteFeedback 删除Feedback记录
// Author [piexlmax](https://github.com/piexlmax)
func (feedbackService *FeedbackService) DeleteFeedback(feedback openfish.Feedback) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
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
	err = global.DB.Transaction(func(tx *gorm.DB) error {
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
	err = global.DB.Save(&feedback).Error
	return err
}

// GetFeedback 根据id获取Feedback记录
// Author [piexlmax](https://github.com/piexlmax)
func (feedbackService *FeedbackService) GetFeedback(id uint) (feedback openfish.Feedback, err error) {
	err = global.DB.Where("id = ?", id).First(&feedback).Error
	return
}

// GetFeedbackInfoList 分页获取Feedback记录
// Author [piexlmax](https://github.com/piexlmax)
func (feedbackService *FeedbackService) GetFeedbackInfoList(info openfishReq.FeedbackSearch) (list []openfish.FeedbackVo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&openfish.Feedback{})
	var feedbacks []openfish.FeedbackVo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.CreatedBy != 0 {
		db = db.Where("feedback.created_by = ?", info.CreatedBy)
	}
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("feedback.created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.FeedbackText != "" {
		db = db.Where("feedback.feedback_text LIKE ?", "%"+info.FeedbackText+"%")
	}
	db = db.Where("feedback.parent_id = ?", 0)
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Joins("left join feedback fb on feedback.id=fb.parent_id").Select("feedback.*, fb.feedback_text reply_text").Limit(limit).Offset(offset).Order("created_at desc").Find(&feedbacks).Error
	return feedbacks, total, err
}
