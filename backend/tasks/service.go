package tasks

import (
	"context"
)

type TaskService interface {
	GetTask(ctx context.Context, id int) (Tarefa, error)
	InsertTask(ctx context.Context, descricao string, prazo string) (int, error)
}

type service struct{}

func NewService() *service {
	return &service{}
}

var repository = NewTaskRepository()

func (s *service) GetTask(ctx context.Context, id int) (Tarefa, error) {
	return repository.GetTask(id)
}

func (s *service) InsertTask(ctx context.Context, descricao string, prazo string) (int, error) {
	return repository.InsertTask(descricao, prazo)
}
