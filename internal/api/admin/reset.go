package admin

import "net/http"

func (h Handler) resetState(w http.ResponseWriter, _ *http.Request) {
	h.reset()
	w.WriteHeader(http.StatusNoContent)
}
