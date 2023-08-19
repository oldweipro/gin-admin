package transaction

import (
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/transaction"
	"gorm.io/gorm"
	"time"
)

type SubscribeService struct {
}

// RenewalSubscription 续费订阅
func (subscribeService *SubscribeService) RenewalSubscription(userPlan *transaction.SubscriptionUser, plan *transaction.SubscriptionPlan) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// 续费时长，扣除鱼币，修改相关业务的时间放给定时任务吧。
		var subscriptionUserRecord transaction.SubscriptionUserRecord
		subscriptionUserRecord.SubscriptionUserId = &userPlan.ID
		subscriptionUserRecord.SubscriptionPlanId = &plan.ID
		subscriptionUserRecord.UserId = userPlan.UserId
		subscriptionUserRecord.BeforeEndTime = userPlan.EndTime
		duration := plan.Duration
		// 判断userPlan.EndTime时间是否小于当前时间
		var futureTime time.Time
		if userPlan.EndTime.Before(time.Now()) {
			futureTime = time.Now().Add(time.Duration(*duration) * time.Second)
		} else {
			futureTime = userPlan.EndTime.Add(time.Duration(*duration) * time.Second)
		}
		subscriptionUserRecord.AfterEndTime = futureTime
		subscriptionUserRecord.Name = plan.Name
		subscriptionUserRecord.Description = plan.Description
		subscriptionUserRecord.Price = plan.Price
		subscriptionUserRecord.Duration = plan.Duration
		subscriptionUserRecord.Quantity = plan.Quantity
		subscriptionUserRecord.MenuId = plan.MenuId
		if err := tx.Create(&subscriptionUserRecord).Error; err != nil {
			return err
		}
		// 更新用户计划
		userPlan.EndTime = futureTime
		if err := tx.Save(&userPlan).Error; err != nil {
			return err
		}
		// 扣除鱼币
		var wallets transaction.Wallets
		err = tx.Where("user_id = ?", userPlan.UserId).First(&wallets).Error
		// 更新用户钱包的鱼币
		balance := *wallets.Balance - *plan.Price
		if err = tx.Model(&transaction.Wallets{}).Where("id = ?", wallets.ID).Update("balance", balance).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
