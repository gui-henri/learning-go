package patient

import (
	"context"

	"github.com/gui-henri/learning-go/db"
	"github.com/gui-henri/learning-go/pkg/fhir"
)

type PatientRepository interface {
	InsertPatient(p fhir.Patient) (paciente, error)
	ListPatients() ([]paciente, error)
}

type patientRepository struct {
	db db.IDB
}

func NewPatientRepository(db db.IDB) *patientRepository {
	return &patientRepository{
		db: db,
	}
}

func (p *patientRepository) InsertPatient(pt fhir.Patient) (paciente, error) {
	sql := `
		INSERT INTO patient (id, last_updated, gender, birth_date, full_name, cpf, resource_json)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	patient, err := NewPaciente(pt)
	if err != nil {
		return paciente{}, err
	}

	p.db.QueryRow(
		context.Background(),
		sql,
		patient.ID,
		patient.LastUpdated,
		patient.Gender,
		patient.BirthDate,
		patient.FullName,
		patient.CPF,
		patient.ResourceJSON,
	)

	return patient, nil

}

func (p *patientRepository) ListPatients() (paciente, error) {
	sql := `
		SELECT * FROM patient LIMIT 100
	`
}
