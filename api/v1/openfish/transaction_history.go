package openfish

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/common/response"
	"github.com/oldweipro/gin-admin/model/openfish"
	openfishReq "github.com/oldweipro/gin-admin/model/openfish/request"
	"github.com/oldweipro/gin-admin/service"
	"github.com/oldweipro/gin-admin/utils"
	"go.uber.org/zap"
)

type TransactionHistoryApi struct {
}

var transactionHistoryService = service.ServiceGroupApp.OpenfishServiceGroup.TransactionHistoryService

// CreateTransactionHistory 创建TransactionHistory
// @Tags TransactionHistory
// @Summary 创建TransactionHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body openfish.TransactionHistory true "创建TransactionHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /transactionHistory/createTransactionHistory [post]
func (transactionHistoryApi *TransactionHistoryApi) CreateTransactionHistory(c *gin.Context) {
	var transactionHistory openfish.TransactionHistory
	err := c.ShouldBindJSON(&transactionHistory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	transactionHistory.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"UserId":    {utils.NotEmpty()},
		"WalletId":  {utils.NotEmpty()},
		"TypeEnum":  {utils.NotEmpty()},
		"Amount":    {utils.NotEmpty()},
		"ProductId": {utils.NotEmpty()},
	}
	if err := utils.Verify(transactionHistory, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := transactionHistoryService.CreateTransactionHistory(&transactionHistory); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTransactionHistory 删除TransactionHistory
// @Tags TransactionHistory
// @Summary 删除TransactionHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body openfish.TransactionHistory true "删除TransactionHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /transactionHistory/deleteTransactionHistory [delete]
func (transactionHistoryApi *TransactionHistoryApi) DeleteTransactionHistory(c *gin.Context) {
	var transactionHistory openfish.TransactionHistory
	err := c.ShouldBindJSON(&transactionHistory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	transactionHistory.DeletedBy = utils.GetUserID(c)
	if err := transactionHistoryService.DeleteTransactionHistory(transactionHistory); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTransactionHistoryByIds 批量删除TransactionHistory
// @Tags TransactionHistory
// @Summary 批量删除TransactionHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TransactionHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /transactionHistory/deleteTransactionHistoryByIds [delete]
func (transactionHistoryApi *TransactionHistoryApi) DeleteTransactionHistoryByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := transactionHistoryService.DeleteTransactionHistoryByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTransactionHistory 更新TransactionHistory
// @Tags TransactionHistory
// @Summary 更新TransactionHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body openfish.TransactionHistory true "更新TransactionHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /transactionHistory/updateTransactionHistory [put]
func (transactionHistoryApi *TransactionHistoryApi) UpdateTransactionHistory(c *gin.Context) {
	var transactionHistory openfish.TransactionHistory
	err := c.ShouldBindJSON(&transactionHistory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	transactionHistory.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"UserId":    {utils.NotEmpty()},
		"WalletId":  {utils.NotEmpty()},
		"TypeEnum":  {utils.NotEmpty()},
		"Amount":    {utils.NotEmpty()},
		"ProductId": {utils.NotEmpty()},
	}
	if err := utils.Verify(transactionHistory, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := transactionHistoryService.UpdateTransactionHistory(transactionHistory); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTransactionHistory 用id查询TransactionHistory
// @Tags TransactionHistory
// @Summary 用id查询TransactionHistory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query openfish.TransactionHistory true "用id查询TransactionHistory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /transactionHistory/findTransactionHistory [get]
func (transactionHistoryApi *TransactionHistoryApi) FindTransactionHistory(c *gin.Context) {
	var transactionHistory openfish.TransactionHistory
	err := c.ShouldBindQuery(&transactionHistory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if retransactionHistory, err := transactionHistoryService.GetTransactionHistory(transactionHistory.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retransactionHistory": retransactionHistory}, c)
	}
}

// GetTransactionHistoryList 分页获取TransactionHistory列表
// @Tags TransactionHistory
// @Summary 分页获取TransactionHistory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query openfishReq.TransactionHistorySearch true "分页获取TransactionHistory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /transactionHistory/getTransactionHistoryList [get]
func (transactionHistoryApi *TransactionHistoryApi) GetTransactionHistoryList(c *gin.Context) {
	var pageInfo openfishReq.TransactionHistorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := transactionHistoryService.GetTransactionHistoryInfoList(pageInfo); err != nil {
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
