package image

type importInfoResponse struct {
	ImportMethods importMethodsDocument `json:"import-methods"`
}

type importMethodsDocument struct {
	Description string   `json:"description"`
	Type        string   `json:"type"`
	Value       []string `json:"value"`
}
