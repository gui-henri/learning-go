package transport_encoding

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	apperrors "github.com/gui-henri/learning-go/pkg/errors"
)

func EncodeRequest[T any](_ context.Context, r *http.Request) (any, error) {
	var request T
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response any) error {
	return json.NewEncoder(w).Encode(response)
}

func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var syntaxError *json.SyntaxError
	switch {
	case errors.Is(err, io.EOF):
		w.WriteHeader(http.StatusBadRequest)
	case errors.As(err, &syntaxError):
		w.WriteHeader(http.StatusBadRequest)
	case errors.Is(err, apperrors.InvalidInput):
		w.WriteHeader(http.StatusBadRequest)
	case errors.Is(err, apperrors.NotFound):
		w.WriteHeader(http.StatusNotFound)
	case errors.Is(err, apperrors.AlreadyFinished):
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
