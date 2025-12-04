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
