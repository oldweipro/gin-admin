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
		redeems = append(redeems, redeem)
	}
	err = global.DB.CreateInBatches(redeems, len(redeems)).Error
	return err
}
