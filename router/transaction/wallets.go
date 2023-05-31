package transaction

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/api/v1"
	"github.com/oldweipro/gin-admin/middleware"
)

type WalletsRouter struct {
}

// InitWalletsRouter 初始化 Wallets 路由信息
func (s *WalletsRouter) InitWalletsRouter(Router *gin.RouterGroup) {
	walletsRouter := Router.Group("wallets").Use(middleware.OperationRecord())
	walletsRouterWithoutRecord := Router.Group("wallets")
	var walletsApi = v1.ApiGroupApp.TransactionApiGroup.WalletsApi
	{
		walletsRouter.POST("createWallets", walletsApi.CreateWallets)             // 新建Wallets
		walletsRouter.DELETE("deleteWallets", walletsApi.DeleteWallets)           // 删除Wallets
		walletsRouter.DELETE("deleteWalletsByIds", walletsApi.DeleteWalletsByIds) // 批量删除Wallets
		walletsRouter.PUT("updateWallets", walletsApi.UpdateWallets)              // 更新Wallets
	}
	{
		walletsRouterWithoutRecord.GET("findWallets", walletsApi.FindWallets)       // 根据ID获取Wallets
		walletsRouterWithoutRecord.GET("getWalletsList", walletsApi.GetWalletsList) // 获取Wallets列表
	}
}
