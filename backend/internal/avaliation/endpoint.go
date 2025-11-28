package avaliation

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func MakeSaveFormEndpoint(s avaliationService) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req := request.(SaveAvaliationRequest)

		err := s.Save(ctx, req.Data)
		if err != nil {
			return SaveAvaliationResponse{Status: "failed", Err: err.Error()}, nil
		}

		return SaveAvaliationResponse{Status: "success"}, nil
	}
}
