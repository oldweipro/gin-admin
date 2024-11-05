package system

import (
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("captcha", baseApi.Captcha)
		baseRouter.POST("registerOrResetPassword", baseApi.RegisterOrResetPassword) // 注册账号或找回密码
		baseRouter.POST("sendSmsCode", baseApi.SendSmsCode)                         // 获取短信验证码
	}
	return baseRouter
}
