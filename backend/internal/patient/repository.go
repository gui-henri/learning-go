package patient

import (
	"context"
	"fmt"

	"github.com/gui-henri/learning-go/db"
	"github.com/gui-henri/learning-go/pkg/errors"
	"github.com/gui-henri/learning-go/pkg/fhir"
)

type PatientRepository interface {
	InsertPatient(p fhir.Patient) (paciente, error)
	ListPatients(count, offset int) (pagedPatientResult, error)
	UpdatePatient(id string, p fhir.Patient) (paciente, error)
	GetPatient(id string) (paciente, error)
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

	_, err = p.db.Exec(
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

	if err != nil {
		fmt.Println("[ERR] Database error happened: ", err)
		return paciente{}, errors.InvalidInput
	}

	return patient, nil

}

type pagedPatientResult struct {
	Patients []paciente
	Total    int
}

func (p *patientRepository) ListPatients(count, offset int) (pagedPatientResult, error) {

	sql := `
		SELECT
            id,
            last_updated,
            active,
            gender,
            birth_date,
            deceased,
            full_name,
            cpf,
            created_at,
            resource_json,
            COUNT(*) OVER() AS total_count 
        FROM 
            patient 
        ORDER BY 
            id 
        LIMIT $1 OFFSET $2
	`

	if count > 100 {
		return pagedPatientResult{}, errors.InvalidInput
	}

	rows, err := p.db.Query(context.Background(), sql, count, offset)

	if err != nil {
		return pagedPatientResult{}, err
	}

	defer rows.Close()

	patients := []paciente{}

	var totalCount int = 0

	for rows.Next() {
		var pt paciente

		err := rows.Scan(
			&pt.ID,
			&pt.LastUpdated,
			&pt.Active,
			&pt.Gender,
			&pt.BirthDate,
			&pt.Deceased,
			&pt.FullName,
			&pt.CPF,
			&pt.CreatedAt,
			&pt.ResourceJSON,
			&totalCount,
		)

		if err != nil {
			return pagedPatientResult{}, errors.InvalidInput
		}
		patients = append(patients, pt)
	}

	return pagedPatientResult{
		Patients: patients,
		Total:    totalCount,
	}, nil
}

func (p *patientRepository) UpdatePatient(id string, pt fhir.Patient) (paciente, error) {
	sql := `
		UPDATE patient
		SET
			last_updated = $1,
			gender = $2,
			birth_date = $3,
			full_name = $4,
			cpf = $5,
			resource_json = $6,
			deceased = $7,
			active = $8
		WHERE id = $9
	`

	patient, err := NewPaciente(pt)
	if err != nil {
		return paciente{}, err
	}

	_, err = p.db.Exec(
		context.Background(),
		sql,
		patient.LastUpdated,
		patient.Gender,
		patient.BirthDate,
		patient.FullName,
		patient.CPF,
		patient.ResourceJSON,
		patient.Deceased,
		patient.Active,
		id,
	)

	if err != nil {
		fmt.Println("[ERR] Database error happened: ", err)
		return paciente{}, errors.InvalidInput
	}

	return patient, nil
}

func (p *patientRepository) GetPatient(id string) (paciente, error) {
	sql := `
		SELECT
			id,
			last_updated,
			active,
			gender,
			birth_date,
			deceased,
			full_name,
			cpf,
			created_at,
			resource_json
		FROM
			patient
		WHERE id = $1
	`

	row := p.db.QueryRow(context.Background(), sql, id)

	var pt paciente

	err := row.Scan(
		&pt.ID,
		&pt.LastUpdated,
		&pt.Active,
		&pt.Gender,
		&pt.BirthDate,
		&pt.Deceased,
		&pt.FullName,
		&pt.CPF,
		&pt.CreatedAt,
		&pt.ResourceJSON,
	)

	if err != nil {
		return paciente{}, err
	}

	return pt, nil
}
