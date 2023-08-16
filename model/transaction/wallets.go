// 自动生成模板Wallets
package transaction

import (
	"github.com/oldweipro/gin-admin/global"
)

// Wallets 结构体
type Wallets struct {
	global.Model
	UserId     uint   `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;"`
	WalletName string `json:"walletName" form:"walletName" gorm:"default:'钱包';column:wallet_name;comment:钱包名称;"`
	Balance    *uint  `json:"balance" form:"balance" gorm:"default:0;column:balance;comment:余额;"`
}

// TableName Wallets 表名
func (Wallets) TableName() string {
	return "wallets"
}
