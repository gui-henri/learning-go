package patient

import (
	"time"

	"github.com/go-playground/validator/v10"
	apperrors "github.com/gui-henri/learning-go/pkg/errors"
	"github.com/gui-henri/learning-go/pkg/fhir"
)

type paciente struct {
	Id                     string    `json:"id"`
	InternalId             int       `json:"internal_id"`
	Version                int       `json:"version" validate:"required"`
	LastUpdated            time.Time `json:"last_updated" validate:"required"`
	Active                 bool      `json:"concluida"`
	Gender                 string    `json:"gender"`
	BirthDate              *string
	Deceased               bool
	DeceasedDatetime       *string
	MaritalStatus          string
	MaritalStatusDisplay   string
	ManagingOrganizationId string
	NameFamily             string
	NameGiven1             string
	NameGiven2             string
	Upid                   string
	CreatedAt              time.Time `json:"criada_em" validate:"lte"`
	ResourceJson           []byte
}

func NewPaciente(p fhir.Patient) (paciente, error) {
	j, err := p.MarshalJSON()
	if err != nil {
		return paciente{}, err
	}
	pt := paciente{
		Version:                1,
		LastUpdated:            time.Now(),
		Active:                 true,
		Gender:                 p.Gender.Code(),
		BirthDate:              p.BirthDate,
		Deceased:               *p.DeceasedBoolean,
		DeceasedDatetime:       p.DeceasedDateTime,
		MaritalStatus:          *p.MaritalStatus.Id,
		MaritalStatusDisplay:   *p.MaritalStatus.Text,
		ManagingOrganizationId: *p.ManagingOrganization.Id,
		NameFamily:             *p.Name[0].Family,
		NameGiven1:             p.Name[0].Given[0],
		NameGiven2:             p.Name[0].Given[1],
		Upid:                   *p.Identifier[0].Id,
		CreatedAt:              time.Now(),
		ResourceJson:           j,
	}

	validate := validator.New()

	err = validate.Struct(pt)

	if err != nil {
		return paciente{}, apperrors.InvalidInput
	}

	return pt, nil
}
