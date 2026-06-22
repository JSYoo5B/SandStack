package network

type networkListResponse struct {
	Networks []networkDocument `json:"networks"`
}

type networkDocument struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	AdminStateUp bool     `json:"admin_state_up"`
	Status       string   `json:"status"`
	Subnets      []string `json:"subnets"`
	TenantID     string   `json:"tenant_id"`
	ProjectID    string   `json:"project_id"`
	Shared       bool     `json:"shared"`
}
