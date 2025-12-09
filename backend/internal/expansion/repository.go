package expansion

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gui-henri/learning-go/db"
	"github.com/gui-henri/learning-go/internal/expansion/domain"
)

type AvaliationRepository interface {
	Save(a []byte) error
	GetOne(id string) (domain.AvaliacaoRequest, error)
	GetAll() ([]domain.AvaliacaoRequest, error)
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
func (s *avaliationRepository) GetOne(id string) (domain.AvaliacaoRequest, error) {
	query := "SELECT resource_json FROM avaliation WHERE id = $1"
	row := s.db.QueryRow(context.Background(), query, id)

	var schm domain.AvaliacaoSchema

	err := row.Scan(&schm.ResourceJSON)
	if err != nil {
		return domain.AvaliacaoRequest{}, err
	}

	var result domain.AvaliacaoRequest
	if err := json.Unmarshal(schm.ResourceJSON, &result); err != nil {
		return domain.AvaliacaoRequest{}, fmt.Errorf("failed to parse json: %w", err)
	}

	return result, nil
}

func (s *avaliationRepository) GetAll() ([]domain.AvaliacaoRequest, error) {
	query := "SELECT id, resource_json, created_at, last_updated, active FROM avaliation"
	rows, err := s.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var avaliations []domain.AvaliacaoRequest
	for rows.Next() {
		var schm domain.AvaliacaoSchema
		if err := rows.Scan(&schm.Id, &schm.ResourceJSON, &schm.CreatedAt, &schm.LastUpdated, &schm.Active); err != nil {
			return nil, err
		}

		var result domain.AvaliacaoRequest
		if err := json.Unmarshal(schm.ResourceJSON, &result); err != nil {
			return nil, fmt.Errorf("failed to parse json: %w", err)
		}
		avaliations = append(avaliations, result)
	}

	return avaliations, nil
}
