package patient

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gui-henri/learning-go/pkg/fhir"
)

type InsertPatientRequest struct {
	Patient fhir.Patient `json:"paciente"`
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
