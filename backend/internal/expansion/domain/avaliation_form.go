package domain

type FieldOption struct {
	Label  string `json:"label"`
	Value  string `json:"value"`
	Points int    `json:"points"`
}

type FormMetadata map[string][]FieldOption

type AvaliationFormOptions struct {
	DadosGerais        FormMetadata `json:"dados_gerais"`
	Contato            FormMetadata `json:"contato"`
	DadosClinicos      FormMetadata `json:"dados_clinicos"`
	CardioRespiratorio FormMetadata `json:"cardio_respiratorio"`
	Nutricional        FormMetadata `json:"nutricional"`
	Eliminacoes        FormMetadata `json:"eliminacoes"`
	Avaliacao          FormMetadata `json:"avaliacao"`
}

func GetAvaliationFormOptions() AvaliationFormOptions {
	return AvaliationFormOptions{
		DadosGerais:        GetDadosGeraisFormOptions(),
		Contato:            GetContatoFormOptions(),
		DadosClinicos:      GetDadosClinicosFormOptions(),
		CardioRespiratorio: GetCardioRespiratorioFormOptions(),
		Nutricional:        GetNutricionalFormOptions(),
		Eliminacoes:        GetEliminacoesFormOptions(),
		Avaliacao:          GetAvaliacaoFormOptions(),
	}
}
