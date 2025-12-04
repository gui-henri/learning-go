package domain

import "time"

type Nutricional struct {
	AlimentacaoOral       bool      `json:"alimentacao_oral"`
	AlimentacaoEnteral    bool      `json:"alimentacao_enteral"`
	AlimentacaoParenteral bool      `json:"alimentacao_parenteral"`
	Sonda                 bool      `json:"sonda"`
	Botton                bool      `json:"botton"`
	Suplemento            bool      `json:"suplemento"`
	ViaEnteral            string    `json:"via_enteral"`
	ViaParenteral         string    `json:"via_parenteral"`
	AdaptadorSonda        string    `json:"adaptador_sonda"`
	RestricaoAlimentar    string    `json:"restricao_alimentar"`
	TipoDieta             string    `json:"tipo_dieta"`
	FormaAdministracao    string    `json:"forma_administracao"`
	MarcaBomba            string    `json:"marca_bomba"`
	DataUltimaTroca       time.Time `json:"data_ultima_troca"`
	Gavando               string    `json:"gavando"`
	VolumeDiario          string    `json:"volume_diario"`
	QualSuplemento        string    `json:"qual_suplemento"`
}
