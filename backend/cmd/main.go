package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gui-henri/learning-go/db"
	"github.com/gui-henri/learning-go/internal/avaliation"
	"github.com/gui-henri/learning-go/pkg/middleware"
)

func main() {
	db.Connect()
	defer db.DB.Close(context.Background())

	mux := http.NewServeMux()
	avaliation.NewHttpTransportLayer(db.DB, mux, "http://gotenberg:3000", "./templates/pep/interne.html")
	handler := middleware.NoCors(mux)
	fmt.Println("Starting server at 8090")
	log.Fatal(http.ListenAndServe(":8090", handler))
}
