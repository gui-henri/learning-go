package patient

import "github.com/gui-henri/learning-go/pkg/fhir"

type InsertPatientRequest struct {
	Patient fhir.Patient `json:"paciente"`
}

type InsertPatientResponse struct {
	Patient *paciente `json:"paciente"`
}

type ListPatientResponse struct {
	Data []paciente `json:"data"`
}
