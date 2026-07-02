package compute

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listKeyPairs(w http.ResponseWriter, _ *http.Request) {
	respond.JSON(w, http.StatusOK, keyPairListResponse{
		KeyPairs: toKeyPairListDocuments(h.service.ListKeyPairs()),
	})
}

func (h Handler) createKeyPair(w http.ResponseWriter, r *http.Request) {
	var request createKeyPairRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	keyPair := h.service.CreateKeyPair(request.createKeyPair())
	respond.JSON(w, http.StatusCreated, keyPairResponse{
		KeyPair: toKeyPairDocument(keyPair),
	})
}

func (h Handler) getKeyPair(w http.ResponseWriter, r *http.Request) {
	keyPair, err := h.service.GetKeyPair(chi.URLParam(r, "keypair_name"))
	if errors.Is(err, appcompute.ErrKeyPairNotFound) {
		respond.Error(w, http.StatusNotFound, "key pair not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "key pair lookup failed")
		return
	}

	respond.JSON(w, http.StatusOK, keyPairResponse{
		KeyPair: toKeyPairDocument(keyPair),
	})
}

func (h Handler) deleteKeyPair(w http.ResponseWriter, r *http.Request) {
	err := h.service.DeleteKeyPair(chi.URLParam(r, "keypair_name"))
	if errors.Is(err, appcompute.ErrKeyPairNotFound) {
		respond.Error(w, http.StatusNotFound, "key pair not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "key pair delete failed")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
