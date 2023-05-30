// 自动生成模板Wallets
package openfish

import (
	"github.com/oldweipro/gin-admin/global"
)

// Wallets 结构体
type Wallets struct {
	global.GVA_MODEL
	UserId     *int   `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;"`
	WalletName string `json:"walletName" form:"walletName" gorm:"column:wallet_name;comment:钱包名称;"`
	Balance    *int   `json:"balance" form:"balance" gorm:"column:balance;comment:余额;"`
}

// TableName Wallets 表名
func (Wallets) TableName() string {
	return "wallets"
}
