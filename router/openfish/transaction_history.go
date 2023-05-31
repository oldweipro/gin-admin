package openfish

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/api/v1"
	"github.com/oldweipro/gin-admin/middleware"
)

type TransactionHistoryRouter struct {
}

// InitTransactionHistoryRouter 初始化 TransactionHistory 路由信息
func (s *TransactionHistoryRouter) InitTransactionHistoryRouter(Router *gin.RouterGroup) {
	transactionHistoryRouter := Router.Group("transactionHistory").Use(middleware.OperationRecord())
	transactionHistoryRouterWithoutRecord := Router.Group("transactionHistory")
	var transactionHistoryApi = v1.ApiGroupApp.OpenfishApiGroup.TransactionHistoryApi
	{
		transactionHistoryRouter.POST("createTransactionHistory", transactionHistoryApi.CreateTransactionHistory)             // 新建TransactionHistory
		transactionHistoryRouter.DELETE("deleteTransactionHistory", transactionHistoryApi.DeleteTransactionHistory)           // 删除TransactionHistory
		transactionHistoryRouter.DELETE("deleteTransactionHistoryByIds", transactionHistoryApi.DeleteTransactionHistoryByIds) // 批量删除TransactionHistory
		transactionHistoryRouter.PUT("updateTransactionHistory", transactionHistoryApi.UpdateTransactionHistory)              // 更新TransactionHistory
	}
	{
		transactionHistoryRouterWithoutRecord.GET("findTransactionHistory", transactionHistoryApi.FindTransactionHistory)       // 根据ID获取TransactionHistory
		transactionHistoryRouterWithoutRecord.GET("getTransactionHistoryList", transactionHistoryApi.GetTransactionHistoryList) // 获取TransactionHistory列表
	}
}
