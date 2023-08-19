package transaction

import (
	"github.com/oldweipro/gin-admin/global"
	"time"
)

// RedeemCode 兑换码表
type RedeemCode struct {
	global.Model
	Code          string    `json:"code" form:"code" gorm:"column:redeem_code;comment:兑换码;type:varchar(20)"`
	Amount        *uint     `json:"amount" form:"amount" gorm:"column:amount;comment:面额;type:int"`
	TotalCount    string    `json:"totalCount" form:"totalCount" gorm:"column:total_count;comment:总数量,该批兑换码的总数量;type:int"`
	LeftCount     *uint     `json:"leftCount" form:"leftCount" gorm:"column:left_count;comment:剩余数量,该批兑换码的剩余可用数量;type:int"`
	PerLimit      *uint     `json:"perLimit" form:"perLimit" gorm:"column:per_limit;comment:每人限兑数量,每人最多可兑换次数;type:int"`
	TotalRedeemed *uint     `json:"totalRedeemed" form:"totalRedeemed" gorm:"column:total_redeemed;comment:总兑换次数,该批码已被兑换的总次数;type:int"`
	Status        string    `json:"status" form:"status" gorm:"column:status;comment:状态,兑换码状态(未使用、已使用、已过期等);type:varchar(20)"`
	ExpireTime    time.Time `json:"expireTime" form:"expireTime" gorm:"column:expire_time;comment:过期时间;type:datetime"`
}

// TableName RedeemCode 表名
func (RedeemCode) TableName() string {
	return "product"
}

// RedeemLog 兑换记录表
type RedeemLog struct {
	global.Model
	UserID       uint      `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID,兑换用户;type:int"`
	RedeemCodeID uint      `json:"redeemCodeId" form:"redeemCodeId" gorm:"column:redeem_code_id;comment:兑换码ID,外键关联兑换码表;type:int"`
	RedeemTime   time.Time `json:"redeemTime" form:"redeemTime" gorm:"column:redeem_time;comment:兑换时间;type:datetime"`
}

// TableName RedeemLog 表名
func (RedeemLog) TableName() string {
	return "product"
}
