package domain

type Cabecalho struct {
	NumeroOrcamento string `json:"numero_orcamento"`
	DataOrcamento   string `json:"data_orcamento"`
	UsuarioEmissor  string `json:"usuario_emissor"`
	Paciente        struct {
		Codigo           string `json:"codigo"`
		Nome             string `json:"nome"`
		NrCarteira       string `json:"nr_carteira"`
		ValidadeCarteira string `json:"validade_carteira"`
	} `json:"paciente"`
	Convenio struct {
		Codigo      string `json:"codigo"`
		Nome        string `json:"nome"`
		CodigoPlano string `json:"codigo_plano"`
		NomePlano   string `json:"nome_plano"`
	} `json:"convenio"`
	Internacao struct {
		DataInicio     string `json:"data_inicio"`
		DataFim        string `json:"data_fim"`
		QuantidadeDias string `json:"quantidade_dias"`
		Tipo           string `json:"tipo"`
		Complexidade   string `json:"complexidade"`
	} `json:"internacao"`
	ValoresProposta struct {
		TotalPeriodo string `json:"total_periodo"`
		MedioDiario  string `json:"medio_diario"`
	} `json:"valores_proposta"`
}
