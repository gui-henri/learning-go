package domain

type FieldOption struct {
	Label  string `json:"label"`
	Value  string `json:"value"`
	Points int    `json:"points,omitempty"`
}

type FormMetadata map[string][]FieldOption

type AvaliationFormOptions struct {
	Version            int
	DadosGerais        FormMetadata `json:"dados_gerais"`
	Contato            FormMetadata `json:"contato"`
	Cuidadores         FormMetadata `json:"cuidadores"`
	SegurancaPaciente  FormMetadata `json:"seguranca"`
	DadosClinicos      FormMetadata `json:"dados_clinicos"`
	CardioRespiratorio FormMetadata `json:"cardio_respiratorio"`
	Nutricional        FormMetadata `json:"nutricional"`
	Eliminacoes        FormMetadata `json:"eliminacoes"`
	Avaliacao          FormMetadata `json:"avaliacao"`
}

func GetAvaliationFormOptions() AvaliationFormOptions {
	return AvaliationFormOptions{
		Version:            1,
		DadosGerais:        GetDadosGeraisFormOptions(),
		Contato:            GetContatoFormOptions(),
		Cuidadores:         GetCuidadoresFormOptions(),
		SegurancaPaciente:  GetSegurancaFormOptions(),
		DadosClinicos:      GetDadosClinicosFormOptions(),
		CardioRespiratorio: GetCardioRespiratorioFormOptions(),
		Nutricional:        GetNutricionalFormOptions(),
		Eliminacoes:        GetEliminacoesFormOptions(),
		Avaliacao:          GetAvaliacaoFormOptions(),
	}
}
