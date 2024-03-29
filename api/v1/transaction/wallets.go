package transaction

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/common/response"
	"github.com/oldweipro/gin-admin/model/transaction"
	transactionReq "github.com/oldweipro/gin-admin/model/transaction/request"
	"github.com/oldweipro/gin-admin/service"
	"github.com/oldweipro/gin-admin/utils"
	"go.uber.org/zap"
)

type WalletsApi struct {
}

var walletsService = service.ServiceGroupApp.TransactionServiceGroup.WalletsService

// CreateWallets 创建Wallets
func (walletsApi *WalletsApi) CreateWallets(c *gin.Context) {
	var wallets transaction.Wallets
	err := c.ShouldBindJSON(&wallets)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"UserId": {utils.NotEmpty()},
	}
	if err := utils.Verify(wallets, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := walletsService.CreateWallets(&wallets); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWallets 删除Wallets
func (walletsApi *WalletsApi) DeleteWallets(c *gin.Context) {
	var wallets transaction.Wallets
	err := c.ShouldBindJSON(&wallets)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := walletsService.DeleteWallets(wallets); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWalletsByIds 批量删除Wallets
func (walletsApi *WalletsApi) DeleteWalletsByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := walletsService.DeleteWalletsByIds(IDS); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWallets 更新Wallets
func (walletsApi *WalletsApi) UpdateWallets(c *gin.Context) {
	var wallets transaction.Wallets
	err := c.ShouldBindJSON(&wallets)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"UserId": {utils.NotEmpty()},
	}
	if err := utils.Verify(wallets, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := walletsService.UpdateWallets(wallets); err != nil {
		global.Logger.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWallets 用id查询Wallets
func (walletsApi *WalletsApi) FindWallets(c *gin.Context) {
	var wallets transaction.Wallets
	err := c.ShouldBindQuery(&wallets)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rewallets, err := walletsService.GetWallets(wallets.ID); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewallets": rewallets}, c)
	}
}

// GetWalletsList 分页获取Wallets列表
func (walletsApi *WalletsApi) GetWalletsList(c *gin.Context) {
	var pageInfo transactionReq.WalletsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := walletsService.GetWalletsInfoList(pageInfo); err != nil {
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

// GetCurrentUserWallets 获取当前用户钱包
func (walletsApi *WalletsApi) GetCurrentUserWallets(c *gin.Context) {
	userID := utils.GetUserID(c)
	if wallets, err := walletsService.GetCurrentUserWallets(userID); err != nil {
		global.Logger.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(wallets, "获取成功", c)
	}
}
