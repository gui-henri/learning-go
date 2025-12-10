package domain

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

func GetNutricionalFormOptions() FormMetadata {
	return FormMetadata{
		"via_enteral": []FieldOption{
			{Label: "Gastrostomia", Value: "gastrostomia"},
			{Label: "Jejunostomia", Value: "jejunostomia"},
			{Label: "Nasoenteral", Value: "nasoenteral"},
		},
		"via_parenteral": []FieldOption{
			{Label: "Central", Value: "central"},
			{Label: "Periférica", Value: "periferica"},
		},
		"adaptador_sonda": []FieldOption{
			{Label: "Torneirinha 3 vias", Value: "torneirinha_3_vias"},
			{Label: "Polifix 2 vias", Value: "polifix_2_vias"},
			{Label: "Outros", Value: "outros"},
		},
		"tipo_dieta": []FieldOption{
			{Label: "Artesanal", Value: "artesanal"},
			{Label: "Industrializada", Value: "industrializada"},
			{Label: "Mista", Value: "mista"},
		},
		"forma_administracao": []FieldOption{
			{Label: "Gravitacional", Value: "gravitacional"},
			{Label: "Bomba de Infusão", Value: "bomba_infusao"},
		},
		"marcas_bomba": []FieldOption{
			{Label: "Kangaroo", Value: "Kangaroo"},
			{Label: "Halyard", Value: "Halyard"},
			{Label: "Kimberly-Clark", Value: "Kimberly-Clark"},
			{Label: "Wilson Cook", Value: "Wilson Cook"},
			{Label: "Folley", Value: "Folley"},
		},
	}
}
