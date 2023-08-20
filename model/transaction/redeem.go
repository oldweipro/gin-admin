package transaction

import (
	"github.com/oldweipro/gin-admin/global"
	"time"
)

// RedeemCode 兑换码表
type RedeemCode struct {
	global.Model
	Code          string    `json:"code" form:"code" gorm:"column:code;comment:兑换码;"`
	Amount        *uint     `json:"amount" form:"amount" gorm:"column:amount;comment:面额;"`
	TotalCount    *uint     `json:"totalCount" form:"totalCount" gorm:"column:total_count;comment:总数量,该兑换码的总数量;"`
	LeftCount     *uint     `json:"leftCount" form:"leftCount" gorm:"column:left_count;comment:剩余数量,该兑换码的剩余可用数量;"`
	PerLimit      *uint     `json:"perLimit" form:"perLimit" gorm:"column:per_limit;comment:每人限兑数量,每人最多可兑换次数;"`
	ExpireTime    time.Time `json:"expireTime" form:"expireTime" gorm:"column:expire_time;comment:过期时间;"`
	Status        *uint     `json:"status" form:"status" gorm:"column:status;comment:状态,兑换码状态(0未使用、1已使用、2已过期等);"`
	TotalRedeemed *uint     `json:"totalRedeemed" form:"totalRedeemed" gorm:"column:total_redeemed;comment:总兑换次数,该码已被兑换的总次数;"`
}

// TableName RedeemCode 表名
func (RedeemCode) TableName() string {
	return "redeem_code"
}

// RedeemLog 兑换记录表
type RedeemLog struct {
	global.Model
	UserId       *uint     `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID,兑换用户;"`
	RedeemCodeId *uint     `json:"redeemCodeId" form:"redeemCodeId" gorm:"column:redeem_code_id;comment:兑换码ID,外键关联兑换码表;"`
	RedeemTime   time.Time `json:"redeemTime" form:"redeemTime" gorm:"column:redeem_time;comment:兑换时间;"`
}

// TableName RedeemLog 表名
func (RedeemLog) TableName() string {
	return "redeem_log"
}
