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

type PatrolTaskApi struct {
}

var patrolTaskService = service.ServiceGroupApp.PatrolServiceGroup.PatrolTaskService

// CreatePatrolTask 创建PatrolTask
// @Tags PatrolTask
// @Summary 创建PatrolTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body patrol.PatrolTask true "创建PatrolTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /patrolTask/createPatrolTask [post]
func (patrolTaskApi *PatrolTaskApi) CreatePatrolTask(c *gin.Context) {
	var patrolTask patrol.PatrolTask
	err := c.ShouldBindJSON(&patrolTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	patrolTask.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"TaskName":       {utils.NotEmpty()},
		"TaskItemIdList": {utils.NotEmpty()},
		"PatrolTimes":    {utils.NotEmpty()},
		"ClockMode":      {utils.NotEmpty()},
	}
	if err := utils.Verify(patrolTask, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := patrolTaskService.CreatePatrolTask(patrolTask); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePatrolTask 删除PatrolTask
// @Tags PatrolTask
// @Summary 删除PatrolTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body patrol.PatrolTask true "删除PatrolTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /patrolTask/deletePatrolTask [delete]
func (patrolTaskApi *PatrolTaskApi) DeletePatrolTask(c *gin.Context) {
	var patrolTask patrol.PatrolTask
	err := c.ShouldBindJSON(&patrolTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	patrolTask.DeletedBy = utils.GetUserID(c)
	if err := patrolTaskService.DeletePatrolTask(patrolTask); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePatrolTaskByIds 批量删除PatrolTask
// @Tags PatrolTask
// @Summary 批量删除PatrolTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PatrolTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /patrolTask/deletePatrolTaskByIds [delete]
func (patrolTaskApi *PatrolTaskApi) DeletePatrolTaskByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := patrolTaskService.DeletePatrolTaskByIds(IDS, deletedBy); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePatrolTask 更新PatrolTask
// @Tags PatrolTask
// @Summary 更新PatrolTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body patrol.PatrolTask true "更新PatrolTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /patrolTask/updatePatrolTask [put]
func (patrolTaskApi *PatrolTaskApi) UpdatePatrolTask(c *gin.Context) {
	var patrolTask patrol.PatrolTask
	err := c.ShouldBindJSON(&patrolTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	patrolTask.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"TaskName":       {utils.NotEmpty()},
		"TaskItemIdList": {utils.NotEmpty()},
		"PatrolTimes":    {utils.NotEmpty()},
		"ClockMode":      {utils.NotEmpty()},
	}
	if err := utils.Verify(patrolTask, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := patrolTaskService.UpdatePatrolTask(patrolTask); err != nil {
		global.Logger.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPatrolTask 用id查询PatrolTask
// @Tags PatrolTask
// @Summary 用id查询PatrolTask
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query patrol.PatrolTask true "用id查询PatrolTask"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /patrolTask/findPatrolTask [get]
func (patrolTaskApi *PatrolTaskApi) FindPatrolTask(c *gin.Context) {
	var patrolTask patrol.PatrolTask
	err := c.ShouldBindQuery(&patrolTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if repatrolTask, err := patrolTaskService.GetPatrolTask(patrolTask.ID); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repatrolTask": repatrolTask}, c)
	}
}

// GetPatrolTaskList 分页获取PatrolTask列表
// @Tags PatrolTask
// @Summary 分页获取PatrolTask列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query patrolReq.PatrolTaskSearch true "分页获取PatrolTask列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /patrolTask/getPatrolTaskList [get]
func (patrolTaskApi *PatrolTaskApi) GetPatrolTaskList(c *gin.Context) {
	var pageInfo patrolReq.PatrolTaskSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := patrolTaskService.GetPatrolTaskInfoList(pageInfo); err != nil {
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
