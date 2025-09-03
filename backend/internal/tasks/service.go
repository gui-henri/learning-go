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

type service struct{}

func NewService() *service {
	return &service{}
}

var repository = NewTaskRepository()

func (s *service) GetTask(ctx context.Context, id int) (*Tarefa, error) {
	return repository.GetTask(id)
}

func (s *service) InsertTask(ctx context.Context, descricao string, prazo string) (int, error) {
	return repository.InsertTask(descricao, prazo)
}

func (s *service) FinishTask(ctx context.Context, id int) (*Tarefa, error) {
	t, err := repository.GetTask(id)
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
	return repository.GetAllIncomplete()
}
