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
