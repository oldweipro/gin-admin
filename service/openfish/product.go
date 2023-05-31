package openfish

import (
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/openfish"
	openfishReq "github.com/oldweipro/gin-admin/model/openfish/request"
)

type ProductService struct {
}

// CreateProduct 创建Product记录
// Author [piexlmax](https://github.com/piexlmax)
func (productService *ProductService) CreateProduct(product *openfish.Product) (err error) {
	err = global.GVA_DB.Create(product).Error
	return err
}

// DeleteProduct 删除Product记录
// Author [piexlmax](https://github.com/piexlmax)
func (productService *ProductService) DeleteProduct(product openfish.Product) (err error) {
	err = global.GVA_DB.Delete(&product).Error
	return err
}

// DeleteProductByIds 批量删除Product记录
// Author [piexlmax](https://github.com/piexlmax)
func (productService *ProductService) DeleteProductByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]openfish.Product{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateProduct 更新Product记录
// Author [piexlmax](https://github.com/piexlmax)
func (productService *ProductService) UpdateProduct(product openfish.Product) (err error) {
	err = global.GVA_DB.Save(&product).Error
	return err
}

// GetProduct 根据id获取Product记录
// Author [piexlmax](https://github.com/piexlmax)
func (productService *ProductService) GetProduct(id uint) (product openfish.Product, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&product).Error
	return
}

// GetProductInfoList 分页获取Product记录
// Author [piexlmax](https://github.com/piexlmax)
func (productService *ProductService) GetProductInfoList(info openfishReq.ProductSearch) (list []openfish.Product, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&openfish.Product{})
	var products []openfish.Product
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Description != "" {
		db = db.Where("description LIKE ?", "%"+info.Description+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&products).Error
	return products, total, err
}
