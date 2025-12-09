package requests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ExportAvaliationRequest struct {
	Id     string
	Format string
}

type ExportAvaliationResponse struct {
	Data []byte
	Err  error
}

func DecodeExportAvaliationRequest(_ context.Context, r *http.Request) (any, error) {
	id := r.PathValue("id")

	return ExportAvaliationRequest{
		Id:     id,
		Format: r.URL.Query().Get("format"), // Get format from ?format=pdf
	}, nil
}

func EncodeExportAvaliationResponse(_ context.Context, w http.ResponseWriter, response any) error {
	resp := response.(ExportAvaliationResponse)

	if resp.Err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusNotFound) // map error types
		json.NewEncoder(w).Encode(map[string]any{
			"error": resp.Err.Error(),
		})
		return nil
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=avaliation.pdf")
	w.Header().Set("Content-Length", fmt.Sprint(len(resp.Data)))

	if _, err := w.Write(resp.Data); err != nil {
		return err
	}

	return nil
}
