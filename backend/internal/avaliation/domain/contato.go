package domain

type Contato struct {
	TelefoneResidencial string        `json:"telefone_residencial"`
	TelefonePaciente    string        `json:"telefone_paciente"`
	EmailPaciente       string        `json:"email_paciente"`
	Responsaveis        []Responsavel `json:"responsaveis"`
}

type Responsavel struct {
	Nome              string `json:"nome"`
	Parentesco        string `json:"parentesco"`
	Email             string `json:"email"`
	Telefone          string `json:"telefone"`
	FormaContato      string `json:"forma_contato"`
	NumeroPrioritario bool   `json:"numero_prioritario"`
}
