package domain

type Cuidadores struct {
	MedicoSolicitante  string `json:"medico_solicitante"`
	ContatoMedico      string `json:"contato_medico"`
	PossuiCuidador     bool   `json:"possui_cuidador"`
	NomeCuidador       string `json:"nome_cuidador"`
	ContatoCuidador    string `json:"contato_cuidador"`
	TurnoCuidador      string `json:"turno_cuidador"`
	PrecisaTreinamento bool   `json:"precisa_treinamento"`
	ObsTreinamento     string `json:"obs_treinamento"`
}
