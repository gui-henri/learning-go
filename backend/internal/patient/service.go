package patient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"

	apperrors "github.com/gui-henri/learning-go/pkg/errors"
	"github.com/gui-henri/learning-go/pkg/fhir"
	"github.com/gui-henri/learning-go/pkg/util"
)

type PatientService interface {
	InsertPatient(ctx context.Context, p fhir.Patient) (paciente, error)
	ListPatients(ctx context.Context, count, currentOffset int) (fhir.Bundle, error)
	UpdatePatient(ctx context.Context, id string, p fhir.Patient) (paciente, error)
}

type patientService struct {
	repository patientRepository
}

func NewPatientService(r patientRepository) *patientService {
	return &patientService{
		repository: r,
	}
}

func buildLinkURL(baseURL string, count, offset int) string {
	u, _ := url.Parse(baseURL)
	q := u.Query()

	q.Set("_count", strconv.Itoa(count))
	q.Set("_offset", strconv.Itoa(offset))

	u.RawQuery = q.Encode()
	return u.String()
}

func (s *patientService) InsertPatient(ctx context.Context, p fhir.Patient) (paciente, error) {
	now := time.Now().UTC()
	formattedDate := now.Format("2006-01-02T15:04:05.000Z07:00")

	if p.Meta == nil {
		p.Meta = &fhir.Meta{}
	}

	p.Meta.VersionId = util.Ptr("1")
	p.Meta.LastUpdated = util.Ptr(formattedDate)
	pt, err := s.repository.InsertPatient(p)

	if err != nil {
		return paciente{}, apperrors.InvalidInput
	}

	return pt, nil
}

func (s *patientService) ListPatients(ctx context.Context, count, currentOffset int) (fhir.Bundle, error) {
	pPage, err := s.repository.ListPatients(count, currentOffset)
	if err != nil {
		return fhir.Bundle{}, err
	}

	var links []fhir.BundleLink

	links = append(links, fhir.BundleLink{
		Relation: "self",
		Url:      buildLinkURL("http://localhost:8090/Patient", count, currentOffset),
	})

	if currentOffset+count < pPage.Total {
		links = append(links, fhir.BundleLink{
			Relation: "next",
			Url:      buildLinkURL("http://localhost:8090/Patient", count, currentOffset+count),
		})
	}

	if currentOffset > 0 {
		prevOffset := max(currentOffset-count, 0)
		links = append(links, fhir.BundleLink{
			Relation: "previous",
			Url:      buildLinkURL("http://localhost:8090/Patient", count, prevOffset),
		})
	}

	var entries []fhir.BundleEntry
	for _, p := range pPage.Patients {
		entries = append(entries, fhir.BundleEntry{
			Resource: p.ResourceJSON,
		})
	}

	bundle := fhir.Bundle{
		Type:  fhir.BundleTypeSearchset,
		Total: &pPage.Total,
		Link:  links,
		Entry: entries,
	}

	return bundle, nil
}

func (s *patientService) UpdatePatient(ctx context.Context, id string, p fhir.Patient) (paciente, error) {
	existingPatient, err := s.repository.GetPatient(id)
	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			return paciente{}, fmt.Errorf("paciente com id '%s' não encontrado: %w", id, apperrors.NotFound)
		}
		return paciente{}, fmt.Errorf("falha ao buscar paciente: %w", err)
	}

	var fhirPatient fhir.Patient
	if err := json.Unmarshal(existingPatient.ResourceJSON, &fhirPatient); err != nil {
		return paciente{}, fmt.Errorf("falha ao desserializar paciente existente: %w", err)
	}

	// Preserva o ID original e o Meta
	p.Id = fhirPatient.Id
	p.Meta = fhirPatient.Meta

	if p.Meta == nil {
		p.Meta = &fhir.Meta{}
	}

	newVersion, err := incrementVersion(p.Meta.VersionId)
	if err != nil {
		return paciente{}, fmt.Errorf("versão inválida no metadado do paciente: %w", err)
	}
	p.Meta.VersionId = util.Ptr(newVersion)
	p.Meta.LastUpdated = util.Ptr(time.Now().UTC().Format("2006-01-02T15:04:05.000Z07:00"))

	pt, err := s.repository.UpdatePatient(id, p)
	if err != nil {
		return paciente{}, fmt.Errorf("falha ao atualizar paciente no repositório: %w", err)
	}

	return pt, nil
}

func incrementVersion(currentVersion *string) (string, error) {
	if currentVersion == nil || *currentVersion == "" {
		return "1", nil
	}

	v, err := strconv.Atoi(*currentVersion)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(v + 1), nil
}
