package memo_nexus

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/oldweipro/gin-admin/api/v1"
	"github.com/oldweipro/gin-admin/middleware"
)

type MemoNexusRouter struct {
}

func (s *MemoNexusRouter) InitMemoNexusRouter(Router *gin.RouterGroup) {
	memoNexusRouter := Router.Group("memoNexus").Use(middleware.OperationRecord())
	memoNexusRouterWithoutRecord := Router.Group("memoNexus")
	var chatTicketApi = v1.ApiGroupApp.MemoNexusApiGroup.MemoNexusApi
	{
		memoNexusRouter.GET("getLoginQrcodeGenerate", chatTicketApi.GetLoginQrcodeGenerate) // 每日签到
		memoNexusRouter.POST("loginQrcodePoll", chatTicketApi.LoginQrcodePoll)              // 每日签到
	}
	{
		memoNexusRouterWithoutRecord.GET("getChatTicketList", chatTicketApi.LoginQrcodePoll) // 获取ChatTicket列表
	}
}
