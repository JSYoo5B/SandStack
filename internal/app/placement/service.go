package placement

import (
	"errors"

	"github.com/JSYoo5B/SandStack/internal/platform/idgen"
)

var ErrResourceProviderNotFound = errors.New("resource provider not found")

type Service struct {
	resourceProviderRepository ResourceProviderRepository
	idGen                      idgen.Generator
}

func NewServiceWithRepositories(
	resourceProviderRepository ResourceProviderRepository,
	idGen idgen.Generator,
) *Service {
	return &Service{
		resourceProviderRepository: resourceProviderRepository,
		idGen:                      idGen,
	}
}

func (s *Service) CreateResourceProvider(
	input CreateResourceProvider,
) ResourceProvider {
	uuid := input.UUID
	if uuid == "" {
		uuid = "rp-" + s.idGen.Hex(16)
	}

	rootProviderUUID := uuid
	if input.ParentProviderUUID != "" {
		rootProviderUUID = input.ParentProviderUUID
	}

	provider := ResourceProvider{
		UUID:               uuid,
		Name:               input.Name,
		Generation:         0,
		ParentProviderUUID: input.ParentProviderUUID,
		RootProviderUUID:   rootProviderUUID,
	}

	return s.resourceProviderRepository.Create(provider)
}

func (s *Service) ListResourceProviders() []ResourceProvider {
	return s.resourceProviderRepository.List()
}

func (s *Service) GetResourceProvider(uuid string) (ResourceProvider, error) {
	return s.resourceProviderRepository.Get(uuid)
}

func (s *Service) DeleteResourceProvider(uuid string) error {
	return s.resourceProviderRepository.Delete(uuid)
}

func (s *Service) Reset() {
	s.resourceProviderRepository.Reset()
}
