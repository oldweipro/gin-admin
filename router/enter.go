package router

import (
	"github.com/oldweipro/gin-admin/router/example"
	"github.com/oldweipro/gin-admin/router/ladder"
	"github.com/oldweipro/gin-admin/router/openfish"
	"github.com/oldweipro/gin-admin/router/patrol"
	"github.com/oldweipro/gin-admin/router/system"
	"github.com/oldweipro/gin-admin/router/transaction"
)

type RouterGroup struct {
	System      system.RouterGroup
	Example     example.RouterGroup
	Patrol      patrol.RouterGroup
	Ladder      ladder.RouterGroup
	Openfish    openfish.RouterGroup
	Transaction transaction.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
