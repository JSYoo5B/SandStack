package image

type imageListResponse struct {
	Images []imageDocument `json:"images"`
}

type imageDocument struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
