package transaction

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/transaction"
	"github.com/oldweipro/gin-admin/model/transaction/request"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type RedeemService struct {
}

// GenerateRedeemCode 生成兑换码
func (redeemService *RedeemService) GenerateRedeemCode(req request.RedeemCodeRequest) (err error) {
	var redeems []transaction.RedeemCode
	for i := 0; i < *req.Pieces; i++ {
		var zero uint = 0
		var redeem transaction.RedeemCode
		// 已被兑换次数
		redeem.TotalCount = req.TotalCount
		redeem.Amount = req.Amount
		redeem.PerLimit = req.PerLimit
		redeem.LeftCount = req.TotalCount
		redeem.ExpireTime = time.UnixMilli(req.ExpireTime)
		newString := "fish" + strings.ReplaceAll(uuid.NewString(), "-", "") + strings.ReplaceAll(uuid.NewString(), "-", "")[:16]
		redeem.Code = newString
		redeem.Status = req.Status
		redeem.TotalRedeemed = &zero
		redeems = append(redeems, redeem)
	}
	err = global.DB.CreateInBatches(redeems, len(redeems)).Error
	return err
}

// CheckRedeemCode 检查兑换码是否有效
func (redeemService *RedeemService) CheckRedeemCode(code string, userId uint) (redeemCode transaction.RedeemCode, err error) {
	err = global.DB.Where("code=? and per_limit>0 and left_count>0 and status=0 and UNIX_TIMESTAMP(expire_time)>unix_timestamp(now())", code).First(&redeemCode).Error
	if err != nil {
		return transaction.RedeemCode{}, errors.New("兑换码失效")
	}
	if *redeemCode.TotalCount > 1 {
		// 这个兑换码可以被兑换多次，查询兑换码和个人兑换记录
		var count int64
		if err = global.DB.Where("redeem_code_id=? and user_id=?", redeemCode.ID, userId).Count(&count).Error; err != nil {
			return transaction.RedeemCode{}, err
		}
		if count >= int64(*redeemCode.PerLimit) {
			return transaction.RedeemCode{}, errors.New("超出兑换限制")
		}
	}
	return
}

// RedeemFishCoin 生成兑换码
func (redeemService *RedeemService) RedeemFishCoin(redeemCode *transaction.RedeemCode, wallets *transaction.Wallets) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		var srcWalletId uint = 0
		// 更新交易记录
		remark := fmt.Sprintf("验证鱼币兑换码: %s;兑换数量: %s", redeemCode.Code, strconv.Itoa(int(*redeemCode.Amount)))
		balance := *wallets.Balance + *redeemCode.Amount
		transactionHistory := transaction.TransactionHistory{
			UserId:        &wallets.UserId,
			SrcWalletId:   &srcWalletId,
			DestWalletId:  &wallets.ID,
			TypeEnum:      "deposit",
			Quantity:      redeemCode.Amount,
			BeforeBalance: wallets.Balance,
			AfterBalance:  &balance,
			ProductId:     &redeemCode.ID,
			Remark:        remark,
			CreatedBy:     wallets.UserId,
		}
		if err = tx.Create(&transactionHistory).Error; err != nil {
			return err
		}
		// 更新用户钱包的鱼币
		if err = tx.Model(&transaction.Wallets{}).
			Where("id = ?", wallets.ID).
			Update("balance", balance).
			Error; err != nil {
			return err
		}
		// 更新兑换码
		*redeemCode.TotalRedeemed = *redeemCode.TotalRedeemed + 1
		*redeemCode.LeftCount = *redeemCode.TotalCount - *redeemCode.TotalRedeemed
		status := 0
		if *redeemCode.LeftCount == 0 {
			status = 1
		}
		if err = tx.Model(&transaction.RedeemCode{}).
			Where("id = ?", redeemCode.ID).
			Update("status", status).
			Update("left_count", *redeemCode.LeftCount).
			Update("total_redeemed", *redeemCode.TotalRedeemed).
			Error; err != nil {
			return err
		}
		// 更新兑换码记录
		var redeemLog transaction.RedeemLog
		redeemLog.RedeemCodeId = &redeemCode.ID
		redeemLog.RedeemTime = time.Now()
		redeemLog.UserId = &wallets.UserId
		if err = tx.Create(&redeemLog).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// GetRedeemCodeList 获取兑换码列表
func (redeemService *RedeemService) GetRedeemCodeList(info request.RedeemCodeSearch) (list []transaction.RedeemCode, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&transaction.RedeemCode{})
	var redeemCodes []transaction.RedeemCode
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Amount != nil {
		db = db.Where("name = ?", info.Amount)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&redeemCodes).Error
	return redeemCodes, total, err
}
