package v1

import (
	"github.com/oldweipro/gin-admin/api/v1/example"
	"github.com/oldweipro/gin-admin/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
}
