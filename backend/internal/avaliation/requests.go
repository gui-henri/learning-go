package avaliation

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type SaveAvaliationRequest struct {
	Data entity.AvaliacaoRequest
}

type SaveAvaliationResponse struct {
	Status string `json:"status"`
	Err    string `json:"error,omitempty"`
}

func decodeSaveFormRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body AvaliacaoRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	fmt.Println(body)

	return SaveAvaliationRequest{Data: body}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
