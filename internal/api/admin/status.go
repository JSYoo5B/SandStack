package admin

import (
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
)

func status(w http.ResponseWriter, r *http.Request) {
	respond.JSON(w, http.StatusOK, statusResponse())
}

func statusResponse() map[string]string {
	return map[string]string{
		"status":  "ok",
		"service": "sandstack",
		"version": "0.1.0",
	}
}
