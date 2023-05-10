package patrol

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/api/v1"
	"github.com/oldweipro/gin-admin/middleware"
)

type CertificationRecordRouter struct {
}

// InitCertificationRecordRouter 初始化 CertificationRecord 路由信息
func (s *CertificationRecordRouter) InitCertificationRecordRouter(Router *gin.RouterGroup) {
	certificationRecordRouter := Router.Group("certificationRecord").Use(middleware.OperationRecord())
	certificationRecordRouterWithoutRecord := Router.Group("certificationRecord")
	var certificationRecordApi = v1.ApiGroupApp.PatrolApiGroup.CertificationRecordApi
	{
		certificationRecordRouter.POST("createCertificationRecord", certificationRecordApi.CreateCertificationRecord)             // 新建CertificationRecord
		certificationRecordRouter.DELETE("deleteCertificationRecord", certificationRecordApi.DeleteCertificationRecord)           // 删除CertificationRecord
		certificationRecordRouter.DELETE("deleteCertificationRecordByIds", certificationRecordApi.DeleteCertificationRecordByIds) // 批量删除CertificationRecord
		certificationRecordRouter.PUT("updateCertificationRecord", certificationRecordApi.UpdateCertificationRecord)              // 更新CertificationRecord
	}
	{
		certificationRecordRouterWithoutRecord.GET("findCertificationRecord", certificationRecordApi.FindCertificationRecord)       // 根据ID获取CertificationRecord
		certificationRecordRouterWithoutRecord.GET("getCertificationRecordList", certificationRecordApi.GetCertificationRecordList) // 获取CertificationRecord列表
	}
}
