package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/api/v1"
	"github.com/oldweipro/gin-admin/middleware"
)

type OpenApiRouter struct {
}

// InitPlatformRouter InitConversationRouter 初始化 Conversation 路由信息
func (p *OpenApiRouter) InitPlatformRouter(Router *gin.RouterGroup) {
	platformRouter := Router.Group("v1").Use(middleware.OperationRecord())
	platformRouterWithoutOperationRecord := Router.Group("v1")
	var platformApi = v1.ApiGroupApp.PlatformApiGroup.OpenApi
	{
		platformRouter.POST("chat/completions", platformApi.ForwardChatCompletionsApi) // 获取当前用户聊天列表
	}
	{
		platformRouterWithoutOperationRecord.OPTIONS("chat/completions", platformApi.ForwardOptionsChatCompletionsApi) // 获取当前用户聊天列表
	}
}
