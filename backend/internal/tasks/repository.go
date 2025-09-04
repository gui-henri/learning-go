package tasks

import (
	"context"
	"time"

	"github.com/gui-henri/learning-go/pkg/errors"
	"github.com/jackc/pgx/v5"
)

var tasks = make([]Tarefa, 0)

type TaskRepository interface {
	GetTask(id int) (Tarefa, error)
	InsertTask(descricao string, prazo string) (int, error)
	GetAllIncomplete() ([]Tarefa, error)
}

type taskRepository struct {
	db *pgx.Conn
}

func NewTaskRepository(db *pgx.Conn) *taskRepository {
	return &taskRepository{
		db: db,
	}
}

func (s *taskRepository) GetTask(id int) (*Tarefa, error) {
	for idx := range tasks {
		if tasks[idx].Id == id {
			return &tasks[idx], nil
		}
	}

	var newTarefa Tarefa
	err := s.db.QueryRow(context.Background(), "select * from tasks where id=$1", id).Scan(
		&newTarefa.Id,
		&newTarefa.Descricao,
		&newTarefa.Prazo,
		&newTarefa.Concluida,
		&newTarefa.CriadaEm,
	)

	if err != nil {
		return &Tarefa{}, errors.NotFound
	}

	return &newTarefa, nil
}

func (s *taskRepository) InsertTask(descricao string, prazo string) (int, error) {
	id := len(tasks)
	tasks = append(tasks, Tarefa{
		Id:        id,
		Descricao: descricao,
		Prazo:     prazo,
		Concluida: false,
		CriadaEm:  time.Now(),
	})

	return id, nil
}

func (s *taskRepository) GetAllIncomplete() []Tarefa {
	incompletedTasks := make([]Tarefa, 0)

	for _, t := range tasks {
		if t.Concluida {
			continue
		}

		incompletedTasks = append(incompletedTasks, t)
	}

	return incompletedTasks
}
