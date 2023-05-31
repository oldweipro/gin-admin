package openfish

import (
	"github.com/gin-gonic/gin"
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/common/response"
	"github.com/oldweipro/gin-admin/model/openfish"
	openfishReq "github.com/oldweipro/gin-admin/model/openfish/request"
	"github.com/oldweipro/gin-admin/service"
	"github.com/oldweipro/gin-admin/utils"
	"go.uber.org/zap"
)

type ProductApi struct {
}

var productService = service.ServiceGroupApp.OpenfishServiceGroup.ProductService

// CreateProduct 创建Product
// @Tags Product
// @Summary 创建Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body openfish.Product true "创建Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /product/createProduct [post]
func (productApi *ProductApi) CreateProduct(c *gin.Context) {
	var product openfish.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":  {utils.NotEmpty()},
		"Price": {utils.NotEmpty()},
	}
	if err := utils.Verify(product, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := productService.CreateProduct(&product); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteProduct 删除Product
// @Tags Product
// @Summary 删除Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body openfish.Product true "删除Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /product/deleteProduct [delete]
func (productApi *ProductApi) DeleteProduct(c *gin.Context) {
	var product openfish.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := productService.DeleteProduct(product); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteProductByIds 批量删除Product
// @Tags Product
// @Summary 批量删除Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /product/deleteProductByIds [delete]
func (productApi *ProductApi) DeleteProductByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := productService.DeleteProductByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateProduct 更新Product
// @Tags Product
// @Summary 更新Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body openfish.Product true "更新Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /product/updateProduct [put]
func (productApi *ProductApi) UpdateProduct(c *gin.Context) {
	var product openfish.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":  {utils.NotEmpty()},
		"Price": {utils.NotEmpty()},
	}
	if err := utils.Verify(product, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := productService.UpdateProduct(product); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindProduct 用id查询Product
// @Tags Product
// @Summary 用id查询Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query openfish.Product true "用id查询Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /product/findProduct [get]
func (productApi *ProductApi) FindProduct(c *gin.Context) {
	var product openfish.Product
	err := c.ShouldBindQuery(&product)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reproduct, err := productService.GetProduct(product.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reproduct": reproduct}, c)
	}
}

// GetProductList 分页获取Product列表
// @Tags Product
// @Summary 分页获取Product列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query openfishReq.ProductSearch true "分页获取Product列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /product/getProductList [get]
func (productApi *ProductApi) GetProductList(c *gin.Context) {
	var pageInfo openfishReq.ProductSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := productService.GetProductInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
