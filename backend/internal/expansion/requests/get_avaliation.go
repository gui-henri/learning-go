package requests

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gui-henri/learning-go/internal/expansion/domain"
)

type GetAvaliationRequest struct {
	Id string
}

type GetAvaliationResponse struct {
	Data domain.AvaliacaoRequest
	Err  error
}

func DecodeGetAvaliationRequest(_ context.Context, r *http.Request) (any, error) {
	id := r.PathValue("id")

	return GetAvaliationRequest{
		Id: id,
	}, nil
}

func EncodeGetAvaliationRequestResponse(_ context.Context, w http.ResponseWriter, response any) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
