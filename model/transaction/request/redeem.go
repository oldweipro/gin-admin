package request

type RedeemRequest struct {
	Pieces     int   `json:"pieces" form:"pieces"`         // 多少个
	TotalCount uint  `json:"totalCount" form:"totalCount"` // 数量
	Amount     uint  `json:"amount" form:"amount"`         // 面额
	PerLimit   uint  `json:"perLimit" form:"perLimit"`     // 频次数
	ExpireTime int64 `json:"expireTime" form:"expireTime"` // 有效期
}
