package openfish

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/api/v1"
	"github.com/oldweipro/gin-admin/middleware"
)

type SecretKeyRouter struct {
}

// InitSecretKeyRouter 初始化 SecretKey 路由信息
func (s *SecretKeyRouter) InitSecretKeyRouter(Router *gin.RouterGroup) {
	secretKeyRouter := Router.Group("secretKey").Use(middleware.OperationRecord())
	secretKeyRouterWithoutRecord := Router.Group("secretKey")
	var secretKeyApi = v1.ApiGroupApp.OpenfishApiGroup.SecretKeyApi
	{
		secretKeyRouter.POST("createSecretKey", secretKeyApi.CreateSecretKey)             // 新建SecretKey
		secretKeyRouter.DELETE("deleteSecretKey", secretKeyApi.DeleteSecretKey)           // 删除SecretKey
		secretKeyRouter.DELETE("deleteSecretKeyByIds", secretKeyApi.DeleteSecretKeyByIds) // 批量删除SecretKey
		secretKeyRouter.PUT("updateSecretKey", secretKeyApi.UpdateSecretKey)              // 更新SecretKey
	}
	{
		secretKeyRouterWithoutRecord.GET("findSecretKey", secretKeyApi.FindSecretKey)       // 根据ID获取SecretKey
		secretKeyRouterWithoutRecord.GET("getSecretKeyList", secretKeyApi.GetSecretKeyList) // 获取SecretKey列表
	}
}
