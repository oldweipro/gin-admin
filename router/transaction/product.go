package transaction

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/api/v1"
	"github.com/oldweipro/gin-admin/middleware"
)

type ProductRouter struct {
}

// InitProductRouter 初始化 Product 路由信息
func (s *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	productRouter := Router.Group("product").Use(middleware.OperationRecord())
	productRouterWithoutRecord := Router.Group("product")
	var productApi = v1.ApiGroupApp.TransactionApiGroup.ProductApi
	{
		productRouter.POST("createProduct", productApi.CreateProduct)             // 新建Product
		productRouter.DELETE("deleteProduct", productApi.DeleteProduct)           // 删除Product
		productRouter.DELETE("deleteProductByIds", productApi.DeleteProductByIds) // 批量删除Product
		productRouter.PUT("updateProduct", productApi.UpdateProduct)              // 更新Product
	}
	{
		productRouterWithoutRecord.GET("findProduct", productApi.FindProduct)       // 根据ID获取Product
		productRouterWithoutRecord.GET("getProductList", productApi.GetProductList) // 获取Product列表
	}
}
