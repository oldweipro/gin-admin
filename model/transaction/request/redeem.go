package request

import (
	"github.com/oldweipro/gin-admin/model/common/request"
	"time"
)

type RedeemCodeRequest struct {
	Pieces     *int  `json:"pieces" form:"pieces"`         // 多少个
	TotalCount *uint `json:"totalCount" form:"totalCount"` // 数量
	Amount     *uint `json:"amount" form:"amount"`         // 面额
	PerLimit   *uint `json:"perLimit" form:"perLimit"`     // 频次数
	Status     *uint `json:"status" form:"status"`         // 频次数
	ExpireTime int64 `json:"expireTime" form:"expireTime"` // 有效期
}

type RedeemCodeSearch struct {
	Amount         *uint      `json:"amount" form:"amount"` // 面额
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type RedeemFishCoinRequest struct {
	RedeemCode string `json:"redeemCode" form:"redeemCode"` // 兑换码
}
