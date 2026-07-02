package volume_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/volume"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/quotasets"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/volumes"
	"github.com/stretchr/testify/suite"
)

type QuotaSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestQuotaSuite(t *testing.T) {
	suite.Run(t, new(QuotaSuite))
}

func (s *QuotaSuite) SetupTest() {
	s.server = httptest.NewServer(
		volume.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *QuotaSuite) TearDownTest() {
	s.server.Close()
}

func (s *QuotaSuite) TestGetQuotaSet() {
	quotaSet, err := quotasets.Get(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/demo"),
		"demo",
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(quotaSet)

	s.Assert().Equal("demo", quotaSet.ID)
	s.Assert().Equal(1000, quotaSet.Volumes)
	s.Assert().Equal(-1, quotaSet.PerVolumeGigabytes)
}

func (s *QuotaSuite) TestGetDefaultQuotaSet() {
	quotaSet, err := quotasets.GetDefaults(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/demo"),
		"demo",
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(quotaSet)

	s.Assert().Equal("demo", quotaSet.ID)
	s.Assert().Equal(1000, quotaSet.Backups)
}

func (s *QuotaSuite) TestUpdateAndDeleteQuotaSet() {
	client := testhelper.ServiceClient(s.server.URL + "/demo")
	quotaSet, err := quotasets.Update(
		s.T().Context(),
		client,
		"demo",
		quotasets.UpdateOpts{
			Volumes: gophercloud.IntToPointer(12),
		},
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(quotaSet)
	s.Assert().Equal(12, quotaSet.Volumes)

	err = quotasets.Delete(
		s.T().Context(),
		client,
		"demo",
	).ExtractErr()
	s.Require().NoError(err)

	quotaSet, err = quotasets.Get(
		s.T().Context(),
		client,
		"demo",
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(quotaSet)
	s.Assert().Equal(1000, quotaSet.Volumes)
}

func (s *QuotaSuite) TestGetQuotaUsageSet() {
	client := testhelper.ServiceClient(s.server.URL + "/demo")
	created, err := volumes.Create(
		s.T().Context(),
		client,
		volumes.CreateOpts{
			Size: 3,
			Name: "database",
		},
		nil,
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(created)

	usageSet, err := quotasets.GetUsage(
		s.T().Context(),
		client,
		"demo",
	).Extract()
	s.Require().NoError(err)

	s.Assert().Equal("demo", usageSet.ID)
	s.Assert().Equal(1, usageSet.Volumes.InUse)
	s.Assert().Equal(3, usageSet.Gigabytes.InUse)
	s.Assert().Equal(1000, usageSet.Volumes.Limit)
}
