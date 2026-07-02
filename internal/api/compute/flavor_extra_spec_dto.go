package compute

type flavorExtraSpecListResponse struct {
	ExtraSpecs map[string]string `json:"extra_specs"`
}

type createFlavorExtraSpecsRequest struct {
	ExtraSpecs map[string]string `json:"extra_specs"`
}
