package response

import (
	"github.com/oldweipro/gin-admin/model/system"
)

type SysAPIResponse struct {
	Api system.SysApi `json:"api"`
}

type SysAPIListResponse struct {
	Apis []system.SysApi `json:"apis"`
}

type SysSyncApis struct {
	NewApis    []system.SysApi `json:"newApis"`
	DeleteApis []system.SysApi `json:"deleteApis"`
}
