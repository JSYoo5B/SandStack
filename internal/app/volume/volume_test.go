package volume_test

import (
	"testing"
	"time"

	"github.com/JSYoo5B/SandStack/internal/app/volume"
	"github.com/JSYoo5B/SandStack/internal/platform/clock"
	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
	storevolume "github.com/JSYoo5B/SandStack/internal/store/volume"
	"github.com/stretchr/testify/suite"
)

type VolumeSuite struct {
	suite.Suite
}

func TestVolumeSuite(t *testing.T) {
	suite.Run(t, new(VolumeSuite))
}

func (s *VolumeSuite) TestCreateVolumeUsesInjectedClock() {
	now := time.Date(2026, 6, 23, 8, 30, 0, 123456000, time.UTC)
	service := volume.NewServiceWithRuntime(
		storevolume.NewMemoryRepository(),
		storevolume.NewMemorySnapshotRepository(),
		storevolume.NewMemoryTransferRepository(),
		clock.Fixed(now),
		idgen.Random(),
	)

	created := service.Create(volume.CreateVolume{
		Size: 1,
		Name: "database",
	})

	s.Assert().Equal("2026-06-23T08:30:00.123456", created.CreatedAt)
	s.Assert().Equal("2026-06-23T08:30:00.123456", created.UpdatedAt)
}

func (s *VolumeSuite) TestCreateVolumeUsesInjectedIDGenerator() {
	service := volume.NewServiceWithRuntime(
		storevolume.NewMemoryRepository(),
		storevolume.NewMemorySnapshotRepository(),
		storevolume.NewMemoryTransferRepository(),
		clock.Fixed(time.Time{}),
		idgen.Fixed("volume-id"),
	)

	created := service.Create(volume.CreateVolume{
		Size: 1,
		Name: "database",
	})

	s.Assert().Equal("vol-volume-id", created.ID)
}

func (s *VolumeSuite) TestGetVolumeMakesCreatedVolumeAvailable() {
	now := time.Date(2026, 6, 23, 8, 30, 0, 123456000, time.UTC)
	service := volume.NewServiceWithRuntime(
		storevolume.NewMemoryRepository(),
		storevolume.NewMemorySnapshotRepository(),
		storevolume.NewMemoryTransferRepository(),
		clock.Fixed(now),
		idgen.Fixed("volume-id"),
	)
	created := service.Create(volume.CreateVolume{
		Size: 1,
		Name: "database",
	})

	found, err := service.Get(created.ID)
	s.Require().NoError(err)

	s.Assert().Equal("creating", created.Status)
	s.Assert().Equal("available", found.Status)
	s.Assert().Equal("2026-06-23T08:30:00.123456", found.UpdatedAt)
}

func (s *VolumeSuite) TestResetClearsVolumes() {
	service := volume.NewServiceWithRuntime(
		storevolume.NewMemoryRepository(),
		storevolume.NewMemorySnapshotRepository(),
		storevolume.NewMemoryTransferRepository(),
		clock.Fixed(time.Time{}),
		idgen.Fixed("volume-id"),
	)
	service.Create(volume.CreateVolume{
		Size: 1,
		Name: "database",
	})
	service.CreateSnapshot(volume.CreateSnapshot{
		Name:     "database-snapshot",
		VolumeID: "vol-volume-id",
	})
	service.CreateTransfer(volume.CreateTransfer{
		Name:     "database-transfer",
		VolumeID: "vol-volume-id",
	})

	service.Reset()

	s.Assert().Empty(service.List())
	s.Assert().Empty(service.ListSnapshots())
	s.Assert().Empty(service.ListTransfers())
}
