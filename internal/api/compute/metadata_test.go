package compute_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/compute"
	appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"
	"github.com/JSYoo5B/SandStack/internal/platform/clock"
	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
	storecompute "github.com/JSYoo5B/SandStack/internal/store/compute"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servers"
	"github.com/stretchr/testify/suite"
)

type MetadataSuite struct {
	suite.Suite
	server  *httptest.Server
	service *appcompute.Service
}

func TestMetadataSuite(t *testing.T) {
	suite.Run(t, new(MetadataSuite))
}

func (s *MetadataSuite) SetupTest() {
	s.service = appcompute.NewServiceWithRuntime(
		storecompute.NewMemoryServerRepository(),
		storecompute.NewMemoryKeyPairRepository(),
		clock.Wall(),
		idgen.Random(),
	)
	s.server = httptest.NewServer(
		compute.NewRouterWithService(
			testhelper.DefaultConfig(),
			s.service,
		),
	)
}

func (s *MetadataSuite) TearDownTest() {
	s.server.Close()
}

func (s *MetadataSuite) TestServerMetadataLifecycle() {
	created := s.service.CreateServer(appcompute.CreateServer{
		Name:     "web",
		ImageID:  "img-1",
		FlavorID: "1",
		Metadata: map[string]string{
			"role": "web",
		},
	})
	client := testhelper.ServiceClient(s.server.URL + "/demo")

	initial, err := servers.Metadata(
		s.T().Context(),
		client,
		created.ID,
	).Extract()
	s.Require().NoError(err)

	reset, err := servers.ResetMetadata(
		s.T().Context(),
		client,
		created.ID,
		servers.MetadataOpts{"role": "api"},
	).Extract()
	s.Require().NoError(err)

	updated, err := servers.UpdateMetadata(
		s.T().Context(),
		client,
		created.ID,
		servers.MetadataOpts{"owner": "qa"},
	).Extract()
	s.Require().NoError(err)

	s.Assert().Equal(map[string]string{"role": "web"}, initial)
	s.Assert().Equal(map[string]string{"role": "api"}, reset)
	s.Assert().Equal(
		map[string]string{
			"role":  "api",
			"owner": "qa",
		},
		updated,
	)
}

func (s *MetadataSuite) TestServerMetadatumLifecycle() {
	created := s.service.CreateServer(appcompute.CreateServer{
		Name:     "web",
		ImageID:  "img-1",
		FlavorID: "1",
		Metadata: map[string]string{
			"role": "web",
		},
	})
	client := testhelper.ServiceClient(s.server.URL + "/demo")

	initial, err := servers.Metadatum(
		s.T().Context(),
		client,
		created.ID,
		"role",
	).Extract()
	s.Require().NoError(err)

	updated, err := servers.CreateMetadatum(
		s.T().Context(),
		client,
		created.ID,
		servers.MetadatumOpts{"role": "api"},
	).Extract()
	s.Require().NoError(err)

	err = servers.DeleteMetadatum(
		s.T().Context(),
		client,
		created.ID,
		"role",
	).ExtractErr()
	s.Require().NoError(err)

	final, err := servers.Metadata(
		s.T().Context(),
		client,
		created.ID,
	).Extract()
	s.Require().NoError(err)

	s.Assert().Equal(map[string]string{"role": "web"}, initial)
	s.Assert().Equal(map[string]string{"role": "api"}, updated)
	s.Assert().Empty(final)
}
