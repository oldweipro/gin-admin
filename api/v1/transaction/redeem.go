package transaction

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/response"
	"github.com/oldweipro/gin-admin/model/transaction/request"
	"github.com/oldweipro/gin-admin/service"
	"go.uber.org/zap"
)

type RedeemApi struct {
}

var redeemService = service.ServiceGroupApp.TransactionServiceGroup.RedeemService

// GenerateRedeemCode 生成兑换码
func (redeemApi *RedeemApi) GenerateRedeemCode(c *gin.Context) {
	var redeemRequest request.RedeemRequest
	err := c.ShouldBindJSON(&redeemRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redeemRequest.Pieces <= 0 {
		response.OkWithMessage("生成兑换码成功，啥也没生成，因为你填写的个数为0", c)
		return
	}
	if err := redeemService.GenerateRedeemCode(redeemRequest); err != nil {
		global.Logger.Error("生成失败!", zap.Error(err))
		response.FailWithMessage("生成失败", c)
	} else {
		response.OkWithMessage("生成兑换码成功", c)
	}
}
