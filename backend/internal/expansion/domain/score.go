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
	AtividadesVidaDiaria    bool   `json:"atividades_vida_diaria"`
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
			{Label: "Não utiliza", Value: "nao_utiliza"},
			{Label: "Até 12 horas por dia", Value: "ate_12h"},
			{Label: "Mais de 12 horas por dia", Value: "mais_12h"},
		},
		"freq_5x": []FieldOption{
			{Label: "Não utiliza", Value: "nao_utiliza"},
			{Label: "Até 5x por dia", Value: "ate_5x"},
			{Label: "Mais que 5x por dia", Value: "mais_5x"},
		},
		"freq_4x": []FieldOption{
			{Label: "Não utiliza", Value: "nao_utiliza"},
			{Label: "Até 4x por dia", Value: "ate_4x"},
			{Label: "Mais que 4x por dia", Value: "mais_4x"},
		},

		"katz": []FieldOption{
			{Label: "Independente (Sem ajuda)", Value: "independente"},
			{Label: "Com ajuda / Dependente", Value: "dependente"},
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
		"exercicio_ventilatorio": []FieldOption{
			{Label: "Ausente", Value: "ausente", Points: 0},
			{Label: "Intermitente", Value: "intermitente", Points: 1},
		},

		"uso_oxigenoterapia": []FieldOption{
			{Label: "Ausente", Value: "ausente", Points: 0},
			{Label: "Intermitente", Value: "intermitente", Points: 1},
			{Label: "Continuo", Value: "continuo", Points: 2},
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

func CalcularScoreKatz(score Score) int {
	total := 0

	values := []bool{
		score.Banho,
		score.Vestir,
		score.HigienePessoal,
		score.Transferencia,
		score.Continencia,
		score.Alimentacao,
	}

	for _, v := range values {
		if v == true {
			total++
		}
	}

	return total
}

func KatzParaPontosNEAD(scoreKatz int) int {
	switch {
	case scoreKatz < 2:
		return 2
	case scoreKatz == 3 || scoreKatz == 4:
		return 1
	default:
		return 0
	}
}

func CalcularPontosNEADBase(score Score, options FormMetadata) int {
	total := 0

	mapper := map[string]string{
		"estado_nutricional":    score.EstadoNutricional,
		"via_alimentacao":       score.ViaAlimentacaoMedicacao,
		"internacoes":           score.InternacoesUltimoAno,
		"aspiracao_vias_aereas": score.AspiracaoViasAereas,
		"lesoes":                score.Lesoes,
		"via_medicacoes":        score.ViaMedicacoes,
		"uso_oxigenoterapia":    score.UsoOxigenoterapia,
		"nivel_consciencia":     score.NivelConsciencia,
	}

	for field, value := range mapper {
		opts := options[field]
		for _, opt := range opts {
			if opt.Value == value {
				total += opt.Points
				break
			}
		}
	}

	return total
}

type AvaliacaoResultado struct {
	ScoreKatz       int `json:"score_katz"`
	PontosKatzNEAD  int `json:"pontos_katz_nead"`
	PontosNEADBase  int `json:"pontos_nead_base"`
	PontosNEADFinal int `json:"pontos_nead_final"`
}

func CalcularResultado(score Score) AvaliacaoResultado {
	options := GetAvaliacaoFormOptions()

	scoreKatz := CalcularScoreKatz(score)
	pontosKatzNEAD := KatzParaPontosNEAD(scoreKatz)
	pontosNEADBase := CalcularPontosNEADBase(score, options)

	return AvaliacaoResultado{
		ScoreKatz:       scoreKatz,
		PontosKatzNEAD:  pontosKatzNEAD,
		PontosNEADBase:  pontosNEADBase,
		PontosNEADFinal: pontosNEADBase + pontosKatzNEAD,
	}
}
