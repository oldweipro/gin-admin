package patrol

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/common/response"
	"github.com/oldweipro/gin-admin/model/patrol"
	patrolReq "github.com/oldweipro/gin-admin/model/patrol/request"
	"github.com/oldweipro/gin-admin/service"
	"github.com/oldweipro/gin-admin/utils"
	"go.uber.org/zap"
	"strings"
)

type AccountApi struct {
}

var accountService = service.ServiceGroupApp.PatrolServiceGroup.AccountService

// CreateAccount 创建Account
// @Tags Account
// @Summary 创建Account
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body patrol.Account true "创建Account"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /account/createAccount [post]
func (accountApi *AccountApi) CreateAccount(c *gin.Context) {
	var account patrol.Account
	err := c.ShouldBindJSON(&account)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	account.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Account": {utils.NotEmpty()},
	}
	if err := utils.Verify(account, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := accountService.CreateAccount(account); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAccount 删除Account
// @Tags Account
// @Summary 删除Account
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body patrol.Account true "删除Account"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /account/deleteAccount [delete]
func (accountApi *AccountApi) DeleteAccount(c *gin.Context) {
	var account patrol.Account
	err := c.ShouldBindJSON(&account)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	account.DeletedBy = utils.GetUserID(c)
	if err := accountService.DeleteAccount(account); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAccountByIds 批量删除Account
// @Tags Account
// @Summary 批量删除Account
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Account"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /account/deleteAccountByIds [delete]
func (accountApi *AccountApi) DeleteAccountByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := accountService.DeleteAccountByIds(IDS, deletedBy); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAccount 更新Account
// @Tags Account
// @Summary 更新Account
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body patrol.Account true "更新Account"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /account/updateAccount [put]
func (accountApi *AccountApi) UpdateAccount(c *gin.Context) {
	var account patrol.Account
	err := c.ShouldBindJSON(&account)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	account.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Account": {utils.NotEmpty()},
	}
	if err := utils.Verify(account, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := accountService.UpdateAccount(account); err != nil {
		global.Logger.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAccount 用id查询Account
// @Tags Account
// @Summary 用id查询Account
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query patrol.Account true "用id查询Account"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /account/findAccount [get]
func (accountApi *AccountApi) FindAccount(c *gin.Context) {
	var account patrol.Account
	err := c.ShouldBindQuery(&account)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reaccount, err := accountService.GetAccount(account.ID); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reaccount": reaccount}, c)
	}
}

// GetAccountList 分页获取Account列表
// @Tags Account
// @Summary 分页获取Account列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query patrolReq.AccountSearch true "分页获取Account列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /account/getAccountList [get]
func (accountApi *AccountApi) GetAccountList(c *gin.Context) {
	var pageInfo patrolReq.AccountSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := accountService.GetAccountInfoList(pageInfo); err != nil {
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

func (accountApi *AccountApi) LoginGameAccount(c *gin.Context) {
	var account patrol.Account
	err := c.ShouldBindJSON(&account)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	account.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Account": {utils.NotEmpty()},
	}
	if err := utils.Verify(account, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	loginStatus := 1
	account.LoginStatus = &loginStatus
	gameAccount, _ := accountService.LoginGameAccount(account)
	if strings.Contains(gameAccount, "<script type=\"text/javascript\">top.location=\"http:") {
		response.OkWithMessage("登陆成功", c)
	} else {
		response.FailWithMessage("登陆失败!页面返回错误", c)
	}
}
