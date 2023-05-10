package patrol

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/api/v1"
	"github.com/oldweipro/gin-admin/middleware"
)

type PatrolTaskRouter struct {
}

// InitPatrolTaskRouter 初始化 PatrolTask 路由信息
func (s *PatrolTaskRouter) InitPatrolTaskRouter(Router *gin.RouterGroup) {
	patrolTaskRouter := Router.Group("patrolTask").Use(middleware.OperationRecord())
	patrolTaskRouterWithoutRecord := Router.Group("patrolTask")
	var patrolTaskApi = v1.ApiGroupApp.PatrolApiGroup.PatrolTaskApi
	{
		patrolTaskRouter.POST("createPatrolTask", patrolTaskApi.CreatePatrolTask)             // 新建PatrolTask
		patrolTaskRouter.DELETE("deletePatrolTask", patrolTaskApi.DeletePatrolTask)           // 删除PatrolTask
		patrolTaskRouter.DELETE("deletePatrolTaskByIds", patrolTaskApi.DeletePatrolTaskByIds) // 批量删除PatrolTask
		patrolTaskRouter.PUT("updatePatrolTask", patrolTaskApi.UpdatePatrolTask)              // 更新PatrolTask
	}
	{
		patrolTaskRouterWithoutRecord.GET("findPatrolTask", patrolTaskApi.FindPatrolTask)       // 根据ID获取PatrolTask
		patrolTaskRouterWithoutRecord.GET("getPatrolTaskList", patrolTaskApi.GetPatrolTaskList) // 获取PatrolTask列表
	}
}
