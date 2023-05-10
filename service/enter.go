package service

import (
	"github.com/oldweipro/gin-admin/service/example"
	"github.com/oldweipro/gin-admin/service/ladder"
	"github.com/oldweipro/gin-admin/service/openfish"
	"github.com/oldweipro/gin-admin/service/patrol"
	"github.com/oldweipro/gin-admin/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	PatrolServiceGroup   patrol.ServiceGroup
	LadderServiceGroup   ladder.ServiceGroup
	OpenfishServiceGroup openfish.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
