package patient

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gui-henri/learning-go/pkg/fhir"
)

type InsertPatientRequest struct {
	Patient fhir.Patient `json:"paciente"`
}

type GetPatientRequest struct {
	ID string `param:"id"`
}

type GetPatientResponse = fhir.Patient

func decodeGetPatientRequest(ctx context.Context, r *http.Request) (request any, err error) {
	var req GetPatientRequest
	req.ID = r.PathValue("id")
	return req, nil
}

type InsertPatientResponse struct {
	Patient *paciente `json:"paciente"`
}

type ListPatientResponse struct {
	Data fhir.Bundle `json:"data"`
}

type ListPatientRequest struct {
	Count  int `query:"_count"`
	Offset int `query:"_offset"`
}

func decodeListPatientRequest(_ context.Context, r *http.Request) (any, error) {
	var req ListPatientRequest
	queryValues := r.URL.Query()

	countStr := queryValues.Get("_count")
	if countStr != "" {
		count, err := strconv.Atoi(countStr)
		if err != nil {
			return nil, err
		}
		req.Count = count
	} else {
		req.Count = 10
	}

	offsetStr := queryValues.Get("_offset")
	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			return nil, err
		}
		req.Offset = offset
	} else {
		req.Offset = 0
	}
	return req, nil

}

type UpdatePatientRequest struct {
	ID      string       `param:"id"`
	Patient fhir.Patient `json:"paciente"`
}

type UpdatePatientResponse struct {
	Patient *paciente `json:"paciente"`
}

func decodeUpdatePatientRequest(ctx context.Context, r *http.Request) (request any, err error) {
	var req UpdatePatientRequest

	var p fhir.Patient
	if r.Body != nil && r.ContentLength > 0 {
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil && err != io.EOF {
			return nil, err
		}
	}

	req.ID = r.PathValue("id")
	req.Patient = p

	return req, nil
}
