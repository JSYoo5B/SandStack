package compute

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appcompute "github.com/JSYoo5B/SandStack/internal/app/compute"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listAggregates(w http.ResponseWriter, _ *http.Request) {
	respond.JSON(w, http.StatusOK, aggregateListResponse{
		Aggregates: toAggregateDocuments(h.service.ListAggregates()),
	})
}

func (h Handler) createAggregate(w http.ResponseWriter, r *http.Request) {
	var request createAggregateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid aggregate request")
		return
	}

	aggregate := h.service.CreateAggregate(request.createAggregate())

	respond.JSON(w, http.StatusOK, aggregateResponse{
		Aggregate: toAggregateDocument(aggregate),
	})
}

func (h Handler) getAggregate(w http.ResponseWriter, r *http.Request) {
	id, ok := aggregateID(w, r)
	if !ok {
		return
	}

	aggregate, err := h.service.GetAggregate(id)
	if errors.Is(err, appcompute.ErrAggregateNotFound) {
		respond.Error(w, http.StatusNotFound, "aggregate not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "aggregate lookup failed")
		return
	}

	respond.JSON(w, http.StatusOK, aggregateResponse{
		Aggregate: toAggregateDocument(aggregate),
	})
}

func (h Handler) deleteAggregate(w http.ResponseWriter, r *http.Request) {
	id, ok := aggregateID(w, r)
	if !ok {
		return
	}

	err := h.service.DeleteAggregate(id)
	if errors.Is(err, appcompute.ErrAggregateNotFound) {
		respond.Error(w, http.StatusNotFound, "aggregate not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "aggregate delete failed")
		return
	}

	respond.JSON(w, http.StatusOK, map[string]any{})
}

func aggregateID(w http.ResponseWriter, r *http.Request) (int, bool) {
	id, err := strconv.Atoi(chi.URLParam(r, "aggregate_id"))
	if err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid aggregate id")
		return 0, false
	}

	return id, true
}
