package transaction

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/api/v1"
)

type ActivationCodeRouter struct {
}

// InitActivationCodeRouter 初始化 激活码 路由信息
func (r *ActivationCodeRouter) InitActivationCodeRouter(Router *gin.RouterGroup) {
	activationCodeRouterWithoutRecord := Router.Group("activationCode")
	var activationCodeApi = v1.ApiGroupApp.TransactionApiGroup.ActivationCodeApi
	{
		activationCodeRouterWithoutRecord.GET("getActivationCodeStatus", activationCodeApi.GetActivationCodeStatus)       // 获取订阅状态
		activationCodeRouterWithoutRecord.GET("getJetBrainsActivationCode", activationCodeApi.GetJetBrainsActivationCode) // 获取激活码
	}
}
