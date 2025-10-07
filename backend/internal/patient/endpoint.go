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

func makeListPatientEndpoint(svc PatientService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (response any, err error) {
		req := request.(ListPatientRequest)
		b, err := svc.ListPatients(context.Background(), req.Count, req.Offset)
		if err != nil {
			return nil, err
		}

		return ListPatientResponse{
			Data: b,
		}, nil
	}
}

func makeUpdatePatientEndpoint(svc PatientService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (response any, err error) {
		req := request.(UpdatePatientRequest)
		t, err := svc.UpdatePatient(ctx, req.ID, req.Patient)
		if err != nil {
			return nil, err
		}
		return UpdatePatientResponse{Patient: &t}, nil
	}
}
