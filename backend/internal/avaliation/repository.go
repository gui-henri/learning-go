package avaliation

import (
	"context"

	"github.com/gui-henri/learning-go/db"
)

type AvaliationRepository interface {
	Save(a []byte) error
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
