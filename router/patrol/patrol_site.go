package patrol

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/api/v1"
	"github.com/oldweipro/gin-admin/middleware"
)

type PatrolSiteRouter struct {
}

// InitPatrolSiteRouter 初始化 PatrolSite 路由信息
func (s *PatrolSiteRouter) InitPatrolSiteRouter(Router *gin.RouterGroup) {
	patrolSiteRouter := Router.Group("patrolSite").Use(middleware.OperationRecord())
	patrolSiteRouterWithoutRecord := Router.Group("patrolSite")
	var patrolSiteApi = v1.ApiGroupApp.PatrolApiGroup.PatrolSiteApi
	{
		patrolSiteRouter.POST("createPatrolSite", patrolSiteApi.CreatePatrolSite)             // 新建PatrolSite
		patrolSiteRouter.DELETE("deletePatrolSite", patrolSiteApi.DeletePatrolSite)           // 删除PatrolSite
		patrolSiteRouter.DELETE("deletePatrolSiteByIds", patrolSiteApi.DeletePatrolSiteByIds) // 批量删除PatrolSite
		patrolSiteRouter.PUT("updatePatrolSite", patrolSiteApi.UpdatePatrolSite)              // 更新PatrolSite
	}
	{
		patrolSiteRouterWithoutRecord.GET("findPatrolSite", patrolSiteApi.FindPatrolSite)       // 根据ID获取PatrolSite
		patrolSiteRouterWithoutRecord.GET("getPatrolSiteList", patrolSiteApi.GetPatrolSiteList) // 获取PatrolSite列表
	}
}
