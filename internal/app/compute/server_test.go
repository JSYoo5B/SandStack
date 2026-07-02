package compute_test

import (
	"testing"
	"time"

	"github.com/JSYoo5B/SandStack/internal/app/compute"
	"github.com/JSYoo5B/SandStack/internal/platform/clock"
	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
	storecompute "github.com/JSYoo5B/SandStack/internal/store/compute"
	"github.com/stretchr/testify/suite"
)

type ServerSuite struct {
	suite.Suite
}

func TestServerSuite(t *testing.T) {
	suite.Run(t, new(ServerSuite))
}

func (s *ServerSuite) TestCreateServerUsesInjectedClock() {
	now := time.Date(2026, 6, 23, 8, 30, 0, 0, time.UTC)
	service := compute.NewServiceWithRuntime(
		storecompute.NewMemoryServerRepository(),
		storecompute.NewMemoryKeyPairRepository(),
		storecompute.NewMemoryServerGroupRepository(),
		clock.Fixed(now),
		idgen.Random(),
	)

	server := service.CreateServer(compute.CreateServer{
		Name:     "web",
		ImageID:  "img-1",
		FlavorID: "1",
	})

	s.Assert().Equal("2026-06-23T08:30:00Z", server.CreatedAt)
	s.Assert().Equal("2026-06-23T08:30:00Z", server.UpdatedAt)
}

func (s *ServerSuite) TestCreateServerUsesInjectedIDGenerator() {
	service := compute.NewServiceWithRuntime(
		storecompute.NewMemoryServerRepository(),
		storecompute.NewMemoryKeyPairRepository(),
		storecompute.NewMemoryServerGroupRepository(),
		clock.Fixed(time.Time{}),
		idgen.Fixed("server-id"),
	)

	server := service.CreateServer(compute.CreateServer{
		Name:     "web",
		ImageID:  "img-1",
		FlavorID: "1",
	})

	s.Assert().Equal("srv-server-id", server.ID)
}

func (s *ServerSuite) TestGetServerActivatesCreatedServer() {
	now := time.Date(2026, 6, 23, 8, 30, 0, 0, time.UTC)
	service := compute.NewServiceWithRuntime(
		storecompute.NewMemoryServerRepository(),
		storecompute.NewMemoryKeyPairRepository(),
		storecompute.NewMemoryServerGroupRepository(),
		clock.Fixed(now),
		idgen.Fixed("server-id"),
	)
	created := service.CreateServer(compute.CreateServer{
		Name:     "web",
		ImageID:  "img-1",
		FlavorID: "1",
	})

	found, err := service.GetServer(created.ID)
	s.Require().NoError(err)

	s.Assert().Equal("BUILD", created.Status)
	s.Assert().Equal("ACTIVE", found.Status)
	s.Assert().Equal(100, found.Progress)
	s.Assert().Equal("2026-06-23T08:30:00Z", found.UpdatedAt)
}

func (s *ServerSuite) TestResetClearsServers() {
	service := compute.NewServiceWithRuntime(
		storecompute.NewMemoryServerRepository(),
		storecompute.NewMemoryKeyPairRepository(),
		storecompute.NewMemoryServerGroupRepository(),
		clock.Fixed(time.Time{}),
		idgen.Fixed("server-id"),
	)
	service.CreateServer(compute.CreateServer{
		Name:     "web",
		ImageID:  "img-1",
		FlavorID: "1",
	})
	service.CreateKeyPair(compute.CreateKeyPair{Name: "default"})

	service.Reset()

	s.Assert().Empty(service.ListServers())
	s.Assert().Empty(service.ListKeyPairs())
}
