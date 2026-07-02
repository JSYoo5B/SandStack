package compute

import "encoding/json"

type serverActionRequest map[string]json.RawMessage

func (r serverActionRequest) Has(action string) bool {
	_, ok := r[action]
	return ok
}

func (r serverActionRequest) SecurityGroupName(action string) string {
	var body struct {
		Name string `json:"name"`
	}
	_ = json.Unmarshal(r[action], &body)
	return body.Name
}
