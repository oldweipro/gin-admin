package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/oldweipro/gin-admin/api/v1"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	certificationRecordApi := v1.ApiGroupApp.PatrolApiGroup.CertificationRecordApi
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("emailLogin", baseApi.EmailLogin)
		baseRouter.POST("captcha", baseApi.Captcha)
		baseRouter.POST("registerWithSmsCode", baseApi.RegisterWithSmsCode) // 用户使用短信验证码注册
	}
	{
		baseRouter.POST("certification", certificationRecordApi.CreateCertificationRecord)
	}
	{
		baseRouter.POST("openFishLogin", baseApi.OpenFishLogin)
		baseRouter.POST("smsCode", baseApi.SmsCode)
	}
	return baseRouter
}
