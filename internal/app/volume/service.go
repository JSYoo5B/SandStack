package volume

import (
	"sync"

	"github.com/JSYoo5B/SandStack/internal/platform/clock"
)

type Service struct {
	mu          sync.RWMutex
	ids         []string
	volumes     map[string]Volume
	volumeTypes []VolumeType
	clock       clock.Clock
}

func NewService() *Service {
	return NewServiceWithClock(clock.Wall())
}

func NewServiceWithClock(clock clock.Clock) *Service {
	return &Service{
		ids:     []string{},
		volumes: map[string]Volume{},
		volumeTypes: []VolumeType{
			{
				ID:          "default",
				Name:        "__DEFAULT__",
				Description: "Default test volume type",
				ExtraSpecs:  map[string]string{},
				IsPublic:    true,
			},
		},
		clock: clock,
	}
}
