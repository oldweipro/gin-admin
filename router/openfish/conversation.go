package openfish

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/api/v1"
	"github.com/oldweipro/gin-admin/middleware"
)

type ConversationRouter struct {
}

// InitConversationRouter 初始化 Conversation 路由信息
func (s *ConversationRouter) InitConversationRouter(Router *gin.RouterGroup) {
	conversationRouter := Router.Group("conversation").Use(middleware.OperationRecord())
	conversationRouterWithoutRecord := Router.Group("conversation")
	var conversationApi = v1.ApiGroupApp.OpenfishApiGroup.ConversationApi
	{
		conversationRouter.POST("createConversation", conversationApi.CreateConversation)             // 新建Conversation
		conversationRouter.POST("chatCompletions", conversationApi.ChatCompletions)                   // AI对话
		conversationRouter.POST("deleteConversation", conversationApi.DeleteConversation)             // 删除Conversation
		conversationRouter.DELETE("deleteConversationByIds", conversationApi.DeleteConversationByIds) // 批量删除Conversation
		conversationRouter.PUT("updateConversation", conversationApi.UpdateConversation)              // 更新Conversation
		conversationRouter.POST("updateConversationName", conversationApi.UpdateConversationName)     // 更新Conversation
	}
	{
		conversationRouterWithoutRecord.GET("findConversation", conversationApi.FindConversation)                             // 根据ID获取Conversation
		conversationRouterWithoutRecord.GET("getConversationList", conversationApi.GetConversationList)                       // 获取Conversation列表
		conversationRouterWithoutRecord.GET("getCurrentUserConversationList", conversationApi.GetCurrentUserConversationList) // 获取当前用户聊天列表
	}
}
