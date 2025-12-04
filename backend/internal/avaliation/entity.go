package avaliation

import (
	"encoding/json"
	"time"

	"github.com/gui-henri/learning-go/internal/avaliation/domain"
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
	DadosGerais domain.DadosGerais `json:"dadosGerais"`
	// Endereco           Endereco           `json:"endereco"`
	// Contato            Contato            `json:"contato"`
	// Cuidadores         Cuidadores         `json:"cuidadores"`
	// Seguranca          Seguranca          `json:"seguranca"`
	// HistoricoClinico   Historico          `json:"historicoClinico"`
	// ExameFisico        ExameFisico        `json:"exameFisico"`
	// CardioRespiratorio CardioRespiratorio `json:"cardioRespiratorio"`
	// Nutricional        Nutricional        `json:"nutricional"`
	// Eliminacoes        Eliminacoes        `json:"eliminacoes"`
	// CondicoesPele      CondicoesPele      `json:"condicoesPele"`
	// Score              Score              `json:"score"`
	// Observacoes        Observacoes        `json:"observacoes"`
}

// --- DADOS GERAIS ---

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
	EstadoGeral          string           `json:"estado_geral"`         // FIXED
	AvaliacaoLocomotora  string           `json:"avaliacao_locomotora"` // FIXED
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
	PadraoRespiratorio  string `json:"padrao_respiratorio"`
	ViaAerea            string `json:"via_aerea"`
	NumeroTubo          string `json:"numero_tubo"`
	DataOrotraqueal     string `json:"data_orotraqueal"`
	SuporteVentilatorio string `json:"suporte_ventilatorio"`
	// VNI
	ModoVNI           string `json:"modo_vni"`
	FrequenciaVNI     string `json:"frequencia_vni"`
	Fio2VNI           string `json:"fio2_vni"`
	MascaraVNI        string `json:"mascara_vni"`
	TamanhoMascaraVNI string `json:"tamanho_mascara_vni"`
	// AVM
	ModoAVM           string `json:"modo_avm"`
	FrequenciaAVM     string `json:"frequencia_avm"`
	Fio2AVM           string `json:"fio2_avm"`
	MascaraAVM        string `json:"mascara_avm"`
	TamanhoMascaraAVM string `json:"tamanho_mascara_avm"`
	// Oxigenoterapia
	UsaOxigenioterapia bool   `json:"usa_oxigenioterapia"`
	FonteOxigenio      string `json:"fonte_oxigenio"`
	FrequenciaOxigenio string `json:"frequencia_oxigenio"`
	VezesDiaOxigenio   string `json:"vezes_dia_oxigenio"`
	ModalidadeOxigenio string `json:"modalidade_oxigenio"`
	FluxoOxigenio      string `json:"fluxo_oxigenio"`
}

type Nutricional struct {
	AlimentacaoOral       bool   `json:"alimentacao_oral"`
	AlimentacaoEnteral    bool   `json:"alimentacao_enteral"`
	AlimentacaoParenteral bool   `json:"alimentacao_parenteral"`
	Sonda                 bool   `json:"sonda"`
	Botton                bool   `json:"botton"`
	Suplemento            bool   `json:"suplemento"`
	ViaEnteral            string `json:"via_enteral"`
	ViaParenteral         string `json:"via_parenteral"`
	AdaptadorSonda        string `json:"adaptador_sonda"`
	RestricaoAlimentar    string `json:"restricao_alimentar"`
	TipoDieta             string `json:"tipo_dieta"`
	FormaAdministracao    string `json:"forma_administracao"`
	MarcaBomba            string `json:"marca_bomba"`
	DataUltimaTroca       string `json:"data_ultima_troca"`
	Gavando               string `json:"gavando"`
	VolumeDiario          string `json:"volume_diario"`
	QualSuplemento        string `json:"qual_suplemento"`
}

type Eliminacoes struct {
	FuncaoIntestinal              string `json:"funcao_intestinal"`
	ViaEvacuacaoFisiologica       bool   `json:"via_evacuacao_fisiologica"`
	ViaEvacuacaoEstomia           bool   `json:"via_evacuacao_estomia"`
	TipoEstomia                   string `json:"tipo_estomia"`
	PlacaEstomia                  string `json:"placa_estomia"`
	BolsaEstomia                  string `json:"bolsa_estomia"`
	Diurese                       string `json:"diurese"`
	NumSndDiurese                 string `json:"num_snd_diurese"`
	UsaFralda                     bool   `json:"usa_fralda"`
	TrocasFraldaDia               string `json:"trocas_fralda_dia"`
	SVA                           bool   `json:"sva"`
	CateterismoIntermitente       bool   `json:"cateterismo_intermitente"`
	VezesCateterismoDia           string `json:"vezes_cateterismo_dia"`
	SVD                           bool   `json:"svd"`
	NumSVD                        string `json:"num_svd"`
	ViasSVD                       string `json:"vias_svd"`
	IrrigacaoSVD                  bool   `json:"irrigacao_svd"`
	DataTrocaSVD                  string `json:"data_troca_svd"`
	VolumeDiureseSVD              string `json:"volume_diurese_svd"`
	Cistostomia                   bool   `json:"cistostomia"`
	NumCistostomia                string `json:"num_cistostomia"`
	ViasCistostomia               string `json:"vias_cistostomia"`
	IrrigacaoCistostomia          bool   `json:"irrigacao_cistostomia"`
	QtdIrrigacaoCistostomia       string `json:"qtd_irrigacao_cistostomia"`
	MedicacaoIrrigacaoCistostomia bool   `json:"medicacao_irrigacao_cistostomia"`
	NomeMedicacaoCistostomia      string `json:"nome_medicacao_cistostomia"`
	VolumeDiureseCistostomia      string `json:"volume_diurese_cistostomia"`
	Preservativo                  bool   `json:"preservativo"`
	VolumeDiuresePreservativo     string `json:"volume_diurese_preservativo"`
}

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

type Score struct {
	DiagnosticoPrimario     string `json:"diagnostico_primario"`
	DiagnosticoSecundario   string `json:"diagnostico_secundario"`
	DomicilioRisco          bool   `json:"domicilio_risco"`
	ImpedimentoDeslocamento bool   `json:"impedimento_deslocamento"`
	AlimentacaoParenteral   string `json:"alimentacao_parenteral"`
	AspiracaoTraqueo        string `json:"aspiracao_traqueo"`
	VentilacaoMecanica      string `json:"ventilacao_mecanica"`
	MedicacaoParenteral     string `json:"medicacao_parenteral"`
	Banho                   string `json:"banho"`
	Vestir                  string `json:"vestir"`
	HigienePessoal          string `json:"higiene_pessoal"`
	Transferencia           string `json:"transferencia"`
	Continencia             string `json:"continencia"`
	Alimentacao             string `json:"alimentacao"`
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

type Observacoes struct {
	Texto string `json:"texto"`
}
