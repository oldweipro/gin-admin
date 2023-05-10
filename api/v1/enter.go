package v1

import (
	"github.com/oldweipro/gin-admin/api/v1/example"
	"github.com/oldweipro/gin-admin/api/v1/ladder"
	"github.com/oldweipro/gin-admin/api/v1/openfish"
	"github.com/oldweipro/gin-admin/api/v1/patrol"
	"github.com/oldweipro/gin-admin/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	PatrolApiGroup   patrol.ApiGroup
	LadderApiGroup   ladder.ApiGroup
	OpenfishApiGroup openfish.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
