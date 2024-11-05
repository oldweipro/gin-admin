package system

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/model/common/response"
	"github.com/oldweipro/gin-admin/model/system/request"
	"github.com/oldweipro/gin-admin/pkg/app"
	utils2 "github.com/oldweipro/gin-admin/pkg/utils"
	"go.uber.org/zap"
)

type AutoCodeTemplateApi struct{}

// Preview
// @Tags      AutoCodeTemplate
// @Summary   预览创建后的代码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.AutoCode                                      true  "预览创建代码"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "预览创建后的代码"
// @Router    /autoCode/preview [post]
func (a *AutoCodeTemplateApi) Preview(c *gin.Context) {
	var info request.AutoCode
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils2.Verify(info, utils2.AutoCodeVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = info.Pretreatment()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	info.PackageT = utils2.FirstUpper(info.Package)
	autoCode, err := autoCodeTemplateService.Preview(c.Request.Context(), info)
	if err != nil {
		app.Logger.Error("预览失败!", zap.Error(err))
		response.FailWithMessage("预览失败", c)
	} else {
		response.OkWithDetailed(gin.H{"autoCode": autoCode}, "预览成功", c)
	}
}

// Create
// @Tags      AutoCodeTemplate
// @Summary   自动代码模板
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.AutoCode  true  "创建自动代码"
// @Success   200   {string}  string                 "{"success":true,"data":{},"msg":"创建成功"}"
// @Router    /autoCode/createTemp [post]
func (a *AutoCodeTemplateApi) Create(c *gin.Context) {
	var info request.AutoCode
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils2.Verify(info, utils2.AutoCodeVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = info.Pretreatment()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = autoCodeTemplateService.Create(c.Request.Context(), info)
	if err != nil {
		app.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// Create
// @Tags      AddFunc
// @Summary   增加方法
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.AutoCode  true  "增加方法"
// @Success   200   {string}  string                 "{"success":true,"data":{},"msg":"创建成功"}"
// @Router    /autoCode/addFunc [post]
func (a *AutoCodeTemplateApi) AddFunc(c *gin.Context) {
	var info request.AutoFunc
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = autoCodeTemplateService.AddFunc(info)
	if err != nil {
		app.Logger.Error("注入失败!", zap.Error(err))
		response.FailWithMessage("注入失败", c)
	} else {
		response.OkWithMessage("注入成功", c)
	}
}
