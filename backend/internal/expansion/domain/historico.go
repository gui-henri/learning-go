package domain

type Historico struct {
	HistoriaDoencaAtual string `json:"historia_doenca_atual"`
	AnamneseEnfermagem  string `json:"anamnese_enfermagem"`
	PlanoTerapeutico    string `json:"plano_terapeutico"`
}
