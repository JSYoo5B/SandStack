package network

import (
	"net/http"

	appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"
	"github.com/JSYoo5B/SandStack/internal/platform/config"
	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
	storenetwork "github.com/JSYoo5B/SandStack/internal/store/network"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	config  config.Config
	service *appnetwork.Service
}

func NewRouter(cfg config.Config) http.Handler {
	return NewHandler(cfg).Router()
}

func NewRouterWithService(
	cfg config.Config,
	service *appnetwork.Service,
) http.Handler {
	return NewHandlerWithService(cfg, service).Router()
}

func NewHandler(cfg config.Config) Handler {
	return NewHandlerWithService(
		cfg,
		appnetwork.NewServiceWithRepositories(
			storenetwork.NewMemoryNetworkRepository(),
			storenetwork.NewMemorySubnetRepository(),
			storenetwork.NewMemoryPortRepository(),
			storenetwork.NewMemorySecurityGroupRepository(),
			storenetwork.NewMemorySecurityGroupRuleRepository(),
			storenetwork.NewMemoryRouterRepository(),
			idgen.Random(),
		),
	)
}

func NewHandlerWithService(
	cfg config.Config,
	service *appnetwork.Service,
) Handler {
	return Handler{
		config:  cfg,
		service: service,
	}
}

func (h Handler) Router() http.Handler {
	router := chi.NewRouter()
	router.Get("/", h.versions)
	router.Get("/networks", h.listNetworks)
	router.Post("/networks", h.createNetwork)
	router.Get("/networks/{network_id}", h.getNetwork)
	router.Delete("/networks/{network_id}", h.deleteNetwork)
	router.Get("/subnets", h.listSubnets)
	router.Post("/subnets", h.createSubnet)
	router.Get("/subnets/{subnet_id}", h.getSubnet)
	router.Delete("/subnets/{subnet_id}", h.deleteSubnet)
	router.Get("/ports", h.listPorts)
	router.Post("/ports", h.createPort)
	router.Get("/ports/{port_id}", h.getPort)
	router.Delete("/ports/{port_id}", h.deletePort)
	router.Get("/security-groups", h.listSecurityGroups)
	router.Post("/security-groups", h.createSecurityGroup)
	router.Get("/security-groups/{security_group_id}", h.getSecurityGroup)
	router.Delete("/security-groups/{security_group_id}", h.deleteSecurityGroup)
	router.Get("/security-group-rules", h.listSecurityGroupRules)
	router.Post("/security-group-rules", h.createSecurityGroupRule)
	router.Get("/security-group-rules/{security_group_rule_id}", h.getSecurityGroupRule)
	router.Delete("/security-group-rules/{security_group_rule_id}", h.deleteSecurityGroupRule)
	router.Get("/routers", h.listRouters)
	router.Post("/routers", h.createRouter)
	router.Get("/routers/{router_id}", h.getRouter)
	router.Delete("/routers/{router_id}", h.deleteRouter)

	return router
}
