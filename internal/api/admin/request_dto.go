package admin

import "github.com/JSYoo5B/SandStack/internal/app/requestlog"

type requestListResponse struct {
	Requests []requestDocument `json:"requests"`
}

type requestDocument struct {
	ID     string `json:"id"`
	Method string `json:"method"`
	Path   string `json:"path"`
	Status int    `json:"status"`
}

func toRequestDocuments(records []requestlog.Record) []requestDocument {
	documents := make([]requestDocument, 0, len(records))
	for _, record := range records {
		documents = append(documents, requestDocument{
			ID:     record.ID,
			Method: record.Method,
			Path:   record.Path,
			Status: record.Status,
		})
	}

	return documents
}
