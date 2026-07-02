package api

import (
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/admin"
	"github.com/JSYoo5B/SandStack/internal/api/compute"
	"github.com/JSYoo5B/SandStack/internal/api/identity"
	"github.com/JSYoo5B/SandStack/internal/api/image"
	"github.com/JSYoo5B/SandStack/internal/api/network"
	"github.com/JSYoo5B/SandStack/internal/api/placement"
	"github.com/JSYoo5B/SandStack/internal/api/volume"
	appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"
	appidentity "github.com/JSYoo5B/SandStack/internal/app/identity"
	appimage "github.com/JSYoo5B/SandStack/internal/app/image"
	appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"
	"github.com/JSYoo5B/SandStack/internal/app/requestlog"
	appvolume "github.com/JSYoo5B/SandStack/internal/app/volume"
	"github.com/JSYoo5B/SandStack/internal/platform/clock"
	"github.com/JSYoo5B/SandStack/internal/platform/config"
	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
	storecompute "github.com/JSYoo5B/SandStack/internal/store/compute"
	storeidentity "github.com/JSYoo5B/SandStack/internal/store/identity"
	storeimage "github.com/JSYoo5B/SandStack/internal/store/image"
	storenetwork "github.com/JSYoo5B/SandStack/internal/store/network"
	storevolume "github.com/JSYoo5B/SandStack/internal/store/volume"
	"github.com/go-chi/chi/v5"
)

func NewRouter(cfg config.Config) http.Handler {
	router := chi.NewRouter()
	router.Use(requestID)
	requests := requestlog.NewService()
	router.Use(recordRequests(requests))
	identityService := appidentity.NewServiceWithRepositories(
		cfg,
		appidentity.Repositories{
			Users:     storeidentity.NewMemoryUserRepository(),
			Projects:  storeidentity.NewMemoryProjectRepository(),
			Roles:     storeidentity.NewMemoryRoleRepository(),
			Tokens:    storeidentity.NewMemoryTokenRepository(),
			Services:  storeidentity.NewMemoryServiceRepository(),
			Endpoints: storeidentity.NewMemoryEndpointRepository(),
		},
	)
	identityHandler := identity.NewHandlerWithService(cfg, identityService)
	computeService := appcompute.NewServiceWithRuntime(
		storecompute.NewMemoryServerRepository(),
		storecompute.NewMemoryKeyPairRepository(),
		storecompute.NewMemoryServerGroupRepository(),
		storecompute.NewMemoryAggregateRepository(),
		clock.Wall(),
		idgen.Random(),
	)
	imageService := appimage.NewServiceWithRuntime(
		storeimage.NewMemoryRepository(),
		clock.Wall(),
		idgen.Random(),
	)
	networkService := appnetwork.NewServiceWithRepositories(
		storenetwork.NewMemoryNetworkRepository(),
		storenetwork.NewMemorySubnetRepository(),
		storenetwork.NewMemoryPortRepository(),
		storenetwork.NewMemorySecurityGroupRepository(),
		storenetwork.NewMemorySecurityGroupRuleRepository(),
		storenetwork.NewMemoryRouterRepository(),
		storenetwork.NewMemoryFloatingIPRepository(),
		storenetwork.NewMemoryRouterInterfaceRepository(),
		idgen.Random(),
	)
	volumeService := appvolume.NewServiceWithRuntime(
		storevolume.NewMemoryRepository(),
		storevolume.NewMemorySnapshotRepository(),
		clock.Wall(),
		idgen.Random(),
	)

	router.Mount("/_sandstack", admin.NewRouterWithState(func() {
		identityService.Reset()
		computeService.Reset()
		imageService.Reset()
		networkService.Reset()
		volumeService.Reset()
		requests.Reset()
	}, requests))

	router.Get("/identity", identityHandler.Discovery())
	router.Get("/identity/", identityHandler.Discovery())
	router.Mount("/identity/v3", identityHandler.Router())

	router.Mount("/compute/v2.1", compute.NewRouterWithService(cfg, computeService))
	router.Mount("/image/v2", image.NewRouterWithService(cfg, imageService))
	router.Mount("/network/v2.0", network.NewRouterWithService(cfg, networkService))
	router.Mount("/placement", placement.NewRouter(cfg))
	router.Mount("/volume/v3", volume.NewRouterWithService(cfg, volumeService))

	return router
}
