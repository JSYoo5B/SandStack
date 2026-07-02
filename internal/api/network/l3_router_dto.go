package network

import appnetwork "github.com/JSYoo5B/SandStack/internal/app/network"

type createRouterRequest struct {
	Router createRouterDocument `json:"router"`
}

type createRouterDocument struct {
	Name                  string                `json:"name"`
	Description           string                `json:"description"`
	AdminStateUp          *bool                 `json:"admin_state_up"`
	Distributed           *bool                 `json:"distributed"`
	ProjectID             string                `json:"project_id"`
	TenantID              string                `json:"tenant_id"`
	GatewayInfo           routerGatewayDocument `json:"external_gateway_info"`
	AvailabilityZoneHints []string              `json:"availability_zone_hints"`
}

func (r createRouterRequest) createRouter() appnetwork.CreateRouter {
	projectID := r.Router.ProjectID
	if projectID == "" {
		projectID = r.Router.TenantID
	}

	return appnetwork.CreateRouter{
		Name:                  r.Router.Name,
		Description:           r.Router.Description,
		AdminStateUp:          r.Router.AdminStateUp,
		Distributed:           r.Router.Distributed,
		ProjectID:             projectID,
		GatewayInfo:           toAppRouterGateway(r.Router.GatewayInfo),
		AvailabilityZoneHints: r.Router.AvailabilityZoneHints,
	}
}

type routerListResponse struct {
	Routers []routerDocument `json:"routers"`
}

type routerResponse struct {
	Router routerDocument `json:"router"`
}

type routerDocument struct {
	ID                    string                `json:"id"`
	Name                  string                `json:"name"`
	Description           string                `json:"description"`
	AdminStateUp          bool                  `json:"admin_state_up"`
	Distributed           bool                  `json:"distributed"`
	Status                string                `json:"status"`
	TenantID              string                `json:"tenant_id"`
	ProjectID             string                `json:"project_id"`
	GatewayInfo           routerGatewayDocument `json:"external_gateway_info"`
	Routes                []routeDocument       `json:"routes"`
	AvailabilityZoneHints []string              `json:"availability_zone_hints"`
	Tags                  []string              `json:"tags"`
}

type routerGatewayDocument struct {
	NetworkID        string                    `json:"network_id,omitempty"`
	EnableSNAT       *bool                     `json:"enable_snat,omitempty"`
	ExternalFixedIPs []externalFixedIPDocument `json:"external_fixed_ips,omitempty"`
	QoSPolicyID      string                    `json:"qos_policy_id,omitempty"`
}

type externalFixedIPDocument struct {
	IPAddress string `json:"ip_address,omitempty"`
	SubnetID  string `json:"subnet_id,omitempty"`
}

type routeDocument struct {
	NextHop         string `json:"nexthop"`
	DestinationCIDR string `json:"destination"`
}

func toRouterDocuments(routers []appnetwork.Router) []routerDocument {
	documents := make([]routerDocument, 0, len(routers))
	for _, router := range routers {
		documents = append(documents, toRouterDocument(router))
	}

	return documents
}

func toRouterDocument(router appnetwork.Router) routerDocument {
	return routerDocument{
		ID:                    router.ID,
		Name:                  router.Name,
		Description:           router.Description,
		AdminStateUp:          router.AdminStateUp,
		Distributed:           router.Distributed,
		Status:                router.Status,
		TenantID:              router.TenantID,
		ProjectID:             router.ProjectID,
		GatewayInfo:           toRouterGatewayDocument(router.GatewayInfo),
		Routes:                toRouteDocuments(router.Routes),
		AvailabilityZoneHints: router.AvailabilityZoneHints,
		Tags:                  router.Tags,
	}
}

func toAppRouterGateway(
	gateway routerGatewayDocument,
) appnetwork.RouterGatewayInfo {
	return appnetwork.RouterGatewayInfo{
		NetworkID:        gateway.NetworkID,
		EnableSNAT:       gateway.EnableSNAT,
		ExternalFixedIPs: toAppExternalFixedIPs(gateway.ExternalFixedIPs),
		QoSPolicyID:      gateway.QoSPolicyID,
	}
}

func toRouterGatewayDocument(
	gateway appnetwork.RouterGatewayInfo,
) routerGatewayDocument {
	return routerGatewayDocument{
		NetworkID:        gateway.NetworkID,
		EnableSNAT:       gateway.EnableSNAT,
		ExternalFixedIPs: toExternalFixedIPDocuments(gateway.ExternalFixedIPs),
		QoSPolicyID:      gateway.QoSPolicyID,
	}
}

func toAppExternalFixedIPs(
	fixedIPs []externalFixedIPDocument,
) []appnetwork.ExternalFixedIP {
	values := make([]appnetwork.ExternalFixedIP, 0, len(fixedIPs))
	for _, fixedIP := range fixedIPs {
		values = append(values, appnetwork.ExternalFixedIP{
			IPAddress: fixedIP.IPAddress,
			SubnetID:  fixedIP.SubnetID,
		})
	}

	return values
}

func toExternalFixedIPDocuments(
	fixedIPs []appnetwork.ExternalFixedIP,
) []externalFixedIPDocument {
	documents := make([]externalFixedIPDocument, 0, len(fixedIPs))
	for _, fixedIP := range fixedIPs {
		documents = append(documents, externalFixedIPDocument{
			IPAddress: fixedIP.IPAddress,
			SubnetID:  fixedIP.SubnetID,
		})
	}

	return documents
}

func toRouteDocuments(routes []appnetwork.Route) []routeDocument {
	documents := make([]routeDocument, 0, len(routes))
	for _, route := range routes {
		documents = append(documents, routeDocument{
			NextHop:         route.NextHop,
			DestinationCIDR: route.DestinationCIDR,
		})
	}

	return documents
}
