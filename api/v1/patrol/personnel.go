package patrol

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/common/response"
	"github.com/oldweipro/gin-admin/model/patrol"
	patrolReq "github.com/oldweipro/gin-admin/model/patrol/request"
	"github.com/oldweipro/gin-admin/service"
	"go.uber.org/zap"
)

type PersonnelApi struct {
}

var personnelService = service.ServiceGroupApp.PatrolServiceGroup.PersonnelService

// SyncPersonnel 同步人员信息
func (personnelApi *PersonnelApi) SyncPersonnel(c *gin.Context) {
	personnelService.SyncPersonnel()
	response.OkWithMessage("调用成功", c)
}

// SyncPersonnelImg 同步人员图片信息
func (personnelApi *PersonnelApi) SyncPersonnelImg(c *gin.Context) {
	personnelService.SyncPersonnelImg()
	response.OkWithMessage("调用成功", c)
}

// CreatePersonnel 创建Personnel
// @Tags Personnel
// @Summary 创建Personnel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body patrol.Personnel true "创建Personnel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /personnel/createPersonnel [post]
func (personnelApi *PersonnelApi) CreatePersonnel(c *gin.Context) {
	var personnel patrol.Personnel
	err := c.ShouldBindJSON(&personnel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := personnelService.CreatePersonnel(personnel); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePersonnel 删除Personnel
// @Tags Personnel
// @Summary 删除Personnel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body patrol.Personnel true "删除Personnel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /personnel/deletePersonnel [delete]
func (personnelApi *PersonnelApi) DeletePersonnel(c *gin.Context) {
	var personnel patrol.Personnel
	err := c.ShouldBindJSON(&personnel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := personnelService.DeletePersonnel(personnel); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePersonnelByIds 批量删除Personnel
// @Tags Personnel
// @Summary 批量删除Personnel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Personnel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /personnel/deletePersonnelByIds [delete]
func (personnelApi *PersonnelApi) DeletePersonnelByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := personnelService.DeletePersonnelByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePersonnel 更新Personnel
// @Tags Personnel
// @Summary 更新Personnel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body patrol.Personnel true "更新Personnel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /personnel/updatePersonnel [put]
func (personnelApi *PersonnelApi) UpdatePersonnel(c *gin.Context) {
	var personnel patrol.Personnel
	err := c.ShouldBindJSON(&personnel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := personnelService.UpdatePersonnel(personnel); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPersonnel 用id查询Personnel
// @Tags Personnel
// @Summary 用id查询Personnel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query patrol.Personnel true "用id查询Personnel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /personnel/findPersonnel [get]
func (personnelApi *PersonnelApi) FindPersonnel(c *gin.Context) {
	var personnel patrol.Personnel
	err := c.ShouldBindQuery(&personnel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if repersonnel, err := personnelService.GetPersonnel(personnel.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repersonnel": repersonnel}, c)
	}
}

// GetPersonnelList 分页获取Personnel列表
// @Tags Personnel
// @Summary 分页获取Personnel列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query patrolReq.PersonnelSearch true "分页获取Personnel列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /personnel/getPersonnelList [get]
func (personnelApi *PersonnelApi) GetPersonnelList(c *gin.Context) {
	var pageInfo patrolReq.PersonnelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := personnelService.GetPersonnelInfoList(pageInfo); err != nil {
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
