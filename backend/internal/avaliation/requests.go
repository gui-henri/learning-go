package avaliation

import (
	"context"
	"encoding/json"
	"net/http"
)

type SaveAvaliationRequest struct {
	Data map[string]interface{}
}

type SaveAvaliationResponse struct {
	Status string `json:"status"`
	Err    string `json:"error,omitempty"`
}

func decodeSaveFormRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return SaveAvaliationRequest{Data: body}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
