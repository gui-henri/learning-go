package requests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gui-henri/learning-go/internal/expansion/domain"
)

type SaveAvaliationRequest struct {
	Data domain.AvaliacaoRequest
}

type SaveAvaliationResponse struct {
	Id  string `json:"id"`
	Err string `json:"error,omitempty"`
}

func DecodeSaveAvaliationRequest(_ context.Context, r *http.Request) (any, error) {
	var body domain.AvaliacaoRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Println("JSON Error:", err)
		return nil, err
	}

	return SaveAvaliationRequest{Data: body}, nil
}

func EncodeSaveAvaliationResponse(_ context.Context, w http.ResponseWriter, response any) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
