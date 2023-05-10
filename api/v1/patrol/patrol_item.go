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

type PatrolItemApi struct {
}

var patrolItemService = service.ServiceGroupApp.PatrolServiceGroup.PatrolItemService

// CreatePatrolItem 创建PatrolItem
// @Tags PatrolItem
// @Summary 创建PatrolItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body patrol.PatrolItem true "创建PatrolItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /patrolItem/createPatrolItem [post]
func (patrolItemApi *PatrolItemApi) CreatePatrolItem(c *gin.Context) {
	var patrolItem patrol.PatrolItem
	err := c.ShouldBindJSON(&patrolItem)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	patrolItem.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"ItemTitle": {utils.NotEmpty()},
		"DeptId":    {utils.NotEmpty()},
	}
	if err := utils.Verify(patrolItem, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := patrolItemService.CreatePatrolItem(patrolItem); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePatrolItem 删除PatrolItem
// @Tags PatrolItem
// @Summary 删除PatrolItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body patrol.PatrolItem true "删除PatrolItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /patrolItem/deletePatrolItem [delete]
func (patrolItemApi *PatrolItemApi) DeletePatrolItem(c *gin.Context) {
	var patrolItem patrol.PatrolItem
	err := c.ShouldBindJSON(&patrolItem)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	patrolItem.DeletedBy = utils.GetUserID(c)
	if err := patrolItemService.DeletePatrolItem(patrolItem); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePatrolItemByIds 批量删除PatrolItem
// @Tags PatrolItem
// @Summary 批量删除PatrolItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PatrolItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /patrolItem/deletePatrolItemByIds [delete]
func (patrolItemApi *PatrolItemApi) DeletePatrolItemByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := patrolItemService.DeletePatrolItemByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePatrolItem 更新PatrolItem
// @Tags PatrolItem
// @Summary 更新PatrolItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body patrol.PatrolItem true "更新PatrolItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /patrolItem/updatePatrolItem [put]
func (patrolItemApi *PatrolItemApi) UpdatePatrolItem(c *gin.Context) {
	var patrolItem patrol.PatrolItem
	err := c.ShouldBindJSON(&patrolItem)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	patrolItem.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"ItemTitle": {utils.NotEmpty()},
		"DeptId":    {utils.NotEmpty()},
	}
	if err := utils.Verify(patrolItem, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := patrolItemService.UpdatePatrolItem(patrolItem); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPatrolItem 用id查询PatrolItem
// @Tags PatrolItem
// @Summary 用id查询PatrolItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query patrol.PatrolItem true "用id查询PatrolItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /patrolItem/findPatrolItem [get]
func (patrolItemApi *PatrolItemApi) FindPatrolItem(c *gin.Context) {
	var patrolItem patrol.PatrolItem
	err := c.ShouldBindQuery(&patrolItem)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if repatrolItem, err := patrolItemService.GetPatrolItem(patrolItem.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repatrolItem": repatrolItem}, c)
	}
}

// GetPatrolItemList 分页获取PatrolItem列表
// @Tags PatrolItem
// @Summary 分页获取PatrolItem列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query patrolReq.PatrolItemSearch true "分页获取PatrolItem列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /patrolItem/getPatrolItemList [get]
func (patrolItemApi *PatrolItemApi) GetPatrolItemList(c *gin.Context) {
	var pageInfo patrolReq.PatrolItemSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := patrolItemService.GetPatrolItemInfoList(pageInfo); err != nil {
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
