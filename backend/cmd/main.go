package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gui-henri/learning-go/internal/tasks"
)

func main() {
	mux := http.NewServeMux()

	tasks.NewHttpTransportLayer(mux)
	fmt.Println("Starting server at 8090")
	log.Fatal(http.ListenAndServe(":8090", mux))
}
