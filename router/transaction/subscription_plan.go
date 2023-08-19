package transaction

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/api/v1"
	"github.com/oldweipro/gin-admin/middleware"
)

type SubscriptionPlanRouter struct {
}

// InitSubscriptionPlanRouter 初始化 SubscriptionPlan 路由信息
func (s *SubscriptionPlanRouter) InitSubscriptionPlanRouter(Router *gin.RouterGroup) {
	subscriptionPlanRouter := Router.Group("subscriptionPlan").Use(middleware.OperationRecord())
	subscriptionPlanRouterWithoutRecord := Router.Group("subscriptionPlan")
	var subscriptionPlanApi = v1.ApiGroupApp.TransactionApiGroup.SubscriptionPlanApi
	var subscribeApi = v1.ApiGroupApp.TransactionApiGroup.SubscribeApi
	{
		subscriptionPlanRouter.POST("createSubscriptionPlan", subscriptionPlanApi.CreateSubscriptionPlan)             // 新建SubscriptionPlan
		subscriptionPlanRouter.DELETE("deleteSubscriptionPlan", subscriptionPlanApi.DeleteSubscriptionPlan)           // 删除SubscriptionPlan
		subscriptionPlanRouter.DELETE("deleteSubscriptionPlanByIds", subscriptionPlanApi.DeleteSubscriptionPlanByIds) // 批量删除SubscriptionPlan
		subscriptionPlanRouter.PUT("updateSubscriptionPlan", subscriptionPlanApi.UpdateSubscriptionPlan)              // 更新SubscriptionPlan
		subscriptionPlanRouter.GET("getCurrentSubscriptionPlan", subscriptionPlanApi.GetCurrentSubscriptionPlan)      // 查询当前用户订阅计划
		subscriptionPlanRouter.GET("getSubscriptionPlanByTag", subscriptionPlanApi.GetSubscriptionPlanByTag)          // 查询某个功能的订阅计划
		subscriptionPlanRouter.POST("subscribePlan", subscribeApi.SubscribePlan)                                      // 订阅
	}
	{
		subscriptionPlanRouterWithoutRecord.GET("getSubscriptionPlan", subscriptionPlanApi.GetSubscriptionPlan)         // 根据ID获取SubscriptionPlan
		subscriptionPlanRouterWithoutRecord.GET("getSubscriptionPlanList", subscriptionPlanApi.GetSubscriptionPlanList) // 获取SubscriptionPlan列表
	}
}
