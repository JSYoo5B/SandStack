package compute

type metadataRequest struct {
	Metadata map[string]string `json:"metadata"`
}

type metadataResponse struct {
	Metadata map[string]string `json:"metadata"`
}
