package domain

import "time"

type DadosGerais struct {
	Genero       *string `json:"genero"` // masculino ou feminino
	Complexidade *string `json:"complexidade"`

	Peso   *string `json:"peso"` // TODO: saber se será utilizado
	Altura *string `json:"altura"`
	Idade  *string `json:"idade"`

	NomePaciente          string     `json:"nome_paciente"`
	PrevisaoAlta          *bool      `json:"previsao_alta"`
	DataNascimento        *string    `json:"data_nascimento"`
	CPF                   *string    `json:"cpf"`
	RG                    *string    `json:"rg"`
	NomePai               *string    `json:"nome_pai"`
	NomeMae               *string    `json:"nome_mae"`
	Convenio              *string    `json:"convenio"`
	Hospital              *string    `json:"hospital"`
	DataAdmissao          *string    `json:"dat_admissao"`
	Apartamento           *string    `json:"apartamento"`
	Carteirinha           string     `json:"carteirinha"`
	VencimentoCarteirinha *time.Time `json:"vencimento_carteirinha"`
}
