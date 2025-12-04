package avaliation

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type SaveAvaliationRequest struct {
	Data AvaliacaoRequest
}

type SaveAvaliationResponse struct {
	Status string `json:"status"`
	Err    string `json:"error,omitempty"`
}

func decodeSaveFormRequest(_ context.Context, r *http.Request) (any, error) {
	var body AvaliacaoRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Println("JSON Error:", err)
		return nil, err
	}

	return SaveAvaliationRequest{Data: body}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response any) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type ExportAvaliationRequest struct {
	Id     int
	Format string
}

type ExportAvaliationResponse struct {
	Data []byte
	Err  error
}

func decodeExportAvaliationRequest(_ context.Context, r *http.Request) (any, error) {
	id := r.PathValue("id")

	// Parse ID from string to int...
	idInt, err := strconv.Atoi(id)

	if err != nil {
		return nil, fmt.Errorf("invalid id format: %s", id)
	}

	return ExportAvaliationRequest{
		Id:     idInt,
		Format: r.URL.Query().Get("format"), // Get format from ?format=pdf
	}, nil
}

func encodeExportAvaliationResponse(_ context.Context, w http.ResponseWriter, response any) error {
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
