package domain

import "time"

type ExameFisico struct {
	EstadoGeral          string           `json:"estado_geral"`         // FIXED
	AvaliacaoLocomotora  string           `json:"avaliacao_locomotora"` // FIXED
	PossuiDreno          bool             `json:"possui_dreno"`
	LocalDreno           string           `json:"local_dreno"`
	DataImplantacaoDreno time.Time        `json:"data_implantacao_dreno"`
	CurativosDreno       string           `json:"curativos_dreno"`
	PossuiAcesso         bool             `json:"possui_acesso"`
	Antimicrobianos      []Antimicrobiano `json:"antimicrobianos"`
}

type Antimicrobiano struct {
	Nome        string `json:"nome_antimicrobiano"`
	DataInicio  string `json:"data_inicio"`
	DataTermino string `json:"data_termino"`
}
