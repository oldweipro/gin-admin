package v1

import (
	"github.com/oldweipro/gin-admin/api/v1/example"
	"github.com/oldweipro/gin-admin/api/v1/ladder"
	"github.com/oldweipro/gin-admin/api/v1/memo_nexus"
	"github.com/oldweipro/gin-admin/api/v1/openfish"
	"github.com/oldweipro/gin-admin/api/v1/patrol"
	"github.com/oldweipro/gin-admin/api/v1/platform"
	"github.com/oldweipro/gin-admin/api/v1/system"
	"github.com/oldweipro/gin-admin/api/v1/transaction"
)

type ApiGroup struct {
	SystemApiGroup      system.ApiGroup
	ExampleApiGroup     example.ApiGroup
	PatrolApiGroup      patrol.ApiGroup
	LadderApiGroup      ladder.ApiGroup
	PlatformApiGroup    platform.ApiGroup
	OpenfishApiGroup    openfish.ApiGroup
	TransactionApiGroup transaction.ApiGroup
	MemoNexusApiGroup   memo_nexus.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
