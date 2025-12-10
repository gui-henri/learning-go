package domain

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

func GetCardioRespiratorioFormOptions() FormMetadata {
	return FormMetadata{
		"padrao_respiratorio": []FieldOption{
			{Label: "Eupnéico", Value: "eupneico"},
			{Label: "Dispnéico", Value: "dispneico"},
			{Label: "Bradipnéico", Value: "bradipneico"},
			{Label: "Taquipnéico", Value: "taquipneico"},
		},
		"via_aerea": []FieldOption{
			{Label: "Fisiológica", Value: "fisiologica"},
			{Label: "Orotraqueal", Value: "orotraqueal"},
			{Label: "Traqueóstomo", Value: "traqueostomo"},
		},
		"suporte_ventilatorio": []FieldOption{
			{Label: "Espontânea", Value: "espontanea"},
			{Label: "Ventilação Não Invasiva (VNI)", Value: "modo_vni"},
			{Label: "Assistência Ventilatória Mecânica", Value: "modo_avm"},
		},
		"modo_vni": []FieldOption{
			{Label: "CPAP Simples", Value: "cpap_simples"},
			{Label: "CPAP Automático", Value: "cpap_automatico"},
			{Label: "BiPAP", Value: "bipap"},
		},
		"frequencia_vni": []FieldOption{
			{Label: "Contínuo", Value: "continuo"},
			{Label: "Intermitente", Value: "intermitente"},
		},
		"mascara_vni": []FieldOption{
			{Label: "Nasal", Value: "nasal"},
			{Label: "Facial", Value: "facial"},
			{Label: "Oronasal", Value: "oronasal"},
		},
		"tamanho_mascara": []FieldOption{
			{Label: "P", Value: "P"},
			{Label: "M", Value: "M"},
			{Label: "G", Value: "G"},
		},
		"fonte_oxigenio": []FieldOption{
			{Label: "Cilindro", Value: "cilindro"},
			{Label: "Concentrador", Value: "concentrador"},
		},
		"frequencia_oxigenio": []FieldOption{
			{Label: "Intermitente", Value: "intermitente"},
			{Label: "Contínua", Value: "continua"},
		},
		"modalidade_oxigenio": []FieldOption{
			{Label: "Cateter Nasal", Value: "cateter_nasal"},
			{Label: "Venturi", Value: "venturi"},
		},
	}
}
