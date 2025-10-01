package tasks

import (
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	apperrors "github.com/gui-henri/learning-go/pkg/errors"
)

type Tarefa struct {
	Id         int       `json:"id"`
	Descricao  string    `json:"descricao" validate:"required"`
	Prazo      string    `json:"prazo" validate:"required"`
	Concluida  bool      `json:"concluida"`
	CriadaEm   time.Time `json:"criada_em" validate:"lte"`
	PacienteID string    `json:"patient_id"`
}

func newTarefa(descricao string, prazo string) (Tarefa, error) {

	descricao = strings.TrimSpace(descricao)
	prazo = strings.TrimSpace(prazo)

	if len(descricao) == 0 || len(prazo) == 0 {
		return Tarefa{}, apperrors.InvalidInput
	}

	tarefa := Tarefa{
		Descricao: descricao,
		Prazo:     prazo,
		Concluida: false,
		CriadaEm:  time.Now(),
	}

	validate := validator.New()

	err := validate.Struct(tarefa)

	if err != nil {
		return Tarefa{}, apperrors.InvalidInput
	}

	return tarefa, nil
}
