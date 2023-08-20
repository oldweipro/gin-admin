package transaction

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/response"
	"github.com/oldweipro/gin-admin/model/transaction/request"
	"github.com/oldweipro/gin-admin/service"
	"github.com/oldweipro/gin-admin/utils"
	"go.uber.org/zap"
)

type RedeemApi struct {
}

var redeemService = service.ServiceGroupApp.TransactionServiceGroup.RedeemService

// GenerateRedeemCode 生成兑换码
func (redeemApi *RedeemApi) GenerateRedeemCode(c *gin.Context) {
	var redeemRequest request.RedeemCodeRequest
	err := c.ShouldBindJSON(&redeemRequest)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if *redeemRequest.Pieces <= 0 {
		response.OkWithMessage("啥也没生成，你填写的生成个数小于1", c)
		return
	}
	if *redeemRequest.TotalCount <= 0 {
		response.OkWithMessage("啥也没生成，你填写的兑换码数量小于1", c)
		return
	}
	if *redeemRequest.PerLimit <= 0 {
		response.OkWithMessage("啥也没生成，你填写的兑换码频次小于1", c)
		return
	}
	if err := redeemService.GenerateRedeemCode(redeemRequest); err != nil {
		global.Logger.Error("生成失败!", zap.Error(err))
		response.FailWithMessage("生成失败", c)
	} else {
		response.OkWithMessage("生成兑换码成功", c)
	}
}

// RedeemFishCoin 兑换鱼币
func (redeemApi *RedeemApi) RedeemFishCoin(c *gin.Context) {
	var redeemFishCoin request.RedeemFishCoinRequest
	err := c.ShouldBindJSON(&redeemFishCoin)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redeemFishCoin.RedeemCode == "" {
		response.FailWithMessage("请填写鱼币兑换码", c)
		return
	}
	userId := utils.GetUserID(c)
	redeemCode, err := redeemService.CheckRedeemCode(redeemFishCoin.RedeemCode, userId)
	if err != nil {
		response.FailWithMessage("兑换码失效", c)
	}
	wallets, err := walletsService.GetCurrentUserWallets(userId)
	if err != nil {
		response.FailWithMessage("账号异常，未开通钱包", c)
	}
	if err := redeemService.RedeemFishCoin(&redeemCode, &wallets); err != nil {
		global.Logger.Error("兑换失败!", zap.Error(err))
		response.FailWithMessage("兑换失败", c)
	} else {
		response.OkWithMessage("兑换成功", c)
	}
}

// GetRedeemCodeList 获取兑换码列表
func (redeemApi *RedeemApi) GetRedeemCodeList(c *gin.Context) {
	var pageInfo request.RedeemCodeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := redeemService.GetRedeemCodeList(pageInfo); err != nil {
		global.Logger.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
