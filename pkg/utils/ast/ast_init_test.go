package ast

import (
	"github.com/oldweipro/gin-admin/pkg/app"
	"path/filepath"
)

func init() {
	app.Config.AutoCode.Root, _ = filepath.Abs("../../../")
	app.Config.AutoCode.Server = "server"
}
