package requestlog_test

import (
	"net/http"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/app/requestlog"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) TestAddThenListRecords() {
	service := requestlog.NewService()

	service.Add(requestlog.Record{
		ID:     "req-1",
		Method: http.MethodGet,
		Path:   "/image/v2/images",
		Status: http.StatusOK,
	})

	records := service.List()

	s.Require().Len(records, 1)
	s.Assert().Equal("req-1", records[0].ID)
	s.Assert().Equal(http.MethodGet, records[0].Method)
	s.Assert().Equal("/image/v2/images", records[0].Path)
	s.Assert().Equal(http.StatusOK, records[0].Status)
}

func (s *ServiceSuite) TestResetClearsRecords() {
	service := requestlog.NewService()
	service.Add(requestlog.Record{ID: "req-1"})

	service.Reset()

	s.Assert().Empty(service.List())
}
