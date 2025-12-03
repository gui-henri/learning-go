package avaliation

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gui-henri/learning-go/db"
)

type AvaliationRepository interface {
	Save(a []byte) error
	GetOne(id int) (AvaliacaoRequest, error)
}

type avaliationRepository struct {
	db db.IDB
}

func NewAvaliationRepository(db db.IDB) *avaliationRepository {
	return &avaliationRepository{
		db: db,
	}
}

func (s *avaliationRepository) Save(a []byte) error {
	query := "INSERT INTO avaliation (resource_json) VALUES ($1)"
	_, err := s.db.Exec(context.Background(), query, a)
	return err
}

func (s *avaliationRepository) GetOne(id int) (AvaliacaoRequest, error) {
	query := "SELECT resource_json FROM avaliation WHERE id=$1"
	row := s.db.QueryRow(context.Background(), query, id)

	var schm AvaliacaoSchema

	err := row.Scan(&schm.ResourceJSON)

	if err != nil {
		return AvaliacaoRequest{}, err
	}

	var result AvaliacaoRequest

	if err := json.Unmarshal(schm.ResourceJSON, &result); err != nil {
		return AvaliacaoRequest{}, fmt.Errorf("failed to parse json: %w", err)
	}

	return result, nil
}
