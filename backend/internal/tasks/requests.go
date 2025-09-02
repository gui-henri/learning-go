package tasks

type GetTaskRequest struct {
	Id int `json:"id"`
}

type GetTaskResponse struct {
	Err  string  `json:"err,omitempty"`
	Task *Tarefa `json:"tarefa"`
}

type InsertTaskRequest struct {
	Descricao string `json:"descricao"`
	Prazo     string `json:"prazo"`
}

type InsertTaskResponse struct {
	Id int `json:"id"`
}

type FinishTaskRequest struct {
	Id int `json:"id"`
}

type FinishTaskResponse struct {
	Task *Tarefa `json:"tarefa"`
}
