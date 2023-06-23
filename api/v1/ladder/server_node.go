package ladder

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/common/response"
	"github.com/oldweipro/gin-admin/model/ladder"
	ladderReq "github.com/oldweipro/gin-admin/model/ladder/request"
	"github.com/oldweipro/gin-admin/service"
	"github.com/oldweipro/gin-admin/utils"
	"go.uber.org/zap"
)

type ServerNodeApi struct {
}

var serverNodeService = service.ServiceGroupApp.LadderServiceGroup.ServerNodeService

// CreateServerNode 创建ServerNode
// @Tags ServerNode
// @Summary 创建ServerNode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ladder.ServerNode true "创建ServerNode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /serverNode/createServerNode [post]
func (serverNodeApi *ServerNodeApi) CreateServerNode(c *gin.Context) {
	var serverNode ladder.ServerNode
	err := c.ShouldBindJSON(&serverNode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	serverNode.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"ServerName":   {utils.NotEmpty()},
		"ServerHost":   {utils.NotEmpty()},
		"ServerPort":   {utils.NotEmpty()},
		"ServerStatus": {utils.NotEmpty()},
	}
	if err := utils.Verify(serverNode, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := serverNodeService.CreateServerNode(&serverNode); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteServerNode 删除ServerNode
// @Tags ServerNode
// @Summary 删除ServerNode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ladder.ServerNode true "删除ServerNode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /serverNode/deleteServerNode [delete]
func (serverNodeApi *ServerNodeApi) DeleteServerNode(c *gin.Context) {
	var serverNode ladder.ServerNode
	err := c.ShouldBindJSON(&serverNode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	serverNode.DeletedBy = utils.GetUserID(c)
	if err := serverNodeService.DeleteServerNode(serverNode); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteServerNodeByIds 批量删除ServerNode
// @Tags ServerNode
// @Summary 批量删除ServerNode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ServerNode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /serverNode/deleteServerNodeByIds [delete]
func (serverNodeApi *ServerNodeApi) DeleteServerNodeByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := serverNodeService.DeleteServerNodeByIds(IDS, deletedBy); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateServerNode 更新ServerNode
// @Tags ServerNode
// @Summary 更新ServerNode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ladder.ServerNode true "更新ServerNode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /serverNode/updateServerNode [put]
func (serverNodeApi *ServerNodeApi) UpdateServerNode(c *gin.Context) {
	var serverNode ladder.ServerNode
	err := c.ShouldBindJSON(&serverNode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	serverNode.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"ServerName":   {utils.NotEmpty()},
		"ServerHost":   {utils.NotEmpty()},
		"ServerPort":   {utils.NotEmpty()},
		"ServerStatus": {utils.NotEmpty()},
	}
	if err := utils.Verify(serverNode, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := serverNodeService.UpdateServerNode(serverNode); err != nil {
		global.Logger.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindServerNode 用id查询ServerNode
// @Tags ServerNode
// @Summary 用id查询ServerNode
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ladder.ServerNode true "用id查询ServerNode"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /serverNode/findServerNode [get]
func (serverNodeApi *ServerNodeApi) FindServerNode(c *gin.Context) {
	var serverNode ladder.ServerNode
	err := c.ShouldBindQuery(&serverNode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reserverNode, err := serverNodeService.GetServerNode(serverNode.ID); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reserverNode": reserverNode}, c)
	}
}

// GetServerNodeList 分页获取ServerNode列表
// @Tags ServerNode
// @Summary 分页获取ServerNode列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ladderReq.ServerNodeSearch true "分页获取ServerNode列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /serverNode/getServerNodeList [get]
func (serverNodeApi *ServerNodeApi) GetServerNodeList(c *gin.Context) {
	var pageInfo ladderReq.ServerNodeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := serverNodeService.GetServerNodeInfoList(pageInfo); err != nil {
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

func (serverNodeApi *ServerNodeApi) GetServerNodeLessList(c *gin.Context) {
	var pageInfo ladderReq.ServerNodeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := serverNodeService.GetServerNodeLessInfoList(pageInfo); err != nil {
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
