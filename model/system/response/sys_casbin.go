package response

import (
	"github.com/oldweipro/gin-admin/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
