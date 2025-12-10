package requests

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gui-henri/learning-go/internal/expansion/domain"
)

type GetAvaliationFormResponse domain.AvaliationFormOptions

func EncodeGetAvaliationFormResponse(_ context.Context, w http.ResponseWriter, response any) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
