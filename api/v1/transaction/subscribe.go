package transaction

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/response"
	"github.com/oldweipro/gin-admin/model/transaction/request"
	"github.com/oldweipro/gin-admin/service"
	"github.com/oldweipro/gin-admin/utils"
	"go.uber.org/zap"
	"sync"
)

type SubscribeApi struct {
}

var subscribeService = service.ServiceGroupApp.TransactionServiceGroup.SubscribeService

var subscribePlanStatus sync.Map

// SubscribePlan 订阅计划 给某人。某个功能模块。按照时长，单价，扣费，记录，
func (subscribeApi *SubscribeApi) SubscribePlan(c *gin.Context) {
	userId := utils.GetUserID(c)
	_, loaded := subscribePlanStatus.LoadOrStore(userId, true)
	if loaded {
		response.FailStatusTooManyRequestsWithDetailed(nil, "请求过多", c)
		return
	}
	defer subscribePlanStatus.Delete(userId)
	var subscribe request.SubscribeRequest
	err := c.ShouldBindJSON(&subscribe)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if subscribe.PlanId == nil {
		response.FailWithMessage("必要参数为空", c)
		return
	}
	// 查询订阅计划，根据订阅计划为用户续费
	if plan, err := subscriptionPlanService.GetSubscriptionPlan(*subscribe.PlanId); err != nil {
		global.Logger.Error("订阅计划不存在!", zap.Error(err))
		response.FailWithMessage("订阅计划不存在", c)
	} else {
		// 查询自己的订阅计划
		if userPlan, err := subscriptionPlanService.GetCurrentSubscriptionPlan(utils.GetUserID(c)); err != nil {
			global.Logger.Error("查询自己的订阅计划失败!", zap.Error(err))
			response.FailWithMessage("查询自己的订阅计划失败!", c)
			return
		} else {
			// 续费订阅计划
			err := subscribeService.RenewalSubscription(&userPlan, &plan)
			if err != nil {
				response.FailWithMessage(err.Error(), c)
				return
			}
		}
		response.OkWithMessage("订阅成功", c)
	}
}
