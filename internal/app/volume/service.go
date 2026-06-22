package volume

import (
	"sync"

	"github.com/JSYoo5B/SandStack/internal/platform/clock"
	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
)

type Service struct {
	mu          sync.RWMutex
	ids         []string
	volumes     map[string]Volume
	volumeTypes []VolumeType
	clock       clock.Clock
	idGen       idgen.Generator
}

func NewService() *Service {
	return NewServiceWithRuntime(clock.Wall(), idgen.Random())
}

func NewServiceWithClock(clock clock.Clock) *Service {
	return NewServiceWithRuntime(clock, idgen.Random())
}

func NewServiceWithRuntime(
	clock clock.Clock,
	idGen idgen.Generator,
) *Service {
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
		idGen: idGen,
	}
}
