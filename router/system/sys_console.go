package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/oldweipro/gin-admin/api/v1"
	"github.com/oldweipro/gin-admin/middleware"
)

type DashboardRouter struct{}

func (s *DashboardRouter) InitDashboardRouter(Router *gin.RouterGroup) {
	sysRouter := Router.Group("dashboard").Use(middleware.OperationRecord())
	systemApi := v1.ApiGroupApp.SystemApiGroup.DashboardApi
	{
		sysRouter.GET("console", systemApi.Console) // 获取配置文件内容
	}
}
