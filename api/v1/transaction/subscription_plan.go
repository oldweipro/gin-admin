package transaction

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/common/response"
	"github.com/oldweipro/gin-admin/model/transaction"
	transactionRequest "github.com/oldweipro/gin-admin/model/transaction/request"
	"github.com/oldweipro/gin-admin/service"
	"github.com/oldweipro/gin-admin/utils"
	"go.uber.org/zap"
)

type SubscriptionPlanApi struct {
}

var subscriptionPlanService = service.ServiceGroupApp.TransactionServiceGroup.SubscriptionPlanService

// CreateSubscriptionPlan 创建SubscriptionPlan
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
		response.OkWithData(resubscriptionPlan, c)
	}
}

// GetCurrentSubscriptionPlan 查询当前用户订阅计划
func (subscriptionPlanApi *SubscriptionPlanApi) GetCurrentSubscriptionPlan(c *gin.Context) {
	userId := utils.GetUserID(c)
	if subscriptionUser, err := subscriptionPlanService.GetCurrentSubscriptionPlan(userId, 1); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.OkWithData(subscriptionUser, c)
	} else {
		response.OkWithData(subscriptionUser, c)
	}
}

// GetSubscriptionPlanByTag 用tag查询SubscriptionPlan
func (subscriptionPlanApi *SubscriptionPlanApi) GetSubscriptionPlanByTag(c *gin.Context) {
	var subscriptionPlan transaction.SubscriptionPlan
	err := c.ShouldBindQuery(&subscriptionPlan)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if subscriptionPlan.Tag == nil {
		response.FailWithMessage("tag为必填参数", c)
		return
	}
	if subscriptionPlan, err := subscriptionPlanService.GetSubscriptionPlanByTag(*subscriptionPlan.Tag); err != nil {
		global.Logger.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(subscriptionPlan, c)
	}
}

// GetSubscriptionPlanList 分页获取SubscriptionPlan列表
func (subscriptionPlanApi *SubscriptionPlanApi) GetSubscriptionPlanList(c *gin.Context) {
	var pageInfo transactionRequest.SubscriptionPlanSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := subscriptionPlanService.GetSubscriptionPlanInfoList(pageInfo); err != nil {
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
