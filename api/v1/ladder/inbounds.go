package ladder

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/common/response"
	"github.com/oldweipro/gin-admin/model/ladder"
	ladderReq "github.com/oldweipro/gin-admin/model/ladder/request"
	"github.com/oldweipro/gin-admin/service"
	"github.com/oldweipro/gin-admin/utils"
	"go.uber.org/zap"
	"net/url"
)

type InboundsApi struct {
}

var inboundsService = service.ServiceGroupApp.LadderServiceGroup.InboundsService
var subscriptionPlanService = service.ServiceGroupApp.TransactionServiceGroup.SubscriptionPlanService

// CreateInbounds 创建Inbounds
// @Tags Inbounds
// @Summary 创建Inbounds
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ladder.Inbounds true "创建Inbounds"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /inbounds/createInbounds [post]
func (inboundsApi *InboundsApi) CreateInbounds(c *gin.Context) {
	var inbounds ladder.Inbounds
	err := c.ShouldBindJSON(&inbounds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	inbounds.CreatedBy = utils.GetUserID(c)
	if err := inboundsService.CreateInbounds(&inbounds); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteInbounds 删除Inbounds
// @Tags Inbounds
// @Summary 删除Inbounds
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ladder.Inbounds true "删除Inbounds"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /inbounds/deleteInbounds [delete]
func (inboundsApi *InboundsApi) DeleteInbounds(c *gin.Context) {
	var inbounds ladder.Inbounds
	err := c.ShouldBindJSON(&inbounds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	inbounds.DeletedBy = utils.GetUserID(c)
	if err := inboundsService.DeleteInbounds(inbounds); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteInboundsByIds 批量删除Inbounds
// @Tags Inbounds
// @Summary 批量删除Inbounds
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Inbounds"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /inbounds/deleteInboundsByIds [delete]
func (inboundsApi *InboundsApi) DeleteInboundsByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := inboundsService.DeleteInboundsByIds(IDS, deletedBy); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateInbounds 更新Inbounds
// @Tags Inbounds
// @Summary 更新Inbounds
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ladder.Inbounds true "更新Inbounds"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /inbounds/updateInbounds [put]
func (inboundsApi *InboundsApi) UpdateInbounds(c *gin.Context) {
	var inbounds ladder.Inbounds
	err := c.ShouldBindJSON(&inbounds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	inbounds.UpdatedBy = utils.GetUserID(c)
	if err := inboundsService.UpdateInbounds(inbounds); err != nil {
		global.Logger.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindInbounds 用id查询Inbounds
// @Tags Inbounds
// @Summary 用id查询Inbounds
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ladder.Inbounds true "用id查询Inbounds"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /inbounds/findInbounds [get]
func (inboundsApi *InboundsApi) FindInbounds(c *gin.Context) {
	var inbounds ladder.Inbounds
	err := c.ShouldBindQuery(&inbounds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reinbounds, err := inboundsService.GetInbounds(inbounds.ID); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reinbounds": reinbounds}, c)
	}
}

// FindInboundsLink 根据服务器ID和当前用户查询节点链接信息
func (inboundsApi *InboundsApi) FindInboundsLink(c *gin.Context) {
	// 先查询你的状态是否激活
	user, subErr := subscriptionPlanService.GetCurrentSubscriptionPlan(utils.GetUserID(c))
	if subErr != nil {
		response.FailWithMessage("请开通", c)
		return
	}
	if *user.Status == 0 {
		response.FailWithMessage("请选择您的订阅计划", c)
		return
	}
	var inbounds ladder.Inbounds
	err := c.ShouldBindQuery(&inbounds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userInfo := utils.GetUserInfo(c)
	// 查询服务器信息
	if serverNode, err := serverNodeService.GetServerNode(*inbounds.Sid); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		// 查询节点信息
		if inboundsLink, err := inboundsService.GetInboundsLink(*userInfo, *inbounds.Sid); err != nil {
			global.Logger.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("查询失败", c)
		} else {
			vMessLink := make(map[string]interface{})
			vMessLink["v"] = "2"
			vMessLink["ps"] = serverNode.Region
			vMessLink["add"] = serverNode.Domain
			vMessLink["port"] = inboundsLink.Port
			vMessLink["id"] = inboundsLink.ClientId
			vMessLink["aid"] = 0
			vMessLink["net"] = "tcp"
			vMessLink["type"] = "none"
			vMessLink["host"] = ""
			vMessLink["path"] = ""
			vMessLink["tls"] = "tls"
			vMessLinkJson, _ := json.MarshalIndent(vMessLink, "", "  ")
			inboundsLink.Link = string(vMessLinkJson)
			vMessLinkJsonBase64 := base64.StdEncoding.EncodeToString(vMessLinkJson)
			inboundsLink.Link64 = "vmess://" + vMessLinkJsonBase64

			// 组装clashSub订阅地址
			prefix := "https://subconverter.oldwei.com/sub?target=clash&url="
			subConverter := inboundsLink.Link64
			suffix := "&insert=false"
			clashInstall := prefix + url.QueryEscape(subConverter) + suffix
			inboundsLink.ClashSub = clashInstall
			response.OkWithData(gin.H{"inboundsData": inboundsLink, "domain": serverNode.Domain, "region": serverNode.Region, "expiryTime": inboundsLink.ExpiryTime}, c)
		}
	}
}

// SetInboundsLink 重置节点链接
func (inboundsApi *InboundsApi) SetInboundsLink(c *gin.Context) {
	var inbounds ladder.Inbounds
	err := c.ShouldBindJSON(&inbounds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 修改当前信息，更新到远端
	userInfo := utils.GetUserInfo(c)
	// 更新节点信息
	if err := inboundsService.SetInboundsLink(*userInfo, inbounds); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithMessage("重置成功", c)
	}
}

// GetInboundsList 分页获取Inbounds列表
// @Tags Inbounds
// @Summary 分页获取Inbounds列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ladderReq.InboundsSearch true "分页获取Inbounds列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /inbounds/getInboundsList [get]
func (inboundsApi *InboundsApi) GetInboundsList(c *gin.Context) {
	var pageInfo ladderReq.InboundsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := inboundsService.GetInboundsInfoList(pageInfo); err != nil {
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
