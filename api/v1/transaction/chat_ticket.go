package transaction

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/common/response"
	"github.com/oldweipro/gin-admin/model/transaction"
	openfishReq "github.com/oldweipro/gin-admin/model/transaction/request"
	"github.com/oldweipro/gin-admin/service"
	"github.com/oldweipro/gin-admin/utils"
	"go.uber.org/zap"
)

type ChatTicketApi struct {
}

var chatTicketService = service.ServiceGroupApp.TransactionServiceGroup.ChatTicketService

// CreateChatTicket 创建ChatTicket
// @Tags ChatTicket
// @Summary 创建ChatTicket
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body openfish.ChatTicket true "创建ChatTicket"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /chatTicket/createChatTicket [post]
func (chatTicketApi *ChatTicketApi) CreateChatTicket(c *gin.Context) {
	var chatTicket transaction.ChatTicket
	err := c.ShouldBindJSON(&chatTicket)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	chatTicket.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Amount":         {utils.NotEmpty()},
		"ExpirationTime": {utils.NotEmpty()},
		"TicketName":     {utils.NotEmpty()},
	}
	if err := utils.Verify(chatTicket, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := chatTicketService.CreateChatTicket(&chatTicket); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteChatTicket 删除ChatTicket
// @Tags ChatTicket
// @Summary 删除ChatTicket
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body openfish.ChatTicket true "删除ChatTicket"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /chatTicket/deleteChatTicket [delete]
func (chatTicketApi *ChatTicketApi) DeleteChatTicket(c *gin.Context) {
	var chatTicket transaction.ChatTicket
	err := c.ShouldBindJSON(&chatTicket)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	chatTicket.DeletedBy = utils.GetUserID(c)
	if err := chatTicketService.DeleteChatTicket(chatTicket); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteChatTicketByIds 批量删除ChatTicket
// @Tags ChatTicket
// @Summary 批量删除ChatTicket
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ChatTicket"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /chatTicket/deleteChatTicketByIds [delete]
func (chatTicketApi *ChatTicketApi) DeleteChatTicketByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := chatTicketService.DeleteChatTicketByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateChatTicket 更新ChatTicket
// @Tags ChatTicket
// @Summary 更新ChatTicket
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body openfish.ChatTicket true "更新ChatTicket"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /chatTicket/updateChatTicket [put]
func (chatTicketApi *ChatTicketApi) UpdateChatTicket(c *gin.Context) {
	var chatTicket transaction.ChatTicket
	err := c.ShouldBindJSON(&chatTicket)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	chatTicket.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Amount":         {utils.NotEmpty()},
		"ExpirationTime": {utils.NotEmpty()},
		"TicketName":     {utils.NotEmpty()},
	}
	if err := utils.Verify(chatTicket, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := chatTicketService.UpdateChatTicket(chatTicket); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindChatTicket 用id查询ChatTicket
// @Tags ChatTicket
// @Summary 用id查询ChatTicket
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query openfish.ChatTicket true "用id查询ChatTicket"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /chatTicket/findChatTicket [get]
func (chatTicketApi *ChatTicketApi) FindChatTicket(c *gin.Context) {
	var chatTicket transaction.ChatTicket
	err := c.ShouldBindQuery(&chatTicket)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rechatTicket, err := chatTicketService.GetChatTicket(chatTicket.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rechatTicket": rechatTicket}, c)
	}
}

// GetChatTicketList 分页获取ChatTicket列表
// @Tags ChatTicket
// @Summary 分页获取ChatTicket列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query openfishReq.ChatTicketSearch true "分页获取ChatTicket列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /chatTicket/getChatTicketList [get]
func (chatTicketApi *ChatTicketApi) GetChatTicketList(c *gin.Context) {
	var pageInfo openfishReq.ChatTicketSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := chatTicketService.GetChatTicketInfoList(pageInfo); err != nil {
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
