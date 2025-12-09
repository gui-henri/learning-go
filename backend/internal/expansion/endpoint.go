package expansion

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/gui-henri/learning-go/internal/expansion/requests"
)

func MakeSaveAvaliationEndpoint(s AvaliationService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req := request.(requests.SaveAvaliationRequest)

		err := s.Save(ctx, req.Data)
		if err != nil {
			return requests.SaveAvaliationResponse{Status: "failed", Err: err.Error()}, nil
		}

		return requests.SaveAvaliationResponse{Status: "success"}, nil
	}
}

func MakeListAvaliationEndpoint(s AvaliationService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		avaliations, err := s.List(ctx)
		if err != nil {
			return requests.ListAvaliationResponse{Err: err.Error()}, nil
		}

		return requests.ListAvaliationResponse{Avaliations: avaliations}, nil
	}
}

func MakeExportAvaliationEndpoint(s AvaliationService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (response any, err error) {
		req := request.(requests.ExportAvaliationRequest)

		pdf, err := s.Export(ctx, req.Id, req.Format)

		return requests.ExportAvaliationResponse{
			Data: pdf,
			Err:  err,
		}, nil
	}
}
