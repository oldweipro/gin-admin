package response

import "github.com/oldweipro/gin-admin/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
