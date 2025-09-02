package tasks

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	transport_encoding "github.com/gui-henri/learning-go/pkg/encoding"
)

func NewHttpTransportLayer(mux *http.ServeMux) *http.ServeMux {
	taskService := NewService()

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

	mux.HandleFunc("GET /tasks", getTaskHandler.ServeHTTP)
	mux.HandleFunc("POST /tasks", insertTaskHandler.ServeHTTP)

	return mux
}
