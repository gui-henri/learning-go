package domain

import "time"

type ExameFisico struct {
	// General Evaluation
	EstadoGeral         string `json:"estado_geral"`
	AvaliacaoLocomotora string `json:"avaliacao_locomotora"`

	// Drains Section
	PossuiDreno          bool       `json:"possui_dreno"`
	LocalDreno           string     `json:"local_dreno,omitempty"`
	DataImplantacaoDreno *time.Time `json:"data_implantacao_dreno,omitempty"`
	CurativosDreno       string     `json:"curativos_dreno,omitempty"`

	PossuiAcesso bool `json:"possui_acesso"`

	CateterTotal bool `json:"cateter_total,omitempty"`

	Piccline          bool       `json:"piccline,omitempty"`
	TipoPiccline      string     `json:"tipo_piccline,omitempty"`
	DataImpPiccline   *time.Time `json:"data_imp_piccline,omitempty"`
	DataTrocaPiccline *time.Time `json:"data_troca_piccline,omitempty"`
	CurativoPiccline  string     `json:"curativo_piccline,omitempty"`

	AcessoCentral    bool       `json:"acesso_central,omitempty"`
	LocalCentral     string     `json:"local_central,omitempty"`
	DataImpCentral   *time.Time `json:"data_imp_central,omitempty"`
	DataTrocaCentral *time.Time `json:"data_troca_central,omitempty"`
	CurativoCentral  string     `json:"curativo_central,omitempty"`

	AcessoPeriferico    bool       `json:"acesso_periferico,omitempty"`
	LocalPeriferico     string     `json:"local_periferico,omitempty"`
	DataTrocaPeriferico *time.Time `json:"data_troca_periferico,omitempty"`
	CurativoPeriferico  string     `json:"curativo_periferico,omitempty"`

	Hipodermoclise bool       `json:"hipodermoclise,omitempty"`
	DataTrocaHipo  *time.Time `json:"data_troca_hipo,omitempty"`
	CurativoHipo   string     `json:"curativo_hipo,omitempty"`

	OutrosAcessos     bool       `json:"outros_acessos,omitempty"`
	NomeOutrosAcessos string     `json:"nome_outros_acessos,omitempty"`
	DataOutroAcesso   *time.Time `json:"data_outro_acesso,omitempty"`
	DataAcessoUltimo  *time.Time `json:"data_acesso_ultimo,omitempty"`

	Antimicrobianos []Antimicrobiano `json:"antimicrobianos"`
}

type Antimicrobiano struct {
	Nome        string `json:"nome_antimicrobiano"`
	DataInicio  string `json:"data_inicio"`
	DataTermino string `json:"data_termino"`
}
