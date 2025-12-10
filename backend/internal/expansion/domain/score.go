package domain

type Score struct {
	DiagnosticoPrimario     string `json:"diagnostico_primario"`
	DiagnosticoSecundario   string `json:"diagnostico_secundario"`
	DomicilioRisco          bool   `json:"domicilio_risco"`
	ImpedimentoDeslocamento bool   `json:"impedimento_deslocamento"`
	AlimentacaoParenteral   string `json:"alimentacao_parenteral"`
	AspiracaoTraqueo        string `json:"aspiracao_traqueo"`
	VentilacaoMecanica      string `json:"ventilacao_mecanica"`
	MedicacaoParenteral     string `json:"medicacao_parenteral"`
	Atividades              bool   `json:"atividades"`
	Banho                   bool   `json:"banho"`
	Vestir                  bool   `json:"vestir"`
	HigienePessoal          bool   `json:"higiene_pessoal"`
	Transferencia           bool   `json:"transferencia"`
	Continencia             bool   `json:"continencia"`
	Alimentacao             bool   `json:"alimentacao"`
	EstadoNutricional       string `json:"estado_nutricional"`
	ViaAlimentacaoMedicacao string `json:"via_alimentacao_medicacao"`
	InternacoesUltimoAno    string `json:"internacoes_ultimo_ano"`
	AspiracaoViasAereas     string `json:"aspiracao_vias_aereas"`
	Lesoes                  string `json:"lesoes"`
	ViaMedicacoes           string `json:"via_medicacoes"`
	ExerciciosVentilatorios string `json:"exercicios_ventilatorios"`
	UsoOxigenoterapia       string `json:"uso_oxigenoterapia"`
	NivelConsciencia        string `json:"nivel_consciencia"`
}

func GetAvaliacaoFormOptions() FormMetadata {
	return FormMetadata{
		"freq_12h": []FieldOption{
			{Label: "Não utiliza", Value: "nao_utiliza", Points: 0},
			{Label: "Até 12 horas por dia", Value: "ate_12h", Points: 0},
			{Label: "Mais de 12 horas por dia", Value: "mais_12h", Points: 0},
		},
		"freq_5x": []FieldOption{
			{Label: "Não utiliza", Value: "nao_utiliza", Points: 0},
			{Label: "Até 5x por dia", Value: "ate_5x", Points: 0},
			{Label: "Mais que 5x por dia", Value: "mais_5x", Points: 0},
		},
		"freq_4x": []FieldOption{
			{Label: "Não utiliza", Value: "nao_utiliza", Points: 0},
			{Label: "Até 4x por dia", Value: "ate_4x", Points: 0},
			{Label: "Mais que 4x por dia", Value: "mais_4x", Points: 0},
		},

		"katz": []FieldOption{
			{Label: "Independente (Sem ajuda)", Value: "independente", Points: 1},
			{Label: "Com ajuda / Dependente", Value: "dependente", Points: 0},
		},

		"estado_nutricional": []FieldOption{
			{Label: "Eutrófico", Value: "eutrofico", Points: 0},
			{Label: "Sobrepeso/Emagrecido", Value: "sobrepeso_emagrecido", Points: 1},
			{Label: "Obeso/Desnutrido", Value: "obeso_desnutrido", Points: 2},
		},
		"via_alimentacao": []FieldOption{
			{Label: "Sem auxílio", Value: "sem_auxilio", Points: 0},
			{Label: "Assistida", Value: "assistida", Points: 1},
			{Label: "Gastrostomia/Jejunostomia", Value: "gastro_jejuno", Points: 2},
			{Label: "SNG/SNE", Value: "sng_sne", Points: 3},
		},

		"internacoes": []FieldOption{
			{Label: "0-1", Value: "0-1", Points: 0},
			{Label: "2-3", Value: "2-3", Points: 1},
			{Label: "4+", Value: "4+", Points: 2},
		},
		"aspiracao_vias_aereas": []FieldOption{
			{Label: "Ausente", Value: "ausente", Points: 0},
			{Label: "Até 5x por dia", Value: "ate_5x", Points: 1},
			{Label: "Mais que 5x por dia", Value: "mais_5x", Points: 2},
		},
		"lesoes": []FieldOption{
			{Label: "Nenhuma ou Única Simples", Value: "nenhuma_simples", Points: 0},
			{Label: "Múltiplas (simples) ou Única (complexo)", Value: "multiplas_simples_unica_complexa", Points: 1},
			{Label: "Múltiplas lesões (complexas)", Value: "multiplas_complexas", Points: 2},
		},
		"via_medicacoes": []FieldOption{
			{Label: "Via enteral", Value: "enteral", Points: 0},
			{Label: "Intramuscular ou subcutânea", Value: "im_sc", Points: 1},
			{Label: "IV até 4x/dia ou Hipodermóclise", Value: "iv_hipo", Points: 2},
		},

		"frequencia_respiratoria": []FieldOption{
			{Label: "Ausente", Value: "ausente", Points: 0},
			{Label: "Intermitente", Value: "intermitente", Points: 1},
			{Label: "Contínuo", Value: "continuo", Points: 2},
		},
		"nivel_consciencia": []FieldOption{
			{Label: "Alerta", Value: "alerta", Points: 0},
			{Label: "Confuso/Desorientado", Value: "confuso", Points: 1},
			{Label: "Comatoso", Value: "comatoso", Points: 2},
		},
	}
}
