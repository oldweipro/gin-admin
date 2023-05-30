package openfish

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/api/v1"
	"github.com/oldweipro/gin-admin/middleware"
)

type ChatTicketRouter struct {
}

// InitChatTicketRouter 初始化 ChatTicket 路由信息
func (s *ChatTicketRouter) InitChatTicketRouter(Router *gin.RouterGroup) {
	chatTicketRouter := Router.Group("chatTicket").Use(middleware.OperationRecord())
	chatTicketRouterWithoutRecord := Router.Group("chatTicket")
	var chatTicketApi = v1.ApiGroupApp.OpenfishApiGroup.ChatTicketApi
	{
		chatTicketRouter.POST("createChatTicket", chatTicketApi.CreateChatTicket)             // 新建ChatTicket
		chatTicketRouter.DELETE("deleteChatTicket", chatTicketApi.DeleteChatTicket)           // 删除ChatTicket
		chatTicketRouter.DELETE("deleteChatTicketByIds", chatTicketApi.DeleteChatTicketByIds) // 批量删除ChatTicket
		chatTicketRouter.PUT("updateChatTicket", chatTicketApi.UpdateChatTicket)              // 更新ChatTicket
	}
	{
		chatTicketRouterWithoutRecord.GET("findChatTicket", chatTicketApi.FindChatTicket)       // 根据ID获取ChatTicket
		chatTicketRouterWithoutRecord.GET("getChatTicketList", chatTicketApi.GetChatTicketList) // 获取ChatTicket列表
	}
}
