package compute

type metadataRequest struct {
	Metadata map[string]string `json:"metadata"`
}

type metadataResponse struct {
	Metadata map[string]string `json:"metadata"`
}

type metadatumRequest struct {
	Meta map[string]string `json:"meta"`
}

type metadatumResponse struct {
	Meta map[string]string `json:"meta"`
}
