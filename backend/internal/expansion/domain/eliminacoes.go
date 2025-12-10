package domain

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

func GetEliminacoesFormOptions() FormMetadata {
	return FormMetadata{
		"funcao_intestinal": []FieldOption{
			{Label: "Normal", Value: "normal"},
			{Label: "Constipado", Value: "constipado"},
			{Label: "Com diarreia", Value: "com_diarreia"},
		},
		"tipo_estomia": []FieldOption{
			{Label: "Ileostomia", Value: "ileostomia"},
			{Label: "Colostomia", Value: "colostomia"},
		},
		"diurese": []FieldOption{
			{Label: "Espontânea", Value: "espontanea"},
			{Label: "Induzida", Value: "induzida"},
		},
		"volume_diurese": []FieldOption{
			{Label: "Normal", Value: "normal"},
			{Label: "Anúria", Value: "anuria"},
			{Label: "Oligúria", Value: "oliguria"},
			{Label: "Poliúria", Value: "poliuria"},
		},
		"vias_sonda": []FieldOption{
			{Label: "2 Vias", Value: "2"},
			{Label: "3 Vias", Value: "3"},
		},
	}
}
