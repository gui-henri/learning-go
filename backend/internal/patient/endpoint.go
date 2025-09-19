package patient

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeInsertPatientEndpoint(svc PatientService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (response any, err error) {
		req := request.(InsertPatientRequest)
		t, err := svc.InsertPatient(ctx, req.Patient)
		if err != nil {
			return nil, err
		}
		return InsertPatientResponse{Patient: &t}, nil
	}
}
