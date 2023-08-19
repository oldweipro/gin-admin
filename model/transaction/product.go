package transaction

import (
	"github.com/oldweipro/gin-admin/global"
)

// Product 结构体
type Product struct {
	global.Model
	Name        string `json:"name" form:"name" gorm:"column:name;comment:表示商品的名称。;"`
	Description string `json:"description" form:"description" gorm:"column:description;comment:表示商品的详细描述信息。;"`
	Price       *uint  `json:"price" form:"price" gorm:"column:price;comment:表示商品的价格。;"`
	Duration    *uint  `json:"duration" form:"duration" gorm:"column:duration;comment:表示商品的有效期。天数表示;"`
	Quantity    *uint  `json:"quantity" form:"quantity" gorm:"column:quantity;comment:表示商品的数量;"`
	MenuId      uint   `json:"menuId" form:"menuId" gorm:"column:menu_id;comment:关联的权限ID;"`
}

// TableName Product 表名
func (Product) TableName() string {
	return "product"
}
