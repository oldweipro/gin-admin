package ladder

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/api/v1"
	"github.com/oldweipro/gin-admin/middleware"
)

type InboundsRouter struct {
}

// InitInboundsRouter 初始化 Inbounds 路由信息
func (s *InboundsRouter) InitInboundsRouter(Router *gin.RouterGroup) {
	inboundsRouter := Router.Group("inbounds").Use(middleware.OperationRecord())
	inboundsRouterWithoutRecord := Router.Group("inbounds")
	var inboundsApi = v1.ApiGroupApp.LadderApiGroup.InboundsApi
	{
		inboundsRouter.POST("createInbounds", inboundsApi.CreateInbounds)             // 新建Inbounds
		inboundsRouter.POST("setInboundsLink", inboundsApi.SetInboundsLink)           // 重置节点链接
		inboundsRouter.DELETE("deleteInbounds", inboundsApi.DeleteInbounds)           // 删除Inbounds
		inboundsRouter.DELETE("deleteInboundsByIds", inboundsApi.DeleteInboundsByIds) // 批量删除Inbounds
		inboundsRouter.PUT("updateInbounds", inboundsApi.UpdateInbounds)              // 更新Inbounds
	}
	{
		inboundsRouterWithoutRecord.GET("findInbounds", inboundsApi.FindInbounds)         // 根据ID获取Inbounds
		inboundsRouterWithoutRecord.GET("getInboundsList", inboundsApi.GetInboundsList)   // 获取Inbounds列表
		inboundsRouterWithoutRecord.GET("findInboundsLink", inboundsApi.FindInboundsLink) // 根据服务器和当前用户查询节点信息
	}
}
