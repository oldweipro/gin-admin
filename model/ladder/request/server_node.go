package request

import (
	"github.com/oldweipro/gin-admin/model/common/request"
	"github.com/oldweipro/gin-admin/model/ladder"
	"time"
)

type ServerNodeSearch struct {
	ladder.ServerNode
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type ServerNodeResponse struct {
	Id         uint   `json:"id" form:"id"`
	Region     string `json:"region" form:"region"`
	Bandwidth  string `json:"bandwidth" form:"bandwidth"`
	Up         int64  `json:"up" form:"up"`
	Down       int64  `json:"down" form:"down"`
	ExpiryTime int64  `json:"expiryTime"`
}
