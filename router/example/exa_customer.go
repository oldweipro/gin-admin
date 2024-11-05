package example

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/pkg/middleware"
)

type CustomerRouter struct{}

func (e *CustomerRouter) InitCustomerRouter(Router *gin.RouterGroup) {
	customerRouter := Router.Group("customer").Use(middleware.OperationRecord())
	customerRouterWithoutRecord := Router.Group("customer")
	{
		customerRouter.POST("customer", exaCustomerApi.CreateExaCustomer)   // 创建客户
		customerRouter.PUT("customer", exaCustomerApi.UpdateExaCustomer)    // 更新客户
		customerRouter.DELETE("customer", exaCustomerApi.DeleteExaCustomer) // 删除客户
	}
	{
		customerRouterWithoutRecord.GET("customer", exaCustomerApi.GetExaCustomer)         // 获取单一客户信息
		customerRouterWithoutRecord.GET("customerList", exaCustomerApi.GetExaCustomerList) // 获取客户列表
	}
}
