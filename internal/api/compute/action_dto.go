package compute

import "encoding/json"

type serverActionRequest map[string]json.RawMessage

func (r serverActionRequest) Has(action string) bool {
	_, ok := r[action]
	return ok
}
