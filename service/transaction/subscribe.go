package transaction

import (
	"errors"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/transaction"
	"gorm.io/gorm"
	"time"
)

type SubscribeService struct {
}

// RenewalSubscription 续费订阅
func (subscribeService *SubscribeService) RenewalSubscription(userPlan *transaction.SubscriptionUser, plan *transaction.SubscriptionPlan) (err error) {
	// 查询钱包余额
	var wallets transaction.Wallets
	err = global.DB.Where("user_id = ?", userPlan.UserId).First(&wallets).Error
	// 检查余额是否充足
	if *wallets.Balance < *plan.Price {
		return errors.New("余额不足")
	}
	// 扣除鱼币
	balance := *wallets.Balance - *plan.Price
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
		var status uint = 1
		userPlan.Status = &status
		if err := tx.Save(&userPlan).Error; err != nil {
			return err
		}
		// 更新消费记录到 transaction
		var srcWalletId uint = 0 // 系统账户
		transactionHistory := transaction.TransactionHistory{
			UserId:        &wallets.UserId,
			SrcWalletId:   &wallets.ID,
			DestWalletId:  &srcWalletId,
			TypeEnum:      "checkin",
			Quantity:      subscriptionUserRecord.Quantity,
			Amount:        subscriptionUserRecord.Price,
			BeforeBalance: wallets.Balance,
			AfterBalance:  &balance,
			ProductId:     &srcWalletId,
			Remark:        subscriptionUserRecord.Description,
			CreatedBy:     wallets.UserId,
		}
		if err = tx.Create(&transactionHistory).Error; err != nil {
			return err
		}
		// 更新用户钱包的鱼币
		if err = tx.Model(&transaction.Wallets{}).Where("id = ?", wallets.ID).Update("balance", balance).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
