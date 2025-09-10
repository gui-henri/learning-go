package tasks

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	transport_encoding "github.com/gui-henri/learning-go/pkg/encoding"
	"github.com/jackc/pgx/v5"
)

func NewHttpTransportLayer(db *pgx.Conn, mux *http.ServeMux) *http.ServeMux {
	taskRepository := NewTaskRepository(db)
	taskService := NewService(*taskRepository)

	getTaskHandler := httptransport.NewServer(
		makeGetTaskEndpoint(taskService),
		transport_encoding.EncodeRequest[GetTaskRequest],
		transport_encoding.EncodeResponse,
		httptransport.ServerErrorEncoder(transport_encoding.EncodeError),
	)

	insertTaskHandler := httptransport.NewServer(
		makeInsertTaskEndpoint(taskService),
		transport_encoding.EncodeRequest[InsertTaskRequest],
		transport_encoding.EncodeResponse,
		httptransport.ServerErrorEncoder(transport_encoding.EncodeError),
	)

	finishTaskHandler := httptransport.NewServer(
		makeFinishTaskEndpoint(taskService),
		transport_encoding.EncodeRequest[FinishTaskRequest],
		transport_encoding.EncodeResponse,
		httptransport.ServerErrorEncoder(transport_encoding.EncodeError),
	)

	getAllIncompleteRequests := httptransport.NewServer(
		makeGetAllIncompleteEndpoint(taskService),
		transport_encoding.EncodeNoBodyRequest,
		transport_encoding.EncodeResponse,
		httptransport.ServerErrorEncoder(transport_encoding.EncodeError),
	)

	getAllRequests := httptransport.NewServer(
		makeGetAllEndpoint(taskService),
		transport_encoding.EncodeNoBodyRequest,
		transport_encoding.EncodeResponse,
		httptransport.ServerErrorEncoder(transport_encoding.EncodeError),
	)

	deleteTask := httptransport.NewServer(
		makeDeleteTaskEndpoint(taskService),
		transport_encoding.EncodeNoBodyRequest,
		transport_encoding.EncodeResponse,
		httptransport.ServerErrorEncoder(transport_encoding.EncodeError),
	)

	mux.HandleFunc("GET /tasks", getTaskHandler.ServeHTTP)
	mux.HandleFunc("POST /tasks", insertTaskHandler.ServeHTTP)
	mux.HandleFunc("POST /tasks/finish", finishTaskHandler.ServeHTTP)
	mux.HandleFunc("GET /tasks/all", getAllRequests.ServeHTTP)
	mux.HandleFunc("GET /tasks/all-unfinished", getAllIncompleteRequests.ServeHTTP)
	mux.HandleFunc("DELETE /tasks", deleteTask.ServeHTTP)

	return mux
}
