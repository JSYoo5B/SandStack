package image_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/JSYoo5B/SandStack/internal/api/image"
	"github.com/JSYoo5B/SandStack/internal/testhelper"
	"github.com/gophercloud/gophercloud/v2/openstack/image/v2/tasks"
	"github.com/stretchr/testify/suite"
)

type TaskSuite struct {
	suite.Suite
	server *httptest.Server
}

func TestTaskSuite(t *testing.T) {
	suite.Run(t, new(TaskSuite))
}

func (s *TaskSuite) SetupTest() {
	s.server = httptest.NewServer(
		image.NewRouter(testhelper.DefaultConfig()),
	)
}

func (s *TaskSuite) TearDownTest() {
	s.server.Close()
}

func (s *TaskSuite) TestCreateGetAndListTask() {
	client := testhelper.ServiceClient(s.server.URL)

	created, err := tasks.Create(
		context.Background(),
		client,
		tasks.CreateOpts{
			Type: "import",
			Input: map[string]any{
				"image_id": "image-1",
			},
		},
	).Extract()
	s.Require().NoError(err)

	found, err := tasks.Get(context.Background(), client, created.ID).Extract()
	s.Require().NoError(err)

	pages, err := tasks.List(client, nil).AllPages(context.Background())
	s.Require().NoError(err)

	listed, err := tasks.ExtractTasks(pages)
	s.Require().NoError(err)

	s.Assert().Equal("import", found.Type)
	s.Assert().Equal(tasks.TaskStatusPending, tasks.TaskStatus(found.Status))
	s.Require().Len(listed, 1)
	s.Assert().Equal(created.ID, listed[0].ID)
}
