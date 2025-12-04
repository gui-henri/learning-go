package avaliation

import (
	"encoding/json"
	"time"

	"github.com/gui-henri/learning-go/internal/avaliation/domain"
)

type AvaliacaoSchema struct {
	Id           int             `db:"id" json:"id"`
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
	DadosGerais        domain.DadosGerais        `json:"dadosGerais"`
	Endereco           domain.Endereco           `json:"endereco"`
	Contato            domain.Contato            `json:"contato"`
	Cuidadores         domain.Cuidadores         `json:"cuidadores"`
	Seguranca          domain.Seguranca          `json:"seguranca"`
	HistoricoClinico   domain.Historico          `json:"historicoClinico"`
	ExameFisico        domain.ExameFisico        `json:"exameFisico"`
	CardioRespiratorio domain.CardioRespiratorio `json:"cardioRespiratorio"`
	Nutricional        domain.Nutricional        `json:"nutricional"`
	Eliminacoes        domain.Eliminacoes        `json:"eliminacoes"`
	CondicoesPele      domain.CondicoesPele      `json:"condicoesPele"`
	Score              domain.Score              `json:"score"`
	Observacoes        domain.Observacoes        `json:"observacoes"`
}

// --- CUIDADORES ---
type Cuidadores struct {
	MedicoSolicitante  string      `json:"medico_solicitante"`
	ContatoMedico      string      `json:"contato_medico"`
	PossuiCuidador     bool        `json:"possui_cuidador"`
	NomeCuidador       string      `json:"nome_cuidador"`
	ContatoCuidador    string      `json:"contato_cuidador"`
	TurnoCuidador      SelectField `json:"turno_cuidador"`
	PrecisaTreinamento bool        `json:"precisa_treinamento"`
	ObsTreinamento     string      `json:"obs_treinamento"`
}

// --- SEGURANÇA ---
type Seguranca struct {
	Alergias []Alergia `json:"alergias"`
}

type Alergia struct {
	TipoAlergia        SelectField `json:"tipo_alergia"` // FIXED: Object in JSON
	QuaisAlergias      string      `json:"quais_alergias"`
	Precaucao          SelectField `json:"precaucao"` // FIXED: Object in JSON
	CuidadosPaliativos string      `json:"cuidados_paliativos"`
}

// --- HISTÓRICO CLÍNICO ---
type Historico struct {
	HistoriaDoencaAtual string `json:"historia_doenca_atual"`
	AnamneseEnfermagem  string `json:"anamnese_enfermagem"`
	PlanoTerapeutico    string `json:"plano_terapeutico"`
}

// --- EXAME FÍSICO ---
type ExameFisico struct {
	EstadoGeral          string           `json:"estado_geral"`         // FIXED
	AvaliacaoLocomotora  string           `json:"avaliacao_locomotora"` // FIXED
	PossuiDreno          bool             `json:"possui_dreno"`
	LocalDreno           string           `json:"local_dreno"`
	DataImplantacaoDreno string           `json:"data_implantacao_dreno"`
	CurativosDreno       string           `json:"curativos_dreno"`
	PossuiAcesso         bool             `json:"possui_acesso"`
	Antimicrobianos      []Antimicrobiano `json:"antimicrobianos"`
}

type Antimicrobiano struct {
	Nome        string `json:"nome_antimicrobiano"`
	DataInicio  string `json:"data_inicio"`
	DataTermino string `json:"data_termino"`
}
