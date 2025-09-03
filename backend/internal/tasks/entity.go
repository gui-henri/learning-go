package tasks

import "time"

type Tarefa struct {
	Id        int       `json:"id"`
	Descricao string    `json:"descricao"`
	Prazo     string    `json:"prazo"`
	Concluida bool      `json:"concluida"`
	CriadaEm  time.Time `json:"criada-em"`
}
