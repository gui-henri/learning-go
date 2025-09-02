package tasks

import (
	"time"

	"github.com/gui-henri/learning-go/pkg/errors"
)

var tasks = make([]Tarefa, 0)

type TaskRepository interface {
	GetTask(id int) (Tarefa, error)
	InsertTask(descricao string, prazo string) (int, error)
}

type taskRepository struct{}

func NewTaskRepository() *taskRepository {
	return &taskRepository{}
}

func (s *taskRepository) GetTask(id int) (Tarefa, error) {
	for idx := range tasks {
		t := tasks[idx]

		if t.Id == id {
			return t, nil
		}
	}

	return Tarefa{}, errors.NotFound
}

func (s *taskRepository) InsertTask(descricao string, prazo string) (int, error) {
	id := len(tasks)
	tasks = append(tasks, Tarefa{
		Id:        id,
		Descricao: descricao,
		Prazo:     prazo,
		Concluida: false,
		CriadaEm:  time.Now().GoString(),
	})

	return id, nil
}
