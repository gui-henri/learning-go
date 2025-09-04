package tasks

import (
	"context"

	apperrors "github.com/gui-henri/learning-go/pkg/errors"
)

type TaskService interface {
	GetTask(ctx context.Context, id int) (*Tarefa, error)
	InsertTask(ctx context.Context, descricao string, prazo string) (int, error)
	FinishTask(ctx context.Context, id int) (*Tarefa, error)
	GetAllIncompleteTasks(ctx context.Context) []Tarefa
}

type service struct {
	repository taskRepository
}

func NewService(r taskRepository) *service {
	return &service{
		repository: r,
	}
}

func (s *service) GetTask(ctx context.Context, id int) (*Tarefa, error) {
	return s.repository.GetTask(id)
}

func (s *service) InsertTask(ctx context.Context, descricao string, prazo string) (int, error) {
	return s.repository.InsertTask(descricao, prazo)
}

func (s *service) FinishTask(ctx context.Context, id int) (*Tarefa, error) {
	t, err := s.repository.GetTask(id)
	if err != nil {
		return &Tarefa{}, err
	}

	if t.Concluida {
		return &Tarefa{}, apperrors.AlreadyFinished
	}

	t.Concluida = true

	return t, nil

}

func (s *service) GetAllIncompleteTasks(ctx context.Context) []Tarefa {
	return s.repository.GetAllIncomplete()
}
