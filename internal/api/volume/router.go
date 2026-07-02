package volume

import (
	"net/http"

	appvolume "github.com/JSYoo5B/SandStack/internal/app/volume"
	"github.com/JSYoo5B/SandStack/internal/platform/clock"
	"github.com/JSYoo5B/SandStack/internal/platform/config"
	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
	storevolume "github.com/JSYoo5B/SandStack/internal/store/volume"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	config  config.Config
	service *appvolume.Service
}

func NewRouter(cfg config.Config) http.Handler {
	return NewHandler(cfg).Router()
}

func NewRouterWithService(
	cfg config.Config,
	service *appvolume.Service,
) http.Handler {
	return NewHandlerWithService(cfg, service).Router()
}

func NewHandler(cfg config.Config) Handler {
	return NewHandlerWithService(
		cfg,
		appvolume.NewServiceWithRuntime(
			storevolume.NewMemoryRepository(),
			storevolume.NewMemorySnapshotRepository(),
			storevolume.NewMemoryTransferRepository(),
			clock.Wall(),
			idgen.Random(),
		),
	)
}

func NewHandlerWithService(
	cfg config.Config,
	service *appvolume.Service,
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
	router.Get("/{project_id}/volumes/detail", h.listVolumes)
	router.Post("/{project_id}/volumes", h.createVolume)
	router.Get("/{project_id}/volumes/{volume_id}", h.getVolume)
	router.Put("/{project_id}/volumes/{volume_id}", h.updateVolume)
	router.Delete("/{project_id}/volumes/{volume_id}", h.deleteVolume)
	router.Get("/{project_id}/snapshots", h.listSnapshots)
	router.Get("/{project_id}/snapshots/detail", h.listSnapshots)
	router.Post("/{project_id}/snapshots", h.createSnapshot)
	router.Get("/{project_id}/snapshots/{snapshot_id}", h.getSnapshot)
	router.Delete("/{project_id}/snapshots/{snapshot_id}", h.deleteSnapshot)
	router.Get("/{project_id}/os-volume-transfer/detail", h.listTransfers)
	router.Post("/{project_id}/os-volume-transfer", h.createTransfer)
	router.Get("/{project_id}/os-volume-transfer/{transfer_id}", h.getTransfer)
	router.Delete("/{project_id}/os-volume-transfer/{transfer_id}", h.deleteTransfer)
	router.Post("/{project_id}/os-volume-transfer/{transfer_id}/accept", h.acceptTransfer)
	router.Get("/{project_id}/types", h.listVolumeTypes)
	router.Get("/{project_id}/types/{type_id}", h.getVolumeType)

	return router
}
