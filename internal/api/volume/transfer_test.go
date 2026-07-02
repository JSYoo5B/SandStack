package volume_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/volume"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/transfers"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/volumes"
	"github.com/stretchr/testify/suite"
)

type TransferSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestTransferSuite(t *testing.T) {
	suite.Run(t, new(TransferSuite))
}

func (s *TransferSuite) SetupTest() {
	s.server = httptest.NewServer(
		volume.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *TransferSuite) TearDownTest() {
	s.server.Close()
}

func (s *TransferSuite) TestCreateTransferThenListTransfers() {
	volume := s.createVolume("database")

	created := s.createTransfer(volume.ID, "database-transfer")
	list := s.listTransfers()

	s.Assert().NotEmpty(created.ID)
	s.Assert().NotEmpty(created.AuthKey)
	s.Assert().Equal(volume.ID, created.VolumeID)
	s.Assert().Equal("database-transfer", created.Name)
	s.Require().Len(list, 1)
	s.Assert().Equal(created.ID, list[0].ID)
}

func (s *TransferSuite) TestGetTransfer() {
	volume := s.createVolume("database")
	created := s.createTransfer(volume.ID, "database-transfer")

	found, err := transfers.Get(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/demo"),
		created.ID,
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(found)

	s.Assert().Equal(created.ID, found.ID)
	s.Assert().Equal(volume.ID, found.VolumeID)
}

func (s *TransferSuite) TestDeleteTransfer() {
	volume := s.createVolume("database")
	created := s.createTransfer(volume.ID, "database-transfer")

	err := transfers.Delete(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/demo"),
		created.ID,
	).ExtractErr()
	s.Require().NoError(err)

	list := s.listTransfers()

	s.Assert().Empty(list)
}

func (s *TransferSuite) TestAcceptTransfer() {
	volume := s.createVolume("database")
	created := s.createTransfer(volume.ID, "database-transfer")

	accepted, err := transfers.Accept(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/demo"),
		created.ID,
		transfers.AcceptOpts{AuthKey: created.AuthKey},
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(accepted)

	list := s.listTransfers()

	s.Assert().Equal(created.ID, accepted.ID)
	s.Assert().Equal(volume.ID, accepted.VolumeID)
	s.Assert().Empty(list)
}

func (s *TransferSuite) listTransfers() []transfers.Transfer {
	pages, err := transfers.List(
		testhelper.ServiceClient(s.server.URL+"/demo"),
		nil,
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	list, err := transfers.ExtractTransfers(pages)
	s.Require().NoError(err)

	return list
}

func (s *TransferSuite) createTransfer(
	volumeID string,
	name string,
) *transfers.Transfer {
	created, err := transfers.Create(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/demo"),
		transfers.CreateOpts{
			VolumeID: volumeID,
			Name:     name,
		},
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(created)

	return created
}

func (s *TransferSuite) createVolume(name string) *volumes.Volume {
	created, err := volumes.Create(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/demo"),
		volumes.CreateOpts{
			Size: 1,
			Name: name,
		},
		nil,
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(created)

	return created
}
