package domain

import (
	"encoding/json"
	"time"
)

type AvaliacaoSchema struct {
	Id           string          `db:"id" json:"id"`
	ResourceJSON json.RawMessage `db:"resource_json" json:"resource_json"`
	CreatedAt    time.Time       `db:"created_at" json:"created_at"`
	LastUpdated  time.Time       `db:"last_updated" json:"last_updated"`
	Active       bool            `db:"active" json:"active"`
}

type AvaliacaoRequest struct {
	DadosGerais        *DadosGerais        `json:"dadosGerais,omitempty"`
	Endereco           *Endereco           `json:"endereco,omitempty"`
	Contato            *Contato            `json:"contato,omitempty"`
	Cuidadores         *Cuidadores         `json:"cuidadores,omitempty"`
	Seguranca          *Seguranca          `json:"seguranca,omitempty"`
	HistoricoClinico   *Historico          `json:"historicoClinico,omitempty"`
	ExameFisico        *ExameFisico        `json:"exameFisico,omitempty"`
	CardioRespiratorio *CardioRespiratorio `json:"cardioRespiratorio,omitempty"`
	Nutricional        *Nutricional        `json:"nutricional,omitempty"`
	Eliminacoes        *Eliminacoes        `json:"eliminacoes,omitempty"`
	CondicoesPele      *CondicoesPele      `json:"condicoesPele,omitempty"`
	Score              *Score              `json:"score,omitempty"`
	Observacoes        *Observacoes        `json:"observacoes,omitempty"`
}

type AvaliacaoListDto struct {
	Id           string
	NomePaciente string
	Idade        int
	CPF          string
}
