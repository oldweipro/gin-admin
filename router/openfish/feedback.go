package openfish

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/api/v1"
	"github.com/oldweipro/gin-admin/middleware"
)

type FeedbackRouter struct {
}

// InitFeedbackRouter 初始化 Feedback 路由信息
func (s *FeedbackRouter) InitFeedbackRouter(Router *gin.RouterGroup) {
	feedbackRouter := Router.Group("feedback").Use(middleware.OperationRecord())
	feedbackRouterWithoutRecord := Router.Group("feedback")
	var feedbackApi = v1.ApiGroupApp.OpenfishApiGroup.FeedbackApi
	{
		feedbackRouter.POST("createFeedback", feedbackApi.CreateFeedback)             // 新建Feedback
		feedbackRouter.DELETE("deleteFeedback", feedbackApi.DeleteFeedback)           // 删除Feedback
		feedbackRouter.DELETE("deleteFeedbackByIds", feedbackApi.DeleteFeedbackByIds) // 批量删除Feedback
		feedbackRouter.PUT("updateFeedback", feedbackApi.UpdateFeedback)              // 更新Feedback
	}
	{
		feedbackRouterWithoutRecord.GET("findFeedback", feedbackApi.FindFeedback)       // 根据ID获取Feedback
		feedbackRouterWithoutRecord.GET("getFeedbackList", feedbackApi.GetFeedbackList) // 获取Feedback列表
	}
}
