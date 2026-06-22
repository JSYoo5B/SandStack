package volume

type volumeListResponse struct {
	Volumes []volumeDocument `json:"volumes"`
}

type volumeDocument struct {
	ID          string            `json:"id"`
	Status      string            `json:"status"`
	Size        int               `json:"size"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Metadata    map[string]string `json:"metadata"`
	Bootable    string            `json:"bootable"`
	Encrypted   bool              `json:"encrypted"`
	Multiattach bool              `json:"multiattach"`
}
