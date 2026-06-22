package compute

import appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"

type flavorListResponse struct {
	Flavors []flavorDocument `json:"flavors"`
}

type flavorResponse struct {
	Flavor flavorDocument `json:"flavor"`
}

type flavorDocument struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	RAM         int               `json:"ram"`
	VCPUs       int               `json:"vcpus"`
	Disk        int               `json:"disk"`
	Swap        int               `json:"swap"`
	RxTxFactor  float64           `json:"rxtx_factor"`
	IsPublic    bool              `json:"os-flavor-access:is_public"`
	Ephemeral   int               `json:"OS-FLV-EXT-DATA:ephemeral"`
	Description string            `json:"description"`
	ExtraSpecs  map[string]string `json:"extra_specs"`
}

func toFlavorDocuments(flavors []appcompute.Flavor) []flavorDocument {
	documents := make([]flavorDocument, 0, len(flavors))
	for _, flavor := range flavors {
		documents = append(documents, toFlavorDocument(flavor))
	}

	return documents
}

func toFlavorDocument(flavor appcompute.Flavor) flavorDocument {
	return flavorDocument{
		ID:          flavor.ID,
		Name:        flavor.Name,
		RAM:         flavor.RAM,
		VCPUs:       flavor.VCPUs,
		Disk:        flavor.Disk,
		Swap:        flavor.Swap,
		RxTxFactor:  flavor.RxTxFactor,
		IsPublic:    flavor.IsPublic,
		Ephemeral:   flavor.Ephemeral,
		Description: flavor.Description,
		ExtraSpecs:  flavor.ExtraSpecs,
	}
}
