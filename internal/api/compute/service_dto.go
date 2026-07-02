package compute

import appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"

type computeServiceListResponse struct {
	Services []computeServiceDocument `json:"services"`
}

type computeServiceDocument struct {
	ID             string `json:"id"`
	Binary         string `json:"binary"`
	DisabledReason string `json:"disabled_reason"`
	ForcedDown     bool   `json:"forced_down"`
	Host           string `json:"host"`
	State          string `json:"state"`
	Status         string `json:"status"`
	UpdatedAt      string `json:"updated_at"`
	Zone           string `json:"zone"`
}

func toComputeServiceDocuments(
	services []appcompute.ComputeService,
) []computeServiceDocument {
	documents := make([]computeServiceDocument, 0, len(services))
	for _, service := range services {
		documents = append(documents, toComputeServiceDocument(service))
	}

	return documents
}

func toComputeServiceDocument(
	service appcompute.ComputeService,
) computeServiceDocument {
	return computeServiceDocument{
		ID:             service.ID,
		Binary:         service.Binary,
		DisabledReason: service.DisabledReason,
		ForcedDown:     service.ForcedDown,
		Host:           service.Host,
		State:          service.State,
		Status:         service.Status,
		UpdatedAt:      service.UpdatedAt,
		Zone:           service.Zone,
	}
}
