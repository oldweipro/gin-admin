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
	"sync"
)

type MailAccountApi struct {
}

var mailAccountService = service.ServiceGroupApp.OpenfishServiceGroup.MailAccountService

// RefreshClaudeChat 产生一次Claude对话
func (mailAccountApi *MailAccountApi) RefreshClaudeChat(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := mailAccountService.RefreshClaudeChat(IDS); err != nil {
		global.Logger.Error("操作失败!", zap.Error(err))
		response.FailWithMessage("操作失败", c)
	} else {
		response.OkWithMessage("操作成功", c)
	}
}

var refreshOpenaiAccessTokenStatus sync.Map

// RefreshOpenaiAccessToken 刷新 openai AccessToken
func (mailAccountApi *MailAccountApi) RefreshOpenaiAccessToken(c *gin.Context) {
	// 获取用户ID
	userID := utils.GetUserID(c)
	// 检查用户的请求状态
	_, loaded := refreshOpenaiAccessTokenStatus.LoadOrStore(userID, true)
	if loaded {
		c.JSON(429, gin.H{"msg": "太多请求了"})
		return
	}
	defer refreshOpenaiAccessTokenStatus.Delete(userID) // 在处理完毕后删除用户的请求状态

	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := mailAccountService.RefreshOpenaiAccessToken(IDS); err != nil {
		global.Logger.Error("操作失败!", zap.Error(err))
		response.FailWithMessage("操作失败", c)
	} else {
		response.OkWithMessage("操作成功", c)
	}
}

// SyncOpenaiInfo 同步openai的信息: sk到期时间、余额
func (mailAccountApi *MailAccountApi) SyncOpenaiInfo(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := mailAccountService.SyncOpenaiInfo(IDS); err != nil {
		global.Logger.Error("操作失败!", zap.Error(err))
		response.FailWithMessage("操作失败", c)
	} else {
		response.OkWithMessage("操作成功", c)
	}
}

// CreateMailAccount 创建
func (mailAccountApi *MailAccountApi) CreateMailAccount(c *gin.Context) {
	var mailAccount openfish.MailAccount
	err := c.ShouldBindJSON(&mailAccount)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Username": {utils.NotEmpty()},
		"Password": {utils.NotEmpty()},
	}
	if err := utils.Verify(mailAccount, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := mailAccountService.CreateMailAccount(&mailAccount); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMailAccount 删除
func (mailAccountApi *MailAccountApi) DeleteMailAccount(c *gin.Context) {
	var mailAccount openfish.MailAccount
	err := c.ShouldBindJSON(&mailAccount)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := mailAccountService.DeleteMailAccount(mailAccount); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMailAccountByIds 批量删除
func (mailAccountApi *MailAccountApi) DeleteMailAccountByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := mailAccountService.DeleteMailAccountByIds(IDS); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMailAccount 更新
func (mailAccountApi *MailAccountApi) UpdateMailAccount(c *gin.Context) {
	var mailAccount openfish.MailAccount
	err := c.ShouldBindJSON(&mailAccount)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Username": {utils.NotEmpty()},
		"Password": {utils.NotEmpty()},
	}
	if err := utils.Verify(mailAccount, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := mailAccountService.UpdateMailAccount(mailAccount); err != nil {
		global.Logger.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMailAccount 用id查询
func (mailAccountApi *MailAccountApi) FindMailAccount(c *gin.Context) {
	var mailAccount openfish.MailAccount
	err := c.ShouldBindQuery(&mailAccount)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if mailAccountData, err := mailAccountService.GetMailAccount(mailAccount.ID); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(mailAccountData, c)
	}
}

// GetMailAccountList GetPromptList 分页获取 mailAccount 列表
func (mailAccountApi *MailAccountApi) GetMailAccountList(c *gin.Context) {
	var pageInfo openfishReq.MailAccountSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := mailAccountService.GetMailAccountInfoList(pageInfo); err != nil {
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
