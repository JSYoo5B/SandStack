package network

type versionsResponse struct {
	Versions []versionDocument `json:"versions"`
}

type versionDocument struct {
	ID     string        `json:"id"`
	Status string        `json:"status"`
	Links  []versionLink `json:"links"`
}

type versionLink struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}
