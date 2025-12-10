package domain

type Dreno struct {
	Local           string `json:"local_dreno,omitempty"`
	DataImplantacao string `json:"data_implantacao_dreno,omitempty"`
	Curativos       string `json:"curativos_dreno,omitempty"`
}

type ExameFisico struct {
	// General Evaluation
	EstadoGeral         string `json:"estado_geral"`
	AvaliacaoLocomotora string `json:"avaliacao_locomotora"`

	// Drains Section
	Drenos []Dreno `json:"drenos,omitempty"`

	PossuiAcesso bool `json:"possui_acesso"`

	CateterTotal bool `json:"cateter_total,omitempty"`

	Piccline          bool   `json:"piccline,omitempty"`
	TipoPiccline      string `json:"tipo_piccline,omitempty"`
	DataImpPiccline   string `json:"data_imp_piccline,omitempty"`
	DataTrocaPiccline string `json:"data_troca_piccline,omitempty"`
	CurativoPiccline  string `json:"curativo_piccline,omitempty"`

	AcessoCentral    bool   `json:"acesso_central,omitempty"`
	LocalCentral     string `json:"local_central,omitempty"`
	DataImpCentral   string `json:"data_imp_central,omitempty"`
	DataTrocaCentral string `json:"data_troca_central,omitempty"`
	CurativoCentral  string `json:"curativo_central,omitempty"`

	AcessoPeriferico    bool   `json:"acesso_periferico,omitempty"`
	LocalPeriferico     string `json:"local_periferico,omitempty"`
	DataTrocaPeriferico string `json:"data_troca_periferico,omitempty"`
	CurativoPeriferico  string `json:"curativo_periferico,omitempty"`

	Hipodermoclise bool   `json:"hipodermoclise,omitempty"`
	DataTrocaHipo  string `json:"data_troca_hipo,omitempty"`
	CurativoHipo   string `json:"curativo_hipo,omitempty"`

	OutrosAcessos     bool   `json:"outros_acessos,omitempty"`
	NomeOutrosAcessos string `json:"nome_outros_acessos,omitempty"`
	DataOutroAcesso   string `json:"data_outro_acesso,omitempty"`
	DataAcessoUltimo  string `json:"data_acesso_ultimo,omitempty"`

	Antimicrobianos []Antimicrobiano `json:"antimicrobianos"`
}

type Antimicrobiano struct {
	Nome        string `json:"nome_antimicrobiano"`
	DataInicio  string `json:"data_inicio"`
	DataTermino string `json:"data_termino"`
}
