package openfish

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/api/v1"
	"github.com/oldweipro/gin-admin/middleware"
)

type MailAccountRouter struct {
}

// InitMailAccountRouter 初始化 MailAccount 路由信息
func (s *MailAccountRouter) InitMailAccountRouter(Router *gin.RouterGroup) {
	mailAccountRouter := Router.Group("mailAccount").Use(middleware.OperationRecord())
	mailAccountRouterWithoutRecord := Router.Group("mailAccount")
	var mailAccountApi = v1.ApiGroupApp.OpenfishApiGroup.MailAccountApi
	{
		mailAccountRouter.POST("createMailAccount", mailAccountApi.CreateMailAccount)               // 新建
		mailAccountRouter.DELETE("deleteMailAccount", mailAccountApi.DeleteMailAccount)             // 删除
		mailAccountRouter.DELETE("deleteMailAccountByIds", mailAccountApi.DeleteMailAccountByIds)   // 批量删除
		mailAccountRouter.PUT("updateMailAccount", mailAccountApi.UpdateMailAccount)                // 更新
		mailAccountRouter.POST("refreshClaudeChat", mailAccountApi.RefreshClaudeChat)               // 产生一次Claude对话
		mailAccountRouter.POST("refreshOpenaiAccessToken", mailAccountApi.RefreshOpenaiAccessToken) // 刷新 openai AccessToken
		mailAccountRouter.POST("syncOpenaiInfo", mailAccountApi.SyncOpenaiInfo)                     // 同步openai的信息: sk到期时间、余额
		mailAccountRouter.POST("syncChatGPTAccessToken", mailAccountApi.SyncChatGPTAccessToken)     // 同步openai的AT
	}
	{
		mailAccountRouterWithoutRecord.GET("getMailAccount", mailAccountApi.FindMailAccount)        // 根据ID获取
		mailAccountRouterWithoutRecord.GET("getMailAccountList", mailAccountApi.GetMailAccountList) // 获取分页列表
	}
}
