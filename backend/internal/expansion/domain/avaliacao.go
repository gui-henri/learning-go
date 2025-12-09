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

type SelectField struct {
	Name  string `json:"name"`
	Code  string `json:"code"`
	Value string `json:"value"` // Added Value because 'Seguranca' uses 'value' instead of 'code'
}

type AvaliacaoRequest struct {
	DadosGerais        DadosGerais        `json:"dadosGerais"`
	Endereco           Endereco           `json:"endereco"`
	Contato            Contato            `json:"contato"`
	Cuidadores         Cuidadores         `json:"cuidadores"`
	Seguranca          Seguranca          `json:"seguranca"`
	HistoricoClinico   Historico          `json:"historicoClinico"`
	ExameFisico        ExameFisico        `json:"exameFisico"`
	CardioRespiratorio CardioRespiratorio `json:"cardioRespiratorio"`
	Nutricional        Nutricional        `json:"nutricional"`
	Eliminacoes        Eliminacoes        `json:"eliminacoes"`
	CondicoesPele      CondicoesPele      `json:"condicoesPele"`
	Score              Score              `json:"score"`
	Observacoes        Observacoes        `json:"observacoes"`
}
