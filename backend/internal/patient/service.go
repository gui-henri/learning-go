package patient

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/gui-henri/learning-go/pkg/fhir"
	"github.com/gui-henri/learning-go/pkg/util"
)

type PatientService interface {
	InsertPatient(ctx context.Context, p fhir.Patient) (paciente, error)
	ListPatients(ctx context.Context, count, currentOffset int) (fhir.Bundle, error)
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
		return paciente{}, err
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
