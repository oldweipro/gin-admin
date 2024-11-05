package service

import (
	"github.com/oldweipro/gin-admin/service/example"
	"github.com/oldweipro/gin-admin/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}
