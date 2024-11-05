package router

import (
	"github.com/oldweipro/gin-admin/router/example"
	"github.com/oldweipro/gin-admin/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}
