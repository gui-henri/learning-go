package avaliation

import (
	"encoding/json"
	"time"
)

type AvaliacaoSchema struct {
	Id           int             `db:"id" json:"id"`
	ResourceJSON json.RawMessage `db:"resource_json" json:"resource_json"`
	CreatedAt    time.Time       `db:"created_at" json:"created_at"`
	LastUpdated  time.Time       `db:"last_updated" json:"last_updated"`
	Active       bool            `db:"active" json:"active"`
}

type SelectField struct {
	Name  string `json:"name"`
	Code  string `json:"code"`
	Value string `json:"value"` // Added Value because 'Seguranca' uses 'value' instead of 'code'
}

type AvaliacaoRequest struct {
	DadosGerais        DadosGerais        `json:"dadosGerais"`
	Endereco           Endereco           `json:"endereco"`
	Contato            Contato            `json:"contato"`
	Cuidadores         Cuidadores         `json:"cuidadores"`
	Seguranca          Seguranca          `json:"seguranca"`
	HistoricoClinico   Historico          `json:"historicoClinico"`
	ExameFisico        ExameFisico        `json:"exameFisico"`
	CardioRespiratorio CardioRespiratorio `json:"cardioRespiratorio"`
	Nutricional        Nutricional        `json:"nutricional"`
	Eliminacoes        Eliminacoes        `json:"eliminacoes"`
	CondicoesPele      CondicoesPele      `json:"condicoesPele"`
	Score              Score              `json:"score"`
	Observacoes        Observacoes        `json:"observacoes"`
}

// --- DADOS GERAIS ---
type DadosGerais struct {
	Genero       SelectField `json:"genero"` // FIXED: Changed from string to Struct
	Complexidade SelectField `json:"complexidade"`

	// FIXED: JSON sends "123" (string), not 123 (number). Changed to string.
	// You can convert these to float/int inside your Service layer using strconv.
	Peso   string `json:"peso"`
	Altura string `json:"altura"`
	Idade  string `json:"idade"`

	NomePaciente          string `json:"nome_paciente"`
	PrevisaoAlta          bool   `json:"previsao_alta"`   // JSON sends boolean true/false
	DataNascimento        string `json:"data_nascimento"` // Matches ISO string
	CPF                   string `json:"cpf"`
	RG                    string `json:"rg"`
	NomePai               string `json:"nome_pai"`
	NomeMae               string `json:"nome_mae"`
	Convenio              string `json:"convenio"`
	Hospital              string `json:"hospital"`
	DataAdmissao          string `json:"dat_admissao"`
	Apartamento           string `json:"apartamento"`
	Carteirinha           string `json:"carteirinha"`
	VencimentoCarteirinha string `json:"vencimento_carteirinha"`
}

// --- ENDEREÇO ---
type Endereco struct {
	CEP             string `json:"cep"`
	Estado          string `json:"estado"`
	Cidade          string `json:"cidade"`
	Bairro          string `json:"bairro"`
	Logradouro      string `json:"logradouro"`
	Numero          string `json:"numero"`
	Complemento     string `json:"complemento"`
	PontoReferencia string `json:"ponto_referencia"`
}

// --- CONTATO ---
type Contato struct {
	TelefoneResidencial string        `json:"telefone_residencial"`
	TelefonePaciente    string        `json:"telefone_paciente"`
	EmailPaciente       string        `json:"email_paciente"`
	Responsaveis        []Responsavel `json:"responsaveis"`
}

type Responsavel struct {
	Nome              string      `json:"nome"`
	Parentesco        string      `json:"parentesco"`
	Email             string      `json:"email"`
	Telefone          string      `json:"telefone"`
	FormaContato      SelectField `json:"forma_contato"`
	NumeroPrioritario bool        `json:"numero_prioritario"`
}

// --- CUIDADORES ---
type Cuidadores struct {
	MedicoSolicitante  string      `json:"medico_solicitante"`
	ContatoMedico      string      `json:"contato_medico"`
	PossuiCuidador     bool        `json:"possui_cuidador"`
	NomeCuidador       string      `json:"nome_cuidador"`
	ContatoCuidador    string      `json:"contato_cuidador"`
	TurnoCuidador      SelectField `json:"turno_cuidador"`
	PrecisaTreinamento bool        `json:"precisa_treinamento"`
	ObsTreinamento     string      `json:"obs_treinamento"`
}

// --- SEGURANÇA ---
type Seguranca struct {
	Alergias []Alergia `json:"alergias"`
}

type Alergia struct {
	TipoAlergia        SelectField `json:"tipo_alergia"` // FIXED: Object in JSON
	QuaisAlergias      string      `json:"quais_alergias"`
	Precaucao          SelectField `json:"precaucao"` // FIXED: Object in JSON
	CuidadosPaliativos string      `json:"cuidados_paliativos"`
}

// --- HISTÓRICO CLÍNICO ---
type Historico struct {
	HistoriaDoencaAtual string `json:"historia_doenca_atual"`
	AnamneseEnfermagem  string `json:"anamnese_enfermagem"`
	PlanoTerapeutico    string `json:"plano_terapeutico"`
}

// --- EXAME FÍSICO ---
type ExameFisico struct {
	EstadoGeral          SelectField      `json:"estado_geral"`         // FIXED
	AvaliacaoLocomotora  SelectField      `json:"avaliacao_locomotora"` // FIXED
	PossuiDreno          bool             `json:"possui_dreno"`
	LocalDreno           string           `json:"local_dreno"`
	DataImplantacaoDreno string           `json:"data_implantacao_dreno"`
	CurativosDreno       string           `json:"curativos_dreno"`
	PossuiAcesso         bool             `json:"possui_acesso"`
	Antimicrobianos      []Antimicrobiano `json:"antimicrobianos"`
}

type Antimicrobiano struct {
	Nome        string `json:"nome_antimicrobiano"`
	DataInicio  string `json:"data_inicio"`
	DataTermino string `json:"data_termino"`
}

type CardioRespiratorio struct {
	PadraoRespiratorio  SelectField `json:"padrao_respiratorio"`
	ViaAerea            SelectField `json:"via_aerea"`
	NumeroTubo          string      `json:"numero_tubo"`
	DataOrotraqueal     string      `json:"data_orotraqueal"`
	SuporteVentilatorio SelectField `json:"suporte_ventilatorio"`
	// VNI
	ModoVNI           SelectField `json:"modo_vni"`
	FrequenciaVNI     SelectField `json:"frequencia_vni"`
	Fio2VNI           string      `json:"fio2_vni"`
	MascaraVNI        SelectField `json:"mascara_vni"`
	TamanhoMascaraVNI SelectField `json:"tamanho_mascara_vni"`
	// AVM
	ModoAVM           SelectField `json:"modo_avm"`
	FrequenciaAVM     SelectField `json:"frequencia_avm"`
	Fio2AVM           string      `json:"fio2_avm"`
	MascaraAVM        SelectField `json:"mascara_avm"`
	TamanhoMascaraAVM SelectField `json:"tamanho_mascara_avm"`
	// Oxigenoterapia
	UsaOxigenioterapia bool        `json:"usa_oxigenioterapia"`
	FonteOxigenio      SelectField `json:"fonte_oxigenio"`
	FrequenciaOxigenio SelectField `json:"frequencia_oxigenio"`
	VezesDiaOxigenio   string      `json:"vezes_dia_oxigenio"`
	ModalidadeOxigenio SelectField `json:"modalidade_oxigenio"`
	FluxoOxigenio      string      `json:"fluxo_oxigenio"`
}

type Nutricional struct {
	AlimentacaoOral       bool        `json:"alimentacao_oral"`
	AlimentacaoEnteral    bool        `json:"alimentacao_enteral"`
	AlimentacaoParenteral bool        `json:"alimentacao_parenteral"`
	Sonda                 bool        `json:"sonda"`
	Botton                bool        `json:"botton"`
	Suplemento            bool        `json:"suplemento"`
	ViaEnteral            SelectField `json:"via_enteral"`
	ViaParenteral         SelectField `json:"via_parenteral"`
	AdaptadorSonda        SelectField `json:"adaptador_sonda"`
	RestricaoAlimentar    string      `json:"restricao_alimentar"`
	TipoDieta             SelectField `json:"tipo_dieta"`
	FormaAdministracao    SelectField `json:"forma_administracao"`
	MarcaBomba            string      `json:"marca_bomba"`
	DataUltimaTroca       string      `json:"data_ultima_troca"`
	Gavando               string      `json:"gavando"`
	VolumeDiario          string      `json:"volume_diario"`
	QualSuplemento        string      `json:"qual_suplemento"`
}

type Eliminacoes struct {
	FuncaoIntestinal              SelectField `json:"funcao_intestinal"`
	ViaEvacuacaoFisiologica       bool        `json:"via_evacuacao_fisiologica"`
	ViaEvacuacaoEstomia           bool        `json:"via_evacuacao_estomia"`
	TipoEstomia                   SelectField `json:"tipo_estomia"`
	PlacaEstomia                  string      `json:"placa_estomia"`
	BolsaEstomia                  string      `json:"bolsa_estomia"`
	Diurese                       SelectField `json:"diurese"`
	NumSndDiurese                 string      `json:"num_snd_diurese"`
	UsaFralda                     bool        `json:"usa_fralda"`
	TrocasFraldaDia               string      `json:"trocas_fralda_dia"`
	SVA                           bool        `json:"sva"`
	CateterismoIntermitente       bool        `json:"cateterismo_intermitente"`
	VezesCateterismoDia           string      `json:"vezes_cateterismo_dia"`
	SVD                           bool        `json:"svd"`
	NumSVD                        string      `json:"num_svd"`
	ViasSVD                       SelectField `json:"vias_svd"`
	IrrigacaoSVD                  bool        `json:"irrigacao_svd"`
	DataTrocaSVD                  string      `json:"data_troca_svd"`
	VolumeDiureseSVD              SelectField `json:"volume_diurese_svd"`
	Cistostomia                   bool        `json:"cistostomia"`
	NumCistostomia                string      `json:"num_cistostomia"`
	ViasCistostomia               SelectField `json:"vias_cistostomia"`
	IrrigacaoCistostomia          bool        `json:"irrigacao_cistostomia"`
	QtdIrrigacaoCistostomia       string      `json:"qtd_irrigacao_cistostomia"`
	MedicacaoIrrigacaoCistostomia bool        `json:"medicacao_irrigacao_cistostomia"`
	NomeMedicacaoCistostomia      string      `json:"nome_medicacao_cistostomia"`
	VolumeDiureseCistostomia      SelectField `json:"volume_diurese_cistostomia"`
	Preservativo                  bool        `json:"preservativo"`
	VolumeDiuresePreservativo     SelectField `json:"volume_diurese_preservativo"`
}

type CondicoesPele struct {
	CondicaoCutaneaMucosa SelectField `json:"condicao_cutanea_mucosa"`
	RealizaCurativo       bool        `json:"realiza_curativo"`
	Curativos             []Curativo  `json:"curativos"`
}

type Curativo struct {
	LocalCurativo     string `json:"local_curativo"`
	TipoCobertura     string `json:"tipo_cobertura"`
	TamanhoCurativo   string `json:"tamanho_curativo"`
	QtdCurativo       string `json:"qtd_curativo"`
	FrequenciaTroca   string `json:"frequencia_troca"`
	MateriaisCurativo string `json:"materiais_curativo"`
}

type Score struct {
	DiagnosticoPrimario     string      `json:"diagnostico_primario"`
	DiagnosticoSecundario   string      `json:"diagnostico_secundario"`
	DomicilioRisco          bool        `json:"domicilio_risco"`
	ImpedimentoDeslocamento bool        `json:"impedimento_deslocamento"`
	AlimentacaoParenteral   SelectField `json:"alimentacao_parenteral"`
	AspiracaoTraqueo        SelectField `json:"aspiracao_traqueo"`
	VentilacaoMecanica      SelectField `json:"ventilacao_mecanica"`
	MedicacaoParenteral     SelectField `json:"medicacao_parenteral"`
	Banho                   SelectField `json:"banho"`
	Vestir                  SelectField `json:"vestir"`
	HigienePessoal          SelectField `json:"higiene_pessoal"`
	Transferencia           SelectField `json:"transferencia"`
	Continencia             SelectField `json:"continencia"`
	Alimentacao             SelectField `json:"alimentacao"`
	EstadoNutricional       SelectField `json:"estado_nutricional"`
	ViaAlimentacaoMedicacao SelectField `json:"via_alimentacao_medicacao"`
	InternacoesUltimoAno    SelectField `json:"internacoes_ultimo_ano"`
	AspiracaoViasAereas     SelectField `json:"aspiracao_vias_aereas"`
	Lesoes                  SelectField `json:"lesoes"`
	ViaMedicacoes           SelectField `json:"via_medicacoes"`
	ExerciciosVentilatorios SelectField `json:"exercicios_ventilatorios"`
	UsoOxigenoterapia       SelectField `json:"uso_oxigenoterapia"`
	NivelConsciencia        SelectField `json:"nivel_consciencia"`
}

type Observacoes struct {
	Texto string `json:"texto"`
}
