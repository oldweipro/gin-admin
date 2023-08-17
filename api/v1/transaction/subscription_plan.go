package transaction

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/common/response"
	"github.com/oldweipro/gin-admin/model/transaction"
	"github.com/oldweipro/gin-admin/service"
	"github.com/oldweipro/gin-admin/utils"
	"go.uber.org/zap"
)

type SubscriptionPlanApi struct {
}

var subscriptionPlanService = service.ServiceGroupApp.TransactionServiceGroup.SubscriptionPlanService

// CreateSubscriptionPlan 创建SubscriptionPlan
// @Tags SubscriptionPlan
// @Summary 创建SubscriptionPlan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body openfish.SubscriptionPlan true "创建SubscriptionPlan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /subscriptionPlan/createSubscriptionPlan [post]
func (subscriptionPlanApi *SubscriptionPlanApi) CreateSubscriptionPlan(c *gin.Context) {
	var subscriptionPlan transaction.SubscriptionPlan
	err := c.ShouldBindJSON(&subscriptionPlan)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":  {utils.NotEmpty()},
		"Price": {utils.NotEmpty()},
	}
	if err := utils.Verify(subscriptionPlan, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := subscriptionPlanService.CreateSubscriptionPlan(&subscriptionPlan); err != nil {
		global.Logger.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSubscriptionPlan 删除SubscriptionPlan
// @Tags SubscriptionPlan
// @Summary 删除SubscriptionPlan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body openfish.SubscriptionPlan true "删除SubscriptionPlan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /subscriptionPlan/deleteSubscriptionPlan [delete]
func (subscriptionPlanApi *SubscriptionPlanApi) DeleteSubscriptionPlan(c *gin.Context) {
	var subscriptionPlan transaction.SubscriptionPlan
	err := c.ShouldBindJSON(&subscriptionPlan)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := subscriptionPlanService.DeleteSubscriptionPlan(subscriptionPlan); err != nil {
		global.Logger.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSubscriptionPlanByIds 批量删除SubscriptionPlan
// @Tags SubscriptionPlan
// @Summary 批量删除SubscriptionPlan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SubscriptionPlan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /subscriptionPlan/deleteSubscriptionPlanByIds [delete]
func (subscriptionPlanApi *SubscriptionPlanApi) DeleteSubscriptionPlanByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := subscriptionPlanService.DeleteSubscriptionPlanByIds(IDS); err != nil {
		global.Logger.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSubscriptionPlan 更新SubscriptionPlan
// @Tags SubscriptionPlan
// @Summary 更新SubscriptionPlan
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body openfish.SubscriptionPlan true "更新SubscriptionPlan"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /subscriptionPlan/updateSubscriptionPlan [put]
func (subscriptionPlanApi *SubscriptionPlanApi) UpdateSubscriptionPlan(c *gin.Context) {
	var subscriptionPlan transaction.SubscriptionPlan
	err := c.ShouldBindJSON(&subscriptionPlan)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":  {utils.NotEmpty()},
		"Price": {utils.NotEmpty()},
	}
	if err := utils.Verify(subscriptionPlan, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := subscriptionPlanService.UpdateSubscriptionPlan(subscriptionPlan); err != nil {
		global.Logger.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// GetSubscriptionPlan 用id查询SubscriptionPlan
func (subscriptionPlanApi *SubscriptionPlanApi) GetSubscriptionPlan(c *gin.Context) {
	var subscriptionPlan transaction.SubscriptionPlan
	err := c.ShouldBindQuery(&subscriptionPlan)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resubscriptionPlan, err := subscriptionPlanService.GetSubscriptionPlan(subscriptionPlan.ID); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resubscriptionPlan": resubscriptionPlan}, c)
	}
}

// GetCurrentSubscriptionPlan 查询当前用户订阅计划
func (subscriptionPlanApi *SubscriptionPlanApi) GetCurrentSubscriptionPlan(c *gin.Context) {
	userID := utils.GetUserID(c)
	if subscriptionUser, err := subscriptionPlanService.GetCurrentSubscriptionPlan(userID); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.OkWithData(gin.H{"subscriptionUser": subscriptionUser}, c)
	} else {
		response.OkWithData(gin.H{"subscriptionUser": subscriptionUser}, c)
	}
}

// GetSubscriptionPlanList 分页获取SubscriptionPlan列表
//func (subscriptionPlanApi *SubscriptionPlanApi) GetSubscriptionPlanList(c *gin.Context) {
//	var pageInfo openfishReq.SubscriptionPlanSearch
//	err := c.ShouldBindQuery(&pageInfo)
//	if err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//	if list, total, err := subscriptionPlanService.GetSubscriptionPlanInfoList(pageInfo); err != nil {
//		global.Logger.Error("获取失败!", zap.Error(err))
//		response.FailWithMessage("获取失败", c)
//	} else {
//		response.OkWithDetailed(response.PageResult{
//			List:     list,
//			Total:    total,
//			Page:     pageInfo.Page,
//			PageSize: pageInfo.PageSize,
//		}, "获取成功", c)
//	}
//}
