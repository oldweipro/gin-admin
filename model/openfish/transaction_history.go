// 自动生成模板TransactionHistory
package openfish

import (
	"github.com/oldweipro/gin-admin/global"
)

// TransactionHistory 结构体
type TransactionHistory struct {
	global.GVA_MODEL
	UserId       *int   `json:"userId" form:"userId" gorm:"column:user_id;comment:与用户表中的UserID字段相关联，表示进行交易的用户;"`
	SrcWalletId  *int   `json:"srcWalletId" form:"srcWalletId" gorm:"column:src_wallet_id;comment:与钱包表中的WalletID字段相关联，表示交易源钱包。;"`
	DestWalletId *int   `json:"destWalletId" form:"destWalletId" gorm:"column:dest_wallet_id;comment:与钱包表中的WalletID字段相关联，表示交易目标钱包。;"`
	TypeEnum     string `json:"typeEnum" form:"typeEnum" gorm:"column:type_enum;type:enum('deposit', 'withdrawal', 'transfer');comment:交易类型，枚举值，表示交易的类型，例如存款、取款、转账等。;"`
	Quantity     *int   `json:"quantity" form:"quantity" gorm:"default:0;column:quantity;comment:表示交易中购买的商品数量。;"`
	Amount       *int   `json:"amount" form:"amount" gorm:"default:0;column:amount;comment:交易涉及的金额。;"`
	ProductId    *int   `json:"productId" form:"productId" gorm:"column:product_id;comment:与商品表中的ProductID字段相关联，表示交易涉及的商品。;"`
	CreatedBy    uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy    uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy    uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName TransactionHistory 表名
func (TransactionHistory) TableName() string {
	return "transaction_history"
}
