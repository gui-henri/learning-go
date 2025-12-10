package requests

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gui-henri/learning-go/internal/expansion/domain"
)

type GetAvaliationFormRequest struct {
	Version int
}

type GetAvaliationFormResponse domain.AvaliationFormOptions

func DecodeGetAvaliationFormRequest(_ context.Context, r *http.Request) (any, error) {
	v := r.Header.Get("If-None-Match")

	vInt, err := strconv.Atoi(v)
	if err != nil {
		vInt = 0
	}

	return GetAvaliationFormRequest{
		Version: vInt,
	}, nil
}

func EncodeGetAvaliationFormResponse(_ context.Context, w http.ResponseWriter, response any) error {
	if response == nil {
		w.WriteHeader(http.StatusNotModified)
		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
