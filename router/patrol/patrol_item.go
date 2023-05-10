package patrol

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/api/v1"
	"github.com/oldweipro/gin-admin/middleware"
)

type PatrolItemRouter struct {
}

// InitPatrolItemRouter 初始化 PatrolItem 路由信息
func (s *PatrolItemRouter) InitPatrolItemRouter(Router *gin.RouterGroup) {
	patrolItemRouter := Router.Group("patrolItem").Use(middleware.OperationRecord())
	patrolItemRouterWithoutRecord := Router.Group("patrolItem")
	var patrolItemApi = v1.ApiGroupApp.PatrolApiGroup.PatrolItemApi
	{
		patrolItemRouter.POST("createPatrolItem", patrolItemApi.CreatePatrolItem)             // 新建PatrolItem
		patrolItemRouter.DELETE("deletePatrolItem", patrolItemApi.DeletePatrolItem)           // 删除PatrolItem
		patrolItemRouter.DELETE("deletePatrolItemByIds", patrolItemApi.DeletePatrolItemByIds) // 批量删除PatrolItem
		patrolItemRouter.PUT("updatePatrolItem", patrolItemApi.UpdatePatrolItem)              // 更新PatrolItem
	}
	{
		patrolItemRouterWithoutRecord.GET("findPatrolItem", patrolItemApi.FindPatrolItem)       // 根据ID获取PatrolItem
		patrolItemRouterWithoutRecord.GET("getPatrolItemList", patrolItemApi.GetPatrolItemList) // 获取PatrolItem列表
	}
}
