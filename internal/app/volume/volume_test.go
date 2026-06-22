package volume_test

import (
	"testing"
	"time"

	"github.com/JSYoo5B/SandStack/internal/app/volume"
	"github.com/JSYoo5B/SandStack/internal/platform/clock"
	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
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
	service := volume.NewServiceWithClock(clock.Fixed(now))

	created := service.Create(volume.CreateVolume{
		Size: 1,
		Name: "database",
	})

	s.Assert().Equal("2026-06-23T08:30:00.123456", created.CreatedAt)
	s.Assert().Equal("2026-06-23T08:30:00.123456", created.UpdatedAt)
}

func (s *VolumeSuite) TestCreateVolumeUsesInjectedIDGenerator() {
	service := volume.NewServiceWithRuntime(
		clock.Fixed(time.Time{}),
		idgen.Fixed("volume-id"),
	)

	created := service.Create(volume.CreateVolume{
		Size: 1,
		Name: "database",
	})

	s.Assert().Equal("vol-volume-id", created.ID)
}
