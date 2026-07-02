package compute_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/compute"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/keypairs"
	"github.com/stretchr/testify/suite"
)

type KeyPairSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestKeyPairSuite(t *testing.T) {
	suite.Run(t, new(KeyPairSuite))
}

func (s *KeyPairSuite) SetupTest() {
	s.server = httptest.NewServer(
		compute.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *KeyPairSuite) TearDownTest() {
	s.server.Close()
}

func (s *KeyPairSuite) TestListKeyPairs() {
	list := s.listKeyPairs()

	s.Assert().Empty(list)
}

func (s *KeyPairSuite) TestCreateKeyPairThenListKeyPairs() {
	created := s.createKeyPair("default")

	list := s.listKeyPairs()

	s.Assert().Equal("default", created.Name)
	s.Assert().Equal("ssh", created.Type)
	s.Assert().NotEmpty(created.Fingerprint)
	s.Require().Len(list, 1)
	s.Assert().Equal("default", list[0].Name)
}

func (s *KeyPairSuite) TestGetKeyPair() {
	created := s.createKeyPair("default")

	found, err := keypairs.Get(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/demo"),
		created.Name,
		nil,
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(found)

	s.Assert().Equal(created.Name, found.Name)
	s.Assert().Equal(created.PublicKey, found.PublicKey)
}

func (s *KeyPairSuite) TestDeleteKeyPair() {
	created := s.createKeyPair("default")

	err := keypairs.Delete(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/demo"),
		created.Name,
		nil,
	).ExtractErr()
	s.Require().NoError(err)

	list := s.listKeyPairs()

	s.Assert().Empty(list)
}

func (s *KeyPairSuite) listKeyPairs() []keypairs.KeyPair {
	pages, err := keypairs.List(
		testhelper.ServiceClient(s.server.URL+"/demo"),
		nil,
	).AllPages(s.T().Context())
	s.Require().NoError(err)

	list, err := keypairs.ExtractKeyPairs(pages)
	s.Require().NoError(err)

	return list
}

func (s *KeyPairSuite) createKeyPair(name string) *keypairs.KeyPair {
	created, err := keypairs.Create(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL+"/demo"),
		keypairs.CreateOpts{
			Name:      name,
			PublicKey: "ssh-rsa test-public-key",
		},
	).Extract()
	s.Require().NoError(err)
	s.Require().NotNil(created)

	return created
}
