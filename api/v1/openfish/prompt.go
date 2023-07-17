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

type PromptApi struct {
}

var promptService = service.ServiceGroupApp.OpenfishServiceGroup.PromptService

// CreatePrompt 创建Prompt
// @Tags Prompt
// @Summary 创建Prompt
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body openfish.Prompt true "创建Prompt"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /prompt/createPrompt [post]
func (promptApi *PromptApi) CreatePrompt(c *gin.Context) {
	var prompt openfish.Prompt
	err := c.ShouldBindJSON(&prompt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":        {utils.NotEmpty()},
		"UseFee":      {utils.NotEmpty()},
		"Content":     {utils.NotEmpty()},
		"IsShare":     {utils.NotEmpty()},
		"Category":    {utils.NotEmpty()},
		"UseDuration": {utils.NotEmpty()},
	}
	if err := utils.Verify(prompt, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := promptService.CreatePrompt(&prompt); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePrompt 删除Prompt
// @Tags Prompt
// @Summary 删除Prompt
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body openfish.Prompt true "删除Prompt"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /prompt/deletePrompt [delete]
func (promptApi *PromptApi) DeletePrompt(c *gin.Context) {
	var prompt openfish.Prompt
	err := c.ShouldBindJSON(&prompt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := promptService.DeletePrompt(prompt); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePromptByIds 批量删除Prompt
// @Tags Prompt
// @Summary 批量删除Prompt
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Prompt"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /prompt/deletePromptByIds [delete]
func (promptApi *PromptApi) DeletePromptByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := promptService.DeletePromptByIds(IDS); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePrompt 更新Prompt
// @Tags Prompt
// @Summary 更新Prompt
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body openfish.Prompt true "更新Prompt"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /prompt/updatePrompt [put]
func (promptApi *PromptApi) UpdatePrompt(c *gin.Context) {
	var prompt openfish.Prompt
	err := c.ShouldBindJSON(&prompt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":        {utils.NotEmpty()},
		"UseFee":      {utils.NotEmpty()},
		"Content":     {utils.NotEmpty()},
		"IsShare":     {utils.NotEmpty()},
		"Category":    {utils.NotEmpty()},
		"UseDuration": {utils.NotEmpty()},
	}
	if err := utils.Verify(prompt, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := promptService.UpdatePrompt(prompt); err != nil {
		global.Logger.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPrompt 用id查询Prompt
// @Tags Prompt
// @Summary 用id查询Prompt
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query openfish.Prompt true "用id查询Prompt"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /prompt/findPrompt [get]
func (promptApi *PromptApi) FindPrompt(c *gin.Context) {
	var prompt openfish.Prompt
	err := c.ShouldBindQuery(&prompt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reprompt, err := promptService.GetPrompt(prompt.ID); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reprompt": reprompt}, c)
	}
}

// GetPromptList 分页获取Prompt列表
// @Tags Prompt
// @Summary 分页获取Prompt列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query openfishReq.PromptSearch true "分页获取Prompt列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /prompt/getPromptList [get]
func (promptApi *PromptApi) GetPromptList(c *gin.Context) {
	var pageInfo openfishReq.PromptSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := promptService.GetPromptInfoList(pageInfo); err != nil {
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

// GetCurrentUserPromptList TODO 获取当前用户prompt
func (promptApi *PromptApi) GetCurrentUserPromptList(c *gin.Context) {
	var pageInfo openfishReq.PromptSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := promptService.GetPromptInfoList(pageInfo); err != nil {
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
