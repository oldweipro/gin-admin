package response

import "github.com/oldweipro/gin-admin/pkg/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
