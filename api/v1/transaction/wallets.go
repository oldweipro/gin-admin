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
// @Tags Wallets
// @Summary 创建Wallets
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body openfish.Wallets true "创建Wallets"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wallets/createWallets [post]
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
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWallets 删除Wallets
// @Tags Wallets
// @Summary 删除Wallets
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body openfish.Wallets true "删除Wallets"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wallets/deleteWallets [delete]
func (walletsApi *WalletsApi) DeleteWallets(c *gin.Context) {
	var wallets transaction.Wallets
	err := c.ShouldBindJSON(&wallets)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := walletsService.DeleteWallets(wallets); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWalletsByIds 批量删除Wallets
// @Tags Wallets
// @Summary 批量删除Wallets
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Wallets"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wallets/deleteWalletsByIds [delete]
func (walletsApi *WalletsApi) DeleteWalletsByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := walletsService.DeleteWalletsByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWallets 更新Wallets
// @Tags Wallets
// @Summary 更新Wallets
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body openfish.Wallets true "更新Wallets"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wallets/updateWallets [put]
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
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWallets 用id查询Wallets
// @Tags Wallets
// @Summary 用id查询Wallets
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query openfish.Wallets true "用id查询Wallets"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wallets/findWallets [get]
func (walletsApi *WalletsApi) FindWallets(c *gin.Context) {
	var wallets transaction.Wallets
	err := c.ShouldBindQuery(&wallets)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rewallets, err := walletsService.GetWallets(wallets.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewallets": rewallets}, c)
	}
}

// GetWalletsList 分页获取Wallets列表
// @Tags Wallets
// @Summary 分页获取Wallets列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query openfishReq.WalletsSearch true "分页获取Wallets列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wallets/getWalletsList [get]
func (walletsApi *WalletsApi) GetWalletsList(c *gin.Context) {
	var pageInfo transactionReq.WalletsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := walletsService.GetWalletsInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
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
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(wallets, "获取成功", c)
	}
}
