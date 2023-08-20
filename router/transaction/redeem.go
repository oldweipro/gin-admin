package transaction

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/api/v1"
	"github.com/oldweipro/gin-admin/middleware"
)

type RedeemRouter struct {
}

// InitRedeemRouter 初始化 兑换码 路由信息
func (r *RedeemRouter) InitRedeemRouter(Router *gin.RouterGroup) {
	redeemRouter := Router.Group("redeem").Use(middleware.OperationRecord())
	redeemRouterWithoutRecord := Router.Group("redeem")
	var redeemApi = v1.ApiGroupApp.TransactionApiGroup.RedeemApi
	{
		redeemRouter.POST("generateRedeemCode", redeemApi.GenerateRedeemCode) // 生成兑换码
	}
	{
		redeemRouterWithoutRecord.GET("getRedeemCodeList", redeemApi.GenerateRedeemCode) // 获取兑换码列表
	}
}
