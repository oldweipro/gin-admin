package patrol

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/api/v1"
	"github.com/oldweipro/gin-admin/middleware"
)

type IdentityRouter struct {
}

// InitIdentityRouter 初始化 Identity 路由信息
func (s *IdentityRouter) InitIdentityRouter(Router *gin.RouterGroup) {
	identityRouter := Router.Group("identity").Use(middleware.OperationRecord())
	identityRouterWithoutRecord := Router.Group("identity")
	var identityApi = v1.ApiGroupApp.PatrolApiGroup.IdentityApi
	{
		identityRouter.POST("createIdentity", identityApi.CreateIdentity)             // 新建Identity
		identityRouter.DELETE("deleteIdentity", identityApi.DeleteIdentity)           // 删除Identity
		identityRouter.DELETE("deleteIdentityByIds", identityApi.DeleteIdentityByIds) // 批量删除Identity
		identityRouter.PUT("updateIdentity", identityApi.UpdateIdentity)              // 更新Identity
	}
	{
		identityRouterWithoutRecord.GET("findIdentity", identityApi.FindIdentity)       // 根据ID获取Identity
		identityRouterWithoutRecord.GET("getIdentityList", identityApi.GetIdentityList) // 获取Identity列表
	}
}
