package volume

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/JSYoo5B/SandStack/internal/api/respond"
	appvolume "github.com/JSYoo5B/SandStack/internal/app/volume"
	"github.com/go-chi/chi/v5"
)

func (h Handler) listTransfers(w http.ResponseWriter, _ *http.Request) {
	respond.JSON(w, http.StatusOK, transferListResponse{
		Transfers: toTransferDocuments(h.service.ListTransfers()),
	})
}

func (h Handler) createTransfer(w http.ResponseWriter, r *http.Request) {
	var request createTransferRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	transfer := h.service.CreateTransfer(request.createTransfer())
	respond.JSON(w, http.StatusAccepted, transferResponse{
		Transfer: toTransferDocument(transfer),
	})
}

func (h Handler) getTransfer(w http.ResponseWriter, r *http.Request) {
	transfer, err := h.service.GetTransfer(chi.URLParam(r, "transfer_id"))
	if errors.Is(err, appvolume.ErrTransferNotFound) {
		respond.Error(w, http.StatusNotFound, "transfer not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "transfer lookup failed")
		return
	}

	respond.JSON(w, http.StatusOK, transferResponse{
		Transfer: toTransferDocument(transfer),
	})
}

func (h Handler) deleteTransfer(w http.ResponseWriter, r *http.Request) {
	err := h.service.DeleteTransfer(chi.URLParam(r, "transfer_id"))
	if errors.Is(err, appvolume.ErrTransferNotFound) {
		respond.Error(w, http.StatusNotFound, "transfer not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "transfer delete failed")
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (h Handler) acceptTransfer(w http.ResponseWriter, r *http.Request) {
	var request acceptTransferRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		respond.Error(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	transfer, err := h.service.AcceptTransfer(
		chi.URLParam(r, "transfer_id"),
		request.Accept.AuthKey,
	)
	if errors.Is(err, appvolume.ErrTransferNotFound) {
		respond.Error(w, http.StatusNotFound, "transfer not found")
		return
	}
	if err != nil {
		respond.Error(w, http.StatusInternalServerError, "transfer accept failed")
		return
	}

	respond.JSON(w, http.StatusAccepted, transferResponse{
		Transfer: toTransferDocument(transfer),
	})
}
