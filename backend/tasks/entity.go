package tasks

type Tarefa struct {
	Id        int    `json:"id"`
	Descricao string `json:"descricao"`
	Prazo     string `json:"prazo"`
	Concluida bool   `json:"concluida"`
	CriadaEm  string `json:"criada-em"`
}
