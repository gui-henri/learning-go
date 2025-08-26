package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var listaTarefas := make(Tarefa, 0)

type NovaTarefaInput struct {
	Descricao string `json:"descricao"`
	Prazo     string `json:"prazo"`
}

type Tarefa struct {
	Id Int	`json:"id"`
	Descricao string `json:"descricao"`
	Prazo     string `json:"prazo"`
	Concluida bool   `json:"concluida"`
	CriadaEm  string `json:"criada-em"`
}

type NovaTarefaOutput struct {
	Id int
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /task", func(w http.ResponseWriter, r *http.Request) {

		var tarefa Tarefa
		var novaTarefaInput NovaTarefaInput
		var novaTarefaOutput NovaTarefaOutput
		err := json.NewDecoder(r.Body).Decode(&novaTarefaInput)

		if err != nil {
			http.Error(w, "Requisição não pôde ser lida", http.StatusBadRequest)
			return
		}

		tarefa = Tarefa{
			Id: len(listaTarefas)
			Descricao: novaTarefaOutput.Descricao,
			Prazo:     novaTarefaInput.Prazo,
			Concluida: false,
			CriadaEm:  time.Now().Format(time.RFC850),
		}

		listaTarefas := append(listaTarefas, tarefa)

		novaTarefaOutput = NovaTarefaOutput{
			Id: tarefa.Id
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(novaTarefaOutput)

	})

	mux.HandleFunc("GET /task/{id}/", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "handling task with id=%v\n", id)
	})

	http.ListenAndServe("localhost:8090", mux)
}
