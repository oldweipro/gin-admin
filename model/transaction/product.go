// 自动生成模板Product
package transaction

import (
	"github.com/oldweipro/gin-admin/global"
)

// Product 结构体
type Product struct {
	global.GVA_MODEL
	Name        string `json:"name" form:"name" gorm:"column:name;comment:表示商品的名称。;"`
	Description string `json:"description" form:"description" gorm:"column:description;comment:表示商品的详细描述信息。;"`
	Price       *int   `json:"price" form:"price" gorm:"column:price;comment:表示商品的价格。;"`
}

// TableName Product 表名
func (Product) TableName() string {
	return "product"
}
