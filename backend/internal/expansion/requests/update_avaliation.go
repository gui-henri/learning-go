package requests

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gui-henri/learning-go/internal/expansion/domain"
)

type UpdateAvaliationRequest struct {
	Id   string
	Data domain.AvaliacaoRequest
}

type UpdateAvaliationResponse struct {
	Status string `json:"status"`
	Err    string `json:"error,omitempty"`
}

func DecodeUpdateAvaliationRequest(_ context.Context, r *http.Request) (any, error) {
	id := r.PathValue("id")

	var body domain.AvaliacaoRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return UpdateAvaliationRequest{Id: id, Data: body}, nil
}

func EncodeUpdateAvaliationResponse(_ context.Context, w http.ResponseWriter, response any) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
