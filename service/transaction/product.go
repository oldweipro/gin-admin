package transaction

import (
	"github.com/oldweipro/gin-admin/global"
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/transaction"
	openfishReq "github.com/oldweipro/gin-admin/model/transaction/request"
)

type ProductService struct {
}

// CreateProduct 创建Product记录
func (productService *ProductService) CreateProduct(product *transaction.Product) (err error) {
	err = global.DB.Create(product).Error
	return err
}

// DeleteProduct 删除Product记录
func (productService *ProductService) DeleteProduct(product transaction.Product) (err error) {
	err = global.DB.Delete(&product).Error
	return err
}

// DeleteProductByIds 批量删除Product记录
func (productService *ProductService) DeleteProductByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]transaction.Product{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateProduct 更新Product记录
func (productService *ProductService) UpdateProduct(product transaction.Product) (err error) {
	err = global.DB.Save(&product).Error
	return err
}

// GetProduct 根据id获取Product记录
func (productService *ProductService) GetProduct(id uint) (product transaction.Product, err error) {
	err = global.DB.Where("id = ?", id).First(&product).Error
	return
}

// GetProductInfoList 分页获取Product记录
func (productService *ProductService) GetProductInfoList(info openfishReq.ProductSearch) (list []transaction.Product, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&transaction.Product{})
	var products []transaction.Product
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
