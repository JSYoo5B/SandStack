package compute

import (
	"net/http"

	appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"
	"github.com/JSYoo5B/SandStack/internal/platform/clock"
	"github.com/JSYoo5B/SandStack/internal/platform/config"
	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
	storecompute "github.com/JSYoo5B/SandStack/internal/store/compute"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	config  config.Config
	service *appcompute.Service
}

func NewRouter(cfg config.Config) http.Handler {
	return NewHandler(cfg).Router()
}

func NewRouterWithService(
	cfg config.Config,
	service *appcompute.Service,
) http.Handler {
	return NewHandlerWithService(cfg, service).Router()
}

func NewHandler(cfg config.Config) Handler {
	return NewHandlerWithService(
		cfg,
		appcompute.NewServiceWithRuntime(
			storecompute.NewMemoryServerRepository(),
			storecompute.NewMemoryKeyPairRepository(),
			storecompute.NewMemoryServerGroupRepository(),
			clock.Wall(),
			idgen.Random(),
		),
	)
}

func NewHandlerWithService(
	cfg config.Config,
	service *appcompute.Service,
) Handler {
	return Handler{
		config:  cfg,
		service: service,
	}
}

func (h Handler) Router() http.Handler {
	router := chi.NewRouter()
	router.Get("/{project_id}", h.version)
	router.Get("/{project_id}/", h.version)
	router.Get("/{project_id}/limits", h.getLimits)
	router.Get("/{project_id}/os-availability-zone", h.listAvailabilityZones)
	router.Get("/{project_id}/os-availability-zone/detail", h.listAvailabilityZones)
	router.Get("/{project_id}/os-services", h.listComputeServices)
	router.Get("/{project_id}/flavors", h.listFlavors)
	router.Get("/{project_id}/flavors/detail", h.listFlavors)
	router.Get("/{project_id}/flavors/{flavor_id}", h.getFlavor)
	router.Get("/{project_id}/flavors/{flavor_id}/os-extra_specs", h.listFlavorExtraSpecs)
	router.Post("/{project_id}/flavors/{flavor_id}/os-extra_specs", h.createFlavorExtraSpecs)
	router.Get("/{project_id}/flavors/{flavor_id}/os-extra_specs/{key}", h.getFlavorExtraSpec)
	router.Put("/{project_id}/flavors/{flavor_id}/os-extra_specs/{key}", h.updateFlavorExtraSpec)
	router.Delete("/{project_id}/flavors/{flavor_id}/os-extra_specs/{key}", h.deleteFlavorExtraSpec)
	router.Get("/{project_id}/servers", h.listServers)
	router.Post("/{project_id}/servers", h.createServer)
	router.Get("/{project_id}/servers/detail", h.listServers)
	router.Get("/{project_id}/servers/{server_id}", h.getServer)
	router.Delete("/{project_id}/servers/{server_id}", h.deleteServer)
	router.Get("/{project_id}/servers/{server_id}/ips", h.listServerAddresses)
	router.Get("/{project_id}/servers/{server_id}/ips/{network}", h.listServerAddressesByNetwork)
	router.Post("/{project_id}/servers/{server_id}/action", h.actionServer)
	router.Get("/{project_id}/servers/{server_id}/metadata", h.getServerMetadata)
	router.Put("/{project_id}/servers/{server_id}/metadata", h.resetServerMetadata)
	router.Post("/{project_id}/servers/{server_id}/metadata", h.updateServerMetadata)
	router.Get("/{project_id}/servers/{server_id}/metadata/{key}", h.getServerMetadatum)
	router.Put("/{project_id}/servers/{server_id}/metadata/{key}", h.setServerMetadatum)
	router.Delete("/{project_id}/servers/{server_id}/metadata/{key}", h.deleteServerMetadatum)
	router.Get("/{project_id}/os-keypairs", h.listKeyPairs)
	router.Post("/{project_id}/os-keypairs", h.createKeyPair)
	router.Get("/{project_id}/os-keypairs/{keypair_name}", h.getKeyPair)
	router.Delete("/{project_id}/os-keypairs/{keypair_name}", h.deleteKeyPair)
	router.Get("/{project_id}/os-server-groups", h.listServerGroups)
	router.Post("/{project_id}/os-server-groups", h.createServerGroup)
	router.Get("/{project_id}/os-server-groups/{group_id}", h.getServerGroup)
	router.Delete("/{project_id}/os-server-groups/{group_id}", h.deleteServerGroup)

	return router
}
