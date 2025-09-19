package patient

import (
	"context"
	"time"

	"github.com/gui-henri/learning-go/pkg/fhir"
	"github.com/gui-henri/learning-go/pkg/util"
)

type PatientService interface {
	InsertPatient(ctx context.Context, p fhir.Patient) (paciente, error)
}

type patientService struct {
	repository PatientRepository
}

func NewPatientService(r PatientRepository) PatientService {
	return &patientService{
		repository: r,
	}
}

func (s *patientService) InsertPatient(ctx context.Context, p fhir.Patient) (paciente, error) {
	now := time.Now().UTC()
	formattedDate := now.Format("2006-01-02T15:04:05.000Z07:00")

	p.Meta.VersionId = util.Ptr("1")
	p.Meta.LastUpdated = util.Ptr(formattedDate)
	pt, err := s.repository.InsertPatient(p)

	if err != nil {
		return paciente{}, err
	}

	return pt, nil
}
