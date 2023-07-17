package openfish

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/api/v1"
	"github.com/oldweipro/gin-admin/middleware"
)

type PromptRouter struct {
}

// InitPromptRouter 初始化 Prompt 路由信息
func (s *PromptRouter) InitPromptRouter(Router *gin.RouterGroup) {
	promptRouter := Router.Group("prompt").Use(middleware.OperationRecord())
	promptRouterWithoutRecord := Router.Group("prompt")
	var promptApi = v1.ApiGroupApp.OpenfishApiGroup.PromptApi
	{
		promptRouter.POST("createPrompt", promptApi.CreatePrompt)             // 新建Prompt
		promptRouter.DELETE("deletePrompt", promptApi.DeletePrompt)           // 删除Prompt
		promptRouter.DELETE("deletePromptByIds", promptApi.DeletePromptByIds) // 批量删除Prompt
		promptRouter.PUT("updatePrompt", promptApi.UpdatePrompt)              // 更新Prompt
	}
	{
		promptRouterWithoutRecord.GET("findPrompt", promptApi.FindPrompt)                             // 根据ID获取Prompt
		promptRouterWithoutRecord.GET("getPromptList", promptApi.GetPromptList)                       // 获取Prompt列表
		promptRouterWithoutRecord.GET("getCurrentUserPromptList", promptApi.GetCurrentUserPromptList) // 获取当前用户Prompt列表
	}
}
