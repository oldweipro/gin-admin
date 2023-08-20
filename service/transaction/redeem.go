package transaction

import (
	"github.com/google/uuid"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/transaction"
	"github.com/oldweipro/gin-admin/model/transaction/request"
	"strings"
	"time"
)

type RedeemService struct {
}

// GenerateRedeemCode 生成兑换码
func (redeemService *RedeemService) GenerateRedeemCode(req request.RedeemRequest) (err error) {
	var redeems []transaction.RedeemCode
	for i := 0; i < req.Pieces; i++ {
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
		redeems = append(redeems, redeem)
	}
	err = global.DB.CreateInBatches(redeems, len(redeems)).Error
	return err
}

// GetRedeemCodeList 获取兑换码列表
func (redeemService *RedeemService) GetRedeemCodeList(info request.RedeemSearch) (list []transaction.RedeemCode, total int64, err error) {
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
