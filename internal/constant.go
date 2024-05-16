package internal

const (
	DashboardURL       = "https://dashboard.cloud-ara.tyk.io"
	cloudPath          = "cloud"
	CoreConfigFileName = "core_config.yaml"
	CoreConfig         = "core_config"
	DefaultConfigDir   = ".tykctl"
	Config             = "config"
	Init               = "init"
	CurrentConfig      = "current_config"

	CurrentService = "current_service"
	Cloud          = "cloud"
	Gateway        = "gateway"
	All            = "all"
)

var AllowedServices = []string{All, Cloud, Gateway}
