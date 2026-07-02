package identity_test

import (
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/identity"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/tokens"
	"github.com/stretchr/testify/suite"
)

type TokenSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestTokenSuite(t *testing.T) {
	suite.Run(t, new(TokenSuite))
}

func (s *TokenSuite) SetupTest() {
	s.server = httptest.NewServer(
		identity.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *TokenSuite) TearDownTest() {
	s.server.Close()
}

func (s *TokenSuite) TestPasswordAuth() {
	result := tokens.Create(
		s.T().Context(),
		testhelper.ServiceClient(s.server.URL),
		testhelper.PasswordAuthOptions(),
	)
	tokenID, err := result.ExtractTokenID()
	s.Require().NoError(err)

	user, err := result.ExtractUser()
	s.Require().NoError(err)
	s.Require().NotNil(user)

	project, err := result.ExtractProject()
	s.Require().NoError(err)
	s.Require().NotNil(project)

	catalog, err := result.ExtractServiceCatalog()
	s.Require().NoError(err)
	s.Require().NotNil(catalog)

	s.Assert().NotEmpty(tokenID)
	s.Assert().Equal("admin", user.Name)
	s.Assert().Equal("demo", project.Name)
	s.Assert().Equal(
		map[string]bool{
			"identity":  true,
			"compute":   true,
			"network":   true,
			"image":     true,
			"volumev3":  true,
			"placement": true,
		},
		catalogServices(catalog.Entries),
	)
	s.Assert().Equal(
		s.server.URL+"/compute/v2.1/demo",
		computeEndpoint(catalog.Entries),
	)
}

func (s *TokenSuite) TestGetIssuedToken() {
	client := testhelper.ServiceClient(s.server.URL)
	result := tokens.Create(
		s.T().Context(),
		client,
		testhelper.PasswordAuthOptions(),
	)
	tokenID, err := result.ExtractTokenID()
	s.Require().NoError(err)
	s.Require().NotEmpty(tokenID)

	getResult := tokens.Get(
		s.T().Context(),
		client,
		tokenID,
	)
	found, err := getResult.ExtractToken()
	s.Require().NoError(err)
	s.Require().NotNil(found)

	user, err := getResult.ExtractUser()
	s.Require().NoError(err)
	s.Require().NotNil(user)
	project, err := getResult.ExtractProject()
	s.Require().NoError(err)
	s.Require().NotNil(project)

	s.Assert().Equal(tokenID, found.ID)
	s.Assert().Equal("admin", user.Name)
	s.Assert().Equal("demo", project.Name)
}

func (s *TokenSuite) TestValidateAndRevokeIssuedToken() {
	client := testhelper.ServiceClient(s.server.URL)
	result := tokens.Create(
		s.T().Context(),
		client,
		testhelper.PasswordAuthOptions(),
	)
	tokenID, err := result.ExtractTokenID()
	s.Require().NoError(err)
	s.Require().NotEmpty(tokenID)

	ok, err := tokens.Validate(
		s.T().Context(),
		client,
		tokenID,
	)
	s.Require().NoError(err)
	s.Assert().True(ok)

	revokeResult := tokens.Revoke(
		s.T().Context(),
		client,
		tokenID,
	)
	s.Require().NoError(revokeResult.Err)

	ok, err = tokens.Validate(
		s.T().Context(),
		client,
		tokenID,
	)
	s.Require().NoError(err)
	s.Assert().False(ok)
}

func computeEndpoint(catalog []tokens.CatalogEntry) string {
	for _, entry := range catalog {
		if entry.Type != "compute" {
			continue
		}

		for _, endpoint := range entry.Endpoints {
			if endpoint.Interface == "public" {
				return endpoint.URL
			}
		}
	}

	return ""
}

func catalogServices(catalog []tokens.CatalogEntry) map[string]bool {
	services := make(map[string]bool, len(catalog))
	for _, entry := range catalog {
		services[entry.Type] = true
	}

	return services
}
