package domain

type CondicoesPele struct {
	CondicaoCutaneaMucosa string     `json:"condicao_cutanea_mucosa"`
	RealizaCurativo       bool       `json:"realiza_curativo"`
	Curativos             []Curativo `json:"curativos"`
}

type Curativo struct {
	LocalCurativo     string `json:"local_curativo"`
	TipoCobertura     string `json:"tipo_cobertura"`
	TamanhoCurativo   string `json:"tamanho_curativo"`
	QtdCurativo       string `json:"qtd_curativo"`
	FrequenciaTroca   string `json:"frequencia_troca"`
	MateriaisCurativo string `json:"materiais_curativo"`
}
