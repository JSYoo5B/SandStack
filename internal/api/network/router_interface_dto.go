package network

import appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"

type routerInterfaceRequest struct {
	SubnetID string `json:"subnet_id"`
	PortID   string `json:"port_id"`
}

func (r routerInterfaceRequest) routerInterfaceRequest() appnetwork.RouterInterfaceRequest {
	return appnetwork.RouterInterfaceRequest{
		SubnetID: r.SubnetID,
		PortID:   r.PortID,
	}
}

type routerInterfaceDocument struct {
	ID       string `json:"id"`
	SubnetID string `json:"subnet_id"`
	PortID   string `json:"port_id"`
	TenantID string `json:"tenant_id"`
}

func toRouterInterfaceDocument(
	routerInterface appnetwork.RouterInterface,
) routerInterfaceDocument {
	return routerInterfaceDocument{
		ID:       routerInterface.ID,
		SubnetID: routerInterface.SubnetID,
		PortID:   routerInterface.PortID,
		TenantID: routerInterface.TenantID,
	}
}
