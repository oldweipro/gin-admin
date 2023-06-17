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
)

type PatrolSiteApi struct {
}

var patrolSiteService = service.ServiceGroupApp.PatrolServiceGroup.PatrolSiteService

// CreatePatrolSite 创建PatrolSite
// @Tags PatrolSite
// @Summary 创建PatrolSite
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body patrol.PatrolSite true "创建PatrolSite"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /patrolSite/createPatrolSite [post]
func (patrolSiteApi *PatrolSiteApi) CreatePatrolSite(c *gin.Context) {
	var patrolSite patrol.PatrolSite
	err := c.ShouldBindJSON(&patrolSite)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	patrolSite.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"SiteName": {utils.NotEmpty()},
		"DeptId":   {utils.NotEmpty()},
	}
	if err := utils.Verify(patrolSite, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := patrolSiteService.CreatePatrolSite(patrolSite); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePatrolSite 删除PatrolSite
// @Tags PatrolSite
// @Summary 删除PatrolSite
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body patrol.PatrolSite true "删除PatrolSite"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /patrolSite/deletePatrolSite [delete]
func (patrolSiteApi *PatrolSiteApi) DeletePatrolSite(c *gin.Context) {
	var patrolSite patrol.PatrolSite
	err := c.ShouldBindJSON(&patrolSite)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	patrolSite.DeletedBy = utils.GetUserID(c)
	if err := patrolSiteService.DeletePatrolSite(patrolSite); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePatrolSiteByIds 批量删除PatrolSite
// @Tags PatrolSite
// @Summary 批量删除PatrolSite
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PatrolSite"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /patrolSite/deletePatrolSiteByIds [delete]
func (patrolSiteApi *PatrolSiteApi) DeletePatrolSiteByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := patrolSiteService.DeletePatrolSiteByIds(IDS, deletedBy); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePatrolSite 更新PatrolSite
// @Tags PatrolSite
// @Summary 更新PatrolSite
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body patrol.PatrolSite true "更新PatrolSite"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /patrolSite/updatePatrolSite [put]
func (patrolSiteApi *PatrolSiteApi) UpdatePatrolSite(c *gin.Context) {
	var patrolSite patrol.PatrolSite
	err := c.ShouldBindJSON(&patrolSite)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	patrolSite.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"SiteName": {utils.NotEmpty()},
		"DeptId":   {utils.NotEmpty()},
	}
	if err := utils.Verify(patrolSite, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := patrolSiteService.UpdatePatrolSite(patrolSite); err != nil {
		global.Logger.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPatrolSite 用id查询PatrolSite
// @Tags PatrolSite
// @Summary 用id查询PatrolSite
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query patrol.PatrolSite true "用id查询PatrolSite"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /patrolSite/findPatrolSite [get]
func (patrolSiteApi *PatrolSiteApi) FindPatrolSite(c *gin.Context) {
	var patrolSite patrol.PatrolSite
	err := c.ShouldBindQuery(&patrolSite)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if repatrolSite, err := patrolSiteService.GetPatrolSite(patrolSite.ID); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repatrolSite": repatrolSite}, c)
	}
}

// GetPatrolSiteList 分页获取PatrolSite列表
// @Tags PatrolSite
// @Summary 分页获取PatrolSite列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query patrolReq.PatrolSiteSearch true "分页获取PatrolSite列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /patrolSite/getPatrolSiteList [get]
func (patrolSiteApi *PatrolSiteApi) GetPatrolSiteList(c *gin.Context) {
	var pageInfo patrolReq.PatrolSiteSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := patrolSiteService.GetPatrolSiteInfoList(pageInfo); err != nil {
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
