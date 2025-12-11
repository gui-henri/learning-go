package domain

type Seguranca struct {
	Alergias []Alergia `json:"alergias"`
}

type Alergia struct {
	TipoAlergia        string `json:"tipo_alergia"` // FIXED: Object in JSON
	QuaisAlergias      string `json:"quais_alergias"`
	Precaucao          string `json:"precaucao"` // FIXED: Object in JSON
	CuidadosPaliativos string `json:"cuidados_paliativos"`
}

func GetSegurancaFormOptions() FormMetadata {
	return FormMetadata{
		"alergia": []FieldOption{
			{Label: "Medicamentosa", Value: "alergia_medicamentosa"},
			{Label: "Alimentar", Value: "alergia_alimentar"},
			{Label: "Respiratória", Value: "alergia_respiratoria"},
			{Label: "Outros", Value: "alergia_outros"},
		},
		"precaucao": []FieldOption{
			{Label: "Padrão", Value: "prec_padrao"},
			{Label: "Contato", Value: "prec_contato"},
			{Label: "Gotículas", Value: "prec_goticulas"},
			{Label: "Aerossóis", Value: "prec_aerossois"},
		},
	}
}
