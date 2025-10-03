package tasks

import (
	"context"
	"fmt"

	"github.com/gui-henri/learning-go/db"
	"github.com/gui-henri/learning-go/pkg/errors"
)

type TaskRepository interface {
	GetTask(id int) (Tarefa, error)
	InsertTask(descricao string, prazo string, pacienteId string) (Tarefa, error)
	GetAll() ([]Tarefa, error)
	GetAllIncomplete() ([]Tarefa, error)
	UpdateTask(t Tarefa) error
	DeleteTask(id int) error
}

type taskRepository struct {
	db db.IDB
}

func NewTaskRepository(db db.IDB) *taskRepository {
	return &taskRepository{
		db: db,
	}
}

func (s *taskRepository) GetTask(id int) (*Tarefa, error) {

	var newTarefa Tarefa
	err := s.db.QueryRow(context.Background(), "select id, descricao, prazo, concluida, criada_em from tasks where id=$1", id).Scan(
		&newTarefa.Id,
		&newTarefa.Descricao,
		&newTarefa.Prazo,
		&newTarefa.Concluida,
		&newTarefa.CriadaEm,
	)

	if err != nil {
		fmt.Println(err)
		return &Tarefa{}, errors.NotFound
	}

	return &newTarefa, nil
}

func (s *taskRepository) InsertTask(descricao string, prazo string, pacienteId string) (*Tarefa, error) {
	tarefaValidation, err := newTarefa(descricao, prazo)

	if err != nil {
		return &Tarefa{}, err
	}

	sql := `
        INSERT INTO tasks (descricao, prazo, patient_id)
        VALUES ($1, $2, $3)
        RETURNING id, descricao, prazo, concluida, criada_em
	`

	var tarefa Tarefa
	// TODO: VALIDAR SE PACIENTE EXISTE ANTES DE ADICIONAR
	err = s.db.QueryRow(context.Background(), sql, tarefaValidation.Descricao, tarefaValidation.Prazo, pacienteId).Scan(
		&tarefa.Id,
		&tarefa.Descricao,
		&tarefa.Prazo,
		&tarefa.Concluida,
		&tarefa.CriadaEm,
	)

	if err != nil {
		return &Tarefa{}, errors.InvalidInput
	}

	return &tarefa, nil
}

func (s *taskRepository) GetAllIncomplete() ([]Tarefa, error) {
	tarefas := make([]Tarefa, 0)
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

func (s *taskRepository) GetAll() ([]Tarefa, error) {
	tarefas := make([]Tarefa, 0)

	sql := `
		select 
			t.id, 
			t.descricao, 
			t.prazo, 
			t.concluida, 
			t.criada_em
			p.full_name
			p.id
		from tasks t
		join patient p ON t.patient_id = p.id
	`

	rows, err := s.db.Query(context.Background(), sql)

	if err != nil {
		fmt.Println("[ERROR] Erro ao realizar query: ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t Tarefa
		err := rows.Scan(&t.Id, &t.Descricao, &t.Prazo, &t.Concluida, &t.CriadaEm, &t.Paciente.FullName, &t.Paciente.ID)
		if err != nil {
			fmt.Println("[ERROR] Erro ao fazer scan dos dados: ", err)
			return nil, err
		}
		tarefas = append(tarefas, t)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("[ERROR] Erro ao ler as linhas do banco: ", err)
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
		fmt.Println("[ERROR] Erro ao executar sql. ", err)
		return err
	}

	return nil
}

func (s *taskRepository) DeleteTask(id int) error {
	sql := `
	DELETE FROM tasks 
	where id = $1
	`
	_, err := s.db.Exec(context.Background(), sql, id)

	if err != nil {
		fmt.Println("[ERROR] Erro ao executar sql. ", err)
		return err
	}

	return nil
}
