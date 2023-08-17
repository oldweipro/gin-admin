package transaction

import (
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/transaction"
)

type SubscriptionPlanService struct {
}

// CreateSubscriptionPlan 创建SubscriptionPlan记录
func (subscriptionPlanService *SubscriptionPlanService) CreateSubscriptionPlan(subscriptionPlan *transaction.SubscriptionPlan) (err error) {
	err = global.DB.Create(subscriptionPlan).Error
	return err
}

// DeleteSubscriptionPlan 删除SubscriptionPlan记录
func (subscriptionPlanService *SubscriptionPlanService) DeleteSubscriptionPlan(subscriptionPlan transaction.SubscriptionPlan) (err error) {
	err = global.DB.Delete(&subscriptionPlan).Error
	return err
}

// DeleteSubscriptionPlanByIds 批量删除SubscriptionPlan记录
func (subscriptionPlanService *SubscriptionPlanService) DeleteSubscriptionPlanByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]transaction.SubscriptionPlan{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSubscriptionPlan 更新SubscriptionPlan记录
func (subscriptionPlanService *SubscriptionPlanService) UpdateSubscriptionPlan(subscriptionPlan transaction.SubscriptionPlan) (err error) {
	err = global.DB.Save(&subscriptionPlan).Error
	return err
}

// GetSubscriptionPlan 根据id获取SubscriptionPlan记录
func (subscriptionPlanService *SubscriptionPlanService) GetSubscriptionPlan(id uint) (subscriptionPlan transaction.SubscriptionPlan, err error) {
	err = global.DB.Where("id = ?", id).First(&subscriptionPlan).Error
	return
}

// GetSubscriptionPlanByTag 根据tag获取SubscriptionPlan记录
func (subscriptionPlanService *SubscriptionPlanService) GetSubscriptionPlanByTag(tag uint) (subscriptionPlan transaction.SubscriptionPlan, err error) {
	err = global.DB.Where("tag = ?", tag).First(&subscriptionPlan).Error
	return
}

// GetCurrentSubscriptionPlan 查询当前用户订阅计划
func (subscriptionPlanService *SubscriptionPlanService) GetCurrentSubscriptionPlan(id uint) (subscriptionUser transaction.SubscriptionUser, err error) {
	err = global.DB.Where("user_id = ?", id).First(&subscriptionUser).Error
	return
}

// GetSubscriptionPlanInfoList 分页获取SubscriptionPlan记录
//func (subscriptionPlanService *SubscriptionPlanService) GetSubscriptionPlanInfoList(info openfishReq.SubscriptionPlanSearch) (list []transaction.SubscriptionPlan, total int64, err error) {
//	limit := info.PageSize
//	offset := info.PageSize * (info.Page - 1)
//	// 创建db
//	db := global.DB.Model(&transaction.SubscriptionPlan{})
//	var subscriptionPlans []transaction.SubscriptionPlan
//	// 如果有条件搜索 下方会自动创建搜索语句
//	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
//		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
//	}
//	if info.Name != "" {
//		db = db.Where("name LIKE ?", "%"+info.Name+"%")
//	}
//	if info.Description != "" {
//		db = db.Where("description LIKE ?", "%"+info.Description+"%")
//	}
//	err = db.Count(&total).Error
//	if err != nil {
//		return
//	}
//
//	err = db.Limit(limit).Offset(offset).Find(&subscriptionPlans).Error
//	return subscriptionPlans, total, err
//}
