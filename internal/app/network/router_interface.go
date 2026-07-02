package network

import "errors"

var ErrRouterInterfaceNotFound = errors.New("router interface not found")

func (s *Service) AddRouterInterface(
	routerID string,
	request RouterInterfaceRequest,
) (RouterInterface, error) {
	router, err := s.routerRepository.Get(routerID)
	if err != nil {
		return RouterInterface{}, err
	}

	portID := request.PortID
	subnetID := request.SubnetID
	if portID == "" {
		portID = "port-" + s.idGen.Hex(16)
	}

	routerInterface := RouterInterface{
		ID:       "ri-" + s.idGen.Hex(16),
		RouterID: routerID,
		SubnetID: subnetID,
		PortID:   portID,
		TenantID: router.ProjectID,
	}

	return s.routerInterfaceRepository.Create(routerInterface), nil
}

func (s *Service) RemoveRouterInterface(
	routerID string,
	request RouterInterfaceRequest,
) (RouterInterface, error) {
	if _, err := s.routerRepository.Get(routerID); err != nil {
		return RouterInterface{}, err
	}

	routerInterface, err := s.routerInterfaceRepository.Find(routerID, request)
	if err != nil {
		return RouterInterface{}, err
	}

	if err := s.routerInterfaceRepository.Delete(routerInterface.ID); err != nil {
		return RouterInterface{}, err
	}

	return routerInterface, nil
}
