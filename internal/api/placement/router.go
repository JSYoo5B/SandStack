package placement

import (
	"net/http"

	appplacement "github.com/JSYoo5B/SandStack/internal/app/placement"
	"github.com/JSYoo5B/SandStack/internal/platform/config"
	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
	storeplacement "github.com/JSYoo5B/SandStack/internal/store/placement"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	config  config.Config
	service *appplacement.Service
}

func NewRouter(cfg config.Config) http.Handler {
	return NewHandler(cfg).Router()
}

func NewRouterWithService(
	cfg config.Config,
	service *appplacement.Service,
) http.Handler {
	return NewHandlerWithService(cfg, service).Router()
}

func NewHandler(cfg config.Config) Handler {
	return NewHandlerWithService(
		cfg,
		appplacement.NewServiceWithRepositories(
			storeplacement.NewMemoryResourceProviderRepository(),
			storeplacement.NewMemoryInventoryRepository(),
			storeplacement.NewMemoryTraitRepository(),
			storeplacement.NewMemoryAggregateRepository(),
			storeplacement.NewMemoryUsageRepository(),
			storeplacement.NewMemoryAllocationRepository(),
			storeplacement.NewMemoryResourceClassRepository(),
			storeplacement.NewMemoryTraitCatalogRepository(),
			idgen.Random(),
		),
	)
}

func NewHandlerWithService(
	cfg config.Config,
	service *appplacement.Service,
) Handler {
	return Handler{
		config:  cfg,
		service: service,
	}
}

func (h Handler) Router() http.Handler {
	router := chi.NewRouter()
	router.Get("/", h.versions)
	router.Get("/resource_providers", h.listResourceProviders)
	router.Post("/resource_providers", h.createResourceProvider)
	router.Get(
		"/resource_providers/{resource_provider_uuid}",
		h.getResourceProvider,
	)
	router.Delete(
		"/resource_providers/{resource_provider_uuid}",
		h.deleteResourceProvider,
	)
	router.Get(
		"/resource_providers/{resource_provider_uuid}/inventories",
		h.getInventories,
	)
	router.Put(
		"/resource_providers/{resource_provider_uuid}/inventories",
		h.updateInventories,
	)
	router.Delete(
		"/resource_providers/{resource_provider_uuid}/inventories",
		h.deleteInventories,
	)
	router.Get(
		"/resource_providers/{resource_provider_uuid}/inventories/{resource_class}",
		h.getInventory,
	)
	router.Put(
		"/resource_providers/{resource_provider_uuid}/inventories/{resource_class}",
		h.updateInventory,
	)
	router.Delete(
		"/resource_providers/{resource_provider_uuid}/inventories/{resource_class}",
		h.deleteInventory,
	)
	router.Get(
		"/resource_providers/{resource_provider_uuid}/traits",
		h.getTraits,
	)
	router.Put(
		"/resource_providers/{resource_provider_uuid}/traits",
		h.updateTraits,
	)
	router.Delete(
		"/resource_providers/{resource_provider_uuid}/traits",
		h.deleteTraits,
	)
	router.Get(
		"/resource_providers/{resource_provider_uuid}/aggregates",
		h.getAggregates,
	)
	router.Put(
		"/resource_providers/{resource_provider_uuid}/aggregates",
		h.updateAggregates,
	)
	router.Get(
		"/resource_providers/{resource_provider_uuid}/usages",
		h.getUsages,
	)
	router.Get(
		"/resource_providers/{resource_provider_uuid}/allocations",
		h.getAllocations,
	)
	router.Get("/allocation_candidates", h.listAllocationCandidates)
	router.Get("/resource_classes", h.listResourceClasses)
	router.Post("/resource_classes", h.createResourceClass)
	router.Get("/resource_classes/{resource_class}", h.getResourceClass)
	router.Put("/resource_classes/{resource_class}", h.updateResourceClass)
	router.Delete("/resource_classes/{resource_class}", h.deleteResourceClass)
	router.Get("/traits", h.listTraits)
	router.Get("/traits/{trait_name}", h.getTrait)
	router.Put("/traits/{trait_name}", h.createTrait)
	router.Delete("/traits/{trait_name}", h.deleteTrait)

	return router
}
