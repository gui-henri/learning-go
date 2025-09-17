package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gui-henri/learning-go/db"
	"github.com/gui-henri/learning-go/internal/tasks"
	"github.com/gui-henri/learning-go/pkg/fhir"
	"github.com/gui-henri/learning-go/pkg/middleware"
)

func Ptr[T any](v T) *T {
	return &v
}

func main() {
	p := fhir.Patient{
		Id: Ptr("1"),
		Meta: &fhir.Meta{
			Id:          Ptr("1"),
			VersionId:   Ptr("1"),
			LastUpdated: Ptr(time.Now().String()),
			Source:      Ptr("https://interne.com.br"),
		},
		Text: &fhir.Narrative{
			Id:     Ptr("1"),
			Status: fhir.NarrativeStatusAdditional,
		},
		Name: []fhir.HumanName{
			{
				Id:     Ptr("1"),
				Text:   Ptr("Guilherme Henrique Freitas da Silva"),
				Family: Ptr("Silva"),
				Given:  []string{"Guilherme", "Henrique", "Freitas"},
			},
		},
		BirthDate:       Ptr("24/09/2002"),
		Gender:          Ptr(fhir.AdministrativeGenderMale),
		Active:          Ptr(true),
		DeceasedBoolean: Ptr(false),
		Address: []fhir.Address{
			{
				City:    Ptr("Vitória de Santo Antão"),
				State:   Ptr("Pernambuco"),
				Country: Ptr("Brasil"),
				Type:    Ptr(fhir.AddressTypePhysical),
				Id:      Ptr("1"),
			},
		},
		Communication: []fhir.PatientCommunication{
			{
				Id: Ptr("1"),
				Language: fhir.CodeableConcept{
					Id:   Ptr("1"),
					Text: Ptr("pt-br"),
					Coding: []fhir.Coding{
						{
							Id:           Ptr("1"),
							UserSelected: Ptr(true),
							Code:         Ptr("pt-br"),
							Version:      Ptr("1"),
						},
					},
				},
				Preferred: Ptr(true),
			},
		},
	}

	jsonBytes, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}

	fmt.Println(string(jsonBytes))

	db.Connect()
	defer db.DB.Close(context.Background())

	mux := http.NewServeMux()
	tasks.NewHttpTransportLayer(db.DB, mux)
	handler := middleware.NoCors(mux)

	fmt.Println("Starting server at 8090")
	log.Fatal(http.ListenAndServe(":8090", handler))
}
