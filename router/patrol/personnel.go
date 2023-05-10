package patrol

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/api/v1"
	"github.com/oldweipro/gin-admin/middleware"
)

type PersonnelRouter struct {
}

// InitPersonnelRouter 初始化 Personnel 路由信息
func (s *PersonnelRouter) InitPersonnelRouter(Router *gin.RouterGroup) {
	personnelRouter := Router.Group("personnel").Use(middleware.OperationRecord())
	personnelRouterWithoutRecord := Router.Group("personnel")
	var personnelApi = v1.ApiGroupApp.PatrolApiGroup.PersonnelApi
	{
		personnelRouter.POST("syncPersonnel", personnelApi.SyncPersonnel)                 // 同步人员信息
		personnelRouter.POST("syncPersonnelImg", personnelApi.SyncPersonnelImg)           // 同步人员图片信息
		personnelRouter.POST("createPersonnel", personnelApi.CreatePersonnel)             // 新建Personnel
		personnelRouter.DELETE("deletePersonnel", personnelApi.DeletePersonnel)           // 删除Personnel
		personnelRouter.DELETE("deletePersonnelByIds", personnelApi.DeletePersonnelByIds) // 批量删除Personnel
		personnelRouter.PUT("updatePersonnel", personnelApi.UpdatePersonnel)              // 更新Personnel
	}
	{
		personnelRouterWithoutRecord.GET("findPersonnel", personnelApi.FindPersonnel)       // 根据ID获取Personnel
		personnelRouterWithoutRecord.GET("getPersonnelList", personnelApi.GetPersonnelList) // 获取Personnel列表
	}
}
