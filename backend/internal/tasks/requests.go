package tasks

type GetTaskRequest struct {
	Id int `json:"id"`
}

type GetTaskResponse struct {
	Err  string  `json:"err,omitempty"`
	Task *Tarefa `json:"tarefa"`
}

type InsertTaskRequest struct {
	Descricao  string `json:"descricao"`
	Prazo      string `json:"prazo"`
	PacienteID string `json:"paciente_id"`
}

type InsertTaskResponse struct {
	Task *Tarefa `json:"tarefa"`
}

type FinishTaskRequest struct {
	Id int `json:"id"`
}

type FinishTaskResponse struct {
	Task *Tarefa `json:"tarefa"`
}

type GetAllIncompleteTasksResponse struct {
	Tasks []Tarefa `json:"tarefas"`
}

type GetAllTasksResponse struct {
	Tasks []Tarefa `json:"tarefas"`
}

type DeleteTaskRequest struct {
	Id int `json:"id"`
}
