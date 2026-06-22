package volume

import "sync"

type Service struct {
	mu          sync.RWMutex
	ids         []string
	volumes     map[string]Volume
	volumeTypes []VolumeType
}

func NewService() *Service {
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
	}
}
