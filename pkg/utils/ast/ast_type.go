package ast

type Type string

func (r Type) String() string {
	return string(r)
}

func (r Type) Group() string {
	switch r {
	case TypePackageApiEnter:
		return "ApiGroup"
	case TypePackageRouterEnter:
		return "RouterGroup"
	case TypePackageServiceEnter:
		return "ServiceGroup"
	case TypePackageApiModuleEnter:
		return "ApiGroup"
	case TypePackageRouterModuleEnter:
		return "RouterGroup"
	case TypePackageServiceModuleEnter:
		return "ServiceGroup"
	default:
		return ""
	}
}

const (
	TypePackageApiEnter           = "PackageApiEnter"           // server/api/v1/enter.go
	TypePackageRouterEnter        = "PackageRouterEnter"        // server/router/enter.go
	TypePackageServiceEnter       = "PackageServiceEnter"       // server/service/enter.go
	TypePackageApiModuleEnter     = "PackageApiModuleEnter"     // server/api/v1/{package}/enter.go
	TypePackageRouterModuleEnter  = "PackageRouterModuleEnter"  // server/router/{package}/enter.go
	TypePackageServiceModuleEnter = "PackageServiceModuleEnter" // server/service/{package}/enter.go
	TypePackageInitializeGorm     = "PackageInitializeGorm"     // server/initialize/gorm_biz.go
	TypePackageInitializeRouter   = "PackageInitializeRouter"   // server/initialize/router_biz.go
)
