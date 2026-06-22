package compute_test

import (
	"testing"
	"time"

	"github.com/JSYoo5B/SandStack/internal/app/compute"
	"github.com/JSYoo5B/SandStack/internal/platform/clock"
	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
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
	service := compute.NewServiceWithClock(clock.Fixed(now))

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
