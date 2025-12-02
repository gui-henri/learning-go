package avaliation

import (
	"context"
	"encoding/json"
)

type AvaliationService interface {
	Save(ctx context.Context, data AvaliacaoRequest) error
}

type avaliationService struct {
	repository avaliationRepository
}

func NewAvaliationService(r avaliationRepository) *avaliationService {
	return &avaliationService{repository: r}
}

func (s *avaliationService) Save(ctx context.Context, data AvaliacaoRequest) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = s.repository.Save(dataBytes)

	return err
}
