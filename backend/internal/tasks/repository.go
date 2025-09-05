package tasks

import (
	"context"

	"github.com/gui-henri/learning-go/pkg/errors"
	"github.com/jackc/pgx/v5"
)

type TaskRepository interface {
	GetTask(id int) (Tarefa, error)
	InsertTask(descricao string, prazo string) (int, error)
	GetAllIncomplete() ([]Tarefa, error)
	UpdateTask(t Tarefa) error
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
	tarefa, err := newTarefa(descricao, prazo)

	if err != nil {
		return 0, err
	}

	sql := `
        INSERT INTO tasks (descricao, prazo)
        VALUES ($1, $2)
        RETURNING id
	`

	var id int
	err = s.db.QueryRow(context.Background(), sql, tarefa.Descricao, tarefa.Prazo).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *taskRepository) GetAllIncomplete() ([]Tarefa, error) {
	var tarefas []Tarefa
	rows, err := s.db.Query(context.Background(), "select * from tasks where concluida=false")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t Tarefa
		err := rows.Scan(&t.Id, &t.Descricao, &t.Prazo, &t.Concluida, &t.CriadaEm)
		if err != nil {
			return nil, err
		}
		tarefas = append(tarefas, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tarefas, nil
}

func (s *taskRepository) UpdateTask(t Tarefa) error {
	sql := `
	UPDATE tasks
	SET descricao = $1, prazo = $2, concluida = $3, criada_em = $4
	WHERE id = $5
	`
	_, err := s.db.Exec(context.Background(), sql, t.Descricao, t.Prazo, t.Concluida, t.CriadaEm, t.Id)

	if err != nil {
		return err
	}

	return nil
}
