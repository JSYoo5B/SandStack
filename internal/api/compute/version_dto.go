package compute

type versionResponse struct {
	Version versionDocument `json:"version"`
}

type versionDocument struct {
	ID      string        `json:"id"`
	Status  string        `json:"status"`
	Version string        `json:"version"`
	Links   []versionLink `json:"links"`
}

type versionLink struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}
