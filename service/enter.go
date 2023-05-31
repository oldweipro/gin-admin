package service

import (
	"github.com/oldweipro/gin-admin/service/example"
	"github.com/oldweipro/gin-admin/service/ladder"
	"github.com/oldweipro/gin-admin/service/openfish"
	"github.com/oldweipro/gin-admin/service/patrol"
	"github.com/oldweipro/gin-admin/service/system"
	"github.com/oldweipro/gin-admin/service/transaction"
)

type ServiceGroup struct {
	SystemServiceGroup      system.ServiceGroup
	ExampleServiceGroup     example.ServiceGroup
	PatrolServiceGroup      patrol.ServiceGroup
	LadderServiceGroup      ladder.ServiceGroup
	OpenfishServiceGroup    openfish.ServiceGroup
	TransactionServiceGroup transaction.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
