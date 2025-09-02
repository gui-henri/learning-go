package main

import (
	"log"
	"net/http"

	"github.com/gui-henri/learning-go/tasks"
)

func main() {
	mux := http.NewServeMux()

	tasks.NewHttpTransportLayer(mux)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
