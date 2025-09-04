package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gui-henri/learning-go/db"
	"github.com/gui-henri/learning-go/internal/tasks"
	"github.com/gui-henri/learning-go/pkg/middleware"
)

func main() {
	db.Connect()
	mux := http.NewServeMux()

	tasks.NewHttpTransportLayer(db.DB, mux)

	handler := middleware.NoCors(mux)

	fmt.Println("Starting server at 8090")
	log.Fatal(http.ListenAndServe(":8090", handler))
}
