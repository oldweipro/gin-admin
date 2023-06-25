package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/api/v1"
)

type OpenApiRouter struct {
}

// InitPlatformRouter InitConversationRouter 初始化 Conversation 路由信息
func (p *OpenApiRouter) InitPlatformRouter(Router *gin.RouterGroup) {
	platformRouter := Router.Group("v1")
	var platformApi = v1.ApiGroupApp.PlatformApiGroup.OpenApi
	{
		platformRouter.POST("chat/completions", platformApi.ForwardChatCompletionsApi) // 获取当前用户聊天列表
	}
}
