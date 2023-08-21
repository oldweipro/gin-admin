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
	"sync"
)

type ChatTicketApi struct {
}

var checkInStatus sync.Map

var chatTicketService = service.ServiceGroupApp.TransactionServiceGroup.ChatTicketService

// CreateChatTicket 创建ChatTicket
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
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteChatTicket 删除ChatTicket
func (chatTicketApi *ChatTicketApi) DeleteChatTicket(c *gin.Context) {
	var chatTicket transaction.ChatTicket
	err := c.ShouldBindJSON(&chatTicket)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	chatTicket.DeletedBy = utils.GetUserID(c)
	if err := chatTicketService.DeleteChatTicket(chatTicket); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteChatTicketByIds 批量删除ChatTicket
func (chatTicketApi *ChatTicketApi) DeleteChatTicketByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := chatTicketService.DeleteChatTicketByIds(IDS, deletedBy); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateChatTicket 更新ChatTicket
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
		global.Logger.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindChatTicket 用id查询ChatTicket
func (chatTicketApi *ChatTicketApi) FindChatTicket(c *gin.Context) {
	var chatTicket transaction.ChatTicket
	err := c.ShouldBindQuery(&chatTicket)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rechatTicket, err := chatTicketService.GetChatTicket(chatTicket.ID); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rechatTicket": rechatTicket}, c)
	}
}

// GetChatTicketList 分页获取ChatTicket列表
func (chatTicketApi *ChatTicketApi) GetChatTicketList(c *gin.Context) {
	var pageInfo openfishReq.ChatTicketSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := chatTicketService.GetChatTicketInfoList(pageInfo); err != nil {
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

// HandleValidateChatTicket 验证鱼币兑换码
func (chatTicketApi *ChatTicketApi) HandleValidateChatTicket(c *gin.Context) {
	var chatTicket transaction.ChatTicket
	err := c.ShouldBindJSON(&chatTicket)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if chatTicket.TicketValue == "" {
		response.FailWithMessage("请填写鱼币兑换码", c)
		return
	}
	userId := utils.GetUserID(c)
	wallets, err := walletsService.GetCurrentUserWallets(userId)
	if err != nil {
		response.FailWithMessage("账号异常，您的钱包未创建", c)
		return
	}
	if err := chatTicketService.HandleValidateChatTicket(chatTicket.TicketValue, &wallets); err != nil {
		global.Logger.Error("验证失败!", zap.Error(err))
		response.FailWithMessage("验证失败", c)
	} else {
		response.OkWithMessage("验证成果", c)
	}
}

// CheckIn 签到
func (chatTicketApi *ChatTicketApi) CheckIn(c *gin.Context) {
	userId := utils.GetUserID(c)
	_, loaded := checkInStatus.LoadOrStore(userId, true)
	if loaded {
		response.FailStatusTooManyRequestsWithDetailed(nil, "请求过多", c)
		return
	}
	defer checkInStatus.Delete(userId)
	// 查询当天是否已签到
	count, err := historyService.GetTodayTransactionHistoryByCurrentUser(userId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if count > 0 {
		response.FailWithMessage("今日已签到", c)
		return
	}
	wallets, err := walletsService.GetCurrentUserWallets(userId)
	if err != nil {
		response.FailWithMessage("账号异常，您的钱包未创建", c)
		return
	}
	// 增加签到记录和金额
	if err := chatTicketService.HandleCheckIn(&wallets); err != nil {
		global.Logger.Error("签到失败!", zap.Error(err))
		response.FailWithMessage("签到失败", c)
	} else {
		response.OkWithMessage("签到成功", c)
	}
}
