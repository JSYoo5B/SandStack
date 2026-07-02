package network

import "errors"

var ErrRouterNotFound = errors.New("router not found")

func (s *Service) CreateRouter(input CreateRouter) Router {
	adminStateUp := true
	if input.AdminStateUp != nil {
		adminStateUp = *input.AdminStateUp
	}

	distributed := false
	if input.Distributed != nil {
		distributed = *input.Distributed
	}

	router := Router{
		ID:                    "router-" + s.idGen.Hex(16),
		Name:                  input.Name,
		Description:           input.Description,
		AdminStateUp:          adminStateUp,
		Distributed:           distributed,
		Status:                "ACTIVE",
		TenantID:              input.ProjectID,
		ProjectID:             input.ProjectID,
		GatewayInfo:           input.GatewayInfo,
		Routes:                []Route{},
		AvailabilityZoneHints: input.AvailabilityZoneHints,
		Tags:                  []string{},
	}

	return s.routerRepository.Create(router)
}

func (s *Service) ListRouters() []Router {
	return s.routerRepository.List()
}

func (s *Service) GetRouter(id string) (Router, error) {
	return s.routerRepository.Get(id)
}

func (s *Service) DeleteRouter(id string) error {
	return s.routerRepository.Delete(id)
}
