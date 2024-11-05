package initialize

import (
	_ "github.com/oldweipro/gin-admin/pkg/source/example"
	_ "github.com/oldweipro/gin-admin/pkg/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
