package ladder

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/api/v1"
	"github.com/oldweipro/gin-admin/middleware"
)

type ServerNodeRouter struct {
}

// InitServerNodeRouter 初始化 ServerNode 路由信息
func (s *ServerNodeRouter) InitServerNodeRouter(Router *gin.RouterGroup) {
	serverNodeRouter := Router.Group("serverNode").Use(middleware.OperationRecord())
	serverNodeRouterWithoutRecord := Router.Group("serverNode")
	var serverNodeApi = v1.ApiGroupApp.LadderApiGroup.ServerNodeApi
	{
		serverNodeRouter.POST("createServerNode", serverNodeApi.CreateServerNode)             // 新建ServerNode
		serverNodeRouter.DELETE("deleteServerNode", serverNodeApi.DeleteServerNode)           // 删除ServerNode
		serverNodeRouter.DELETE("deleteServerNodeByIds", serverNodeApi.DeleteServerNodeByIds) // 批量删除ServerNode
		serverNodeRouter.PUT("updateServerNode", serverNodeApi.UpdateServerNode)              // 更新ServerNode
	}
	{
		serverNodeRouterWithoutRecord.GET("findServerNode", serverNodeApi.FindServerNode)               // 根据ID获取ServerNode
		serverNodeRouterWithoutRecord.GET("getServerNodeList", serverNodeApi.GetServerNodeList)         // 获取ServerNode列表
		serverNodeRouterWithoutRecord.GET("getServerNodeLessList", serverNodeApi.GetServerNodeLessList) // 获取ServerNodeLess列表
	}
}
