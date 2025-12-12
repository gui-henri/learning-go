package expansion

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/gui-henri/learning-go/internal/expansion/requests"
)

func MakeSaveAvaliationEndpoint(s AvaliationService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req := request.(requests.SaveAvaliationRequest)

		id, err := s.Save(ctx, req.Data)
		if err != nil {
			return requests.SaveAvaliationResponse{Err: err.Error()}, nil
		}

		return requests.SaveAvaliationResponse{Id: id}, nil
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

func MakeGetAvaliationEndpoint(s AvaliationService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (response any, err error) {
		req := request.(requests.GetAvaliationRequest)

		avaliation, err := s.Get(ctx, req.Id)

		return requests.GetAvaliationResponse{
			Data: avaliation,
			Err:  err,
		}, nil
	}
}

func MakeUpdateAvaliationEndpoint(s AvaliationService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req := request.(requests.UpdateAvaliationRequest)

		err := s.Update(ctx, req.Id, req.Data)
		if err != nil {
			return requests.UpdateAvaliationResponse{Status: "failed", Err: err.Error()}, nil
		}

		return requests.UpdateAvaliationResponse{Status: "success"}, nil
	}
}

func MakeGetAvaliationFormOptionsEndpoint(s AvaliationService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (response any, err error) {
		req := request.(requests.GetAvaliationFormRequest)

		opts, ok := s.GetFormOptions(ctx, req.Version)
		if !ok {
			return opts, nil
		}

		return nil, nil
	}
}
