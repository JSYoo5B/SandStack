package admin

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	"github.com/JSYoo5B/SandStack/internal/app/fault"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listFaults(w http.ResponseWriter, _ *http.Request) {
	respond.JSON(w, http.StatusOK, faultListResponse{
		Rules: toFaultRuleDocuments(h.faults.List()),
	})
}

func (h Handler) createFault(w http.ResponseWriter, r *http.Request) {
	var request createFaultRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	rule := request.rule()
	h.faults.Add(rule)
	respond.JSON(w, http.StatusCreated, faultResponse{
		Rule: toFaultRuleDocument(rule),
	})
}

func (h Handler) enableFault(w http.ResponseWriter, r *http.Request) {
	h.setFaultEnabled(w, r, true)
}

func (h Handler) disableFault(w http.ResponseWriter, r *http.Request) {
	h.setFaultEnabled(w, r, false)
}

func (h Handler) setFaultEnabled(
	w http.ResponseWriter,
	r *http.Request,
	enabled bool,
) {
	id := chi.URLParam(r, "fault_id")
	var err error
	if enabled {
		err = h.faults.Enable(id)
	} else {
		err = h.faults.Disable(id)
	}

	if errors.Is(err, fault.ErrRuleNotFound) {
		respond.Error(w, http.StatusNotFound, "fault rule not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "fault rule update failed")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
