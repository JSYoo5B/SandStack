package image_test

import (
	"testing"
	"time"

	"github.com/JSYoo5B/SandStack/internal/app/image"
	"github.com/JSYoo5B/SandStack/internal/platform/clock"
	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
	storeimage "github.com/JSYoo5B/SandStack/internal/store/image"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite
}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) TestCreateImageUsesInjectedClock() {
	now := time.Date(2026, 6, 23, 8, 30, 0, 0, time.UTC)
	service := image.NewServiceWithRuntime(
		storeimage.NewMemoryRepository(),
		storeimage.NewMemoryDataRepository(),
		storeimage.NewMemoryMemberRepository(),
		storeimage.NewMemoryTaskRepository(),
		clock.Fixed(now),
		idgen.Random(),
	)

	created := service.Create(image.CreateImage{
		Name:            "ubuntu",
		ContainerFormat: "bare",
		DiskFormat:      "qcow2",
	})

	s.Assert().Equal("2026-06-23T08:30:00Z", created.CreatedAt)
	s.Assert().Equal("2026-06-23T08:30:00Z", created.UpdatedAt)
}

func (s *ServiceSuite) TestCreateImageUsesInjectedIDGenerator() {
	service := image.NewServiceWithRuntime(
		storeimage.NewMemoryRepository(),
		storeimage.NewMemoryDataRepository(),
		storeimage.NewMemoryMemberRepository(),
		storeimage.NewMemoryTaskRepository(),
		clock.Fixed(time.Time{}),
		idgen.Fixed("image-id"),
	)

	created := service.Create(image.CreateImage{
		Name:            "ubuntu",
		ContainerFormat: "bare",
		DiskFormat:      "qcow2",
	})

	s.Assert().Equal("img-image-id", created.ID)
}

func (s *ServiceSuite) TestResetClearsImages() {
	service := image.NewServiceWithRuntime(
		storeimage.NewMemoryRepository(),
		storeimage.NewMemoryDataRepository(),
		storeimage.NewMemoryMemberRepository(),
		storeimage.NewMemoryTaskRepository(),
		clock.Fixed(time.Time{}),
		idgen.Fixed("image-id"),
	)
	created := service.Create(image.CreateImage{
		Name:            "ubuntu",
		ContainerFormat: "bare",
		DiskFormat:      "qcow2",
	})
	err := service.UploadData(created.ID, []byte("image-data"))
	s.Require().NoError(err)
	_, err = service.CreateMember(created.ID, "project-1")
	s.Require().NoError(err)
	task := service.CreateTask(image.CreateTask{Type: "import"})

	service.Reset()

	s.Assert().Empty(service.List())
	_, err = service.DownloadData(created.ID)
	s.ErrorIs(err, image.ErrImageNotFound)
	members, err := service.ListMembers(created.ID)
	s.ErrorIs(err, image.ErrImageNotFound)
	s.Nil(members)
	_, err = service.GetTask(task.ID)
	s.ErrorIs(err, image.ErrTaskNotFound)
}
