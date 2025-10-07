package patient

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gui-henri/learning-go/pkg/errors"
	"github.com/gui-henri/learning-go/pkg/fhir"
	"github.com/gui-henri/learning-go/pkg/util"
)

type paciente struct {
	InternalID   int64           `db:"internal_id" json:"internal_id"` //criado no banco
	ID           string          `db:"id" json:"id"`
	LastUpdated  *time.Time      `db:"last_updated" json:"last_updated,omitempty"`
	Active       bool            `db:"active" json:"active"` //forçado true
	Gender       *string         `db:"gender" json:"gender,omitempty"`
	BirthDate    *time.Time      `db:"birth_date" json:"birth_date,omitempty"`
	Deceased     *bool           `db:"deceased" json:"deceased,omitempty"` //forçado false
	FullName     *string         `db:"full_name" json:"full_name,omitempty"`
	CPF          *string         `db:"cpf" json:"cpf,omitempty"`
	CreatedAt    time.Time       `db:"created_at" json:"created_at"` //criado no banco
	ResourceJSON json.RawMessage `db:"resource_json" json:"resource_json"`
}

type Paciente struct {
	ID       string  `db:"id" json:"id"`
	FullName *string `db:"full_name" json:"full_name,omitempty"`
}

const (
	SystemCPFOficial = "http://hl7.org.br/fhir/r4/sid/CPF"
	SystemCPFAntigo  = "urn:cpf"
)

func NewPaciente(p fhir.Patient) (paciente, error) {

	id := uuid.NewString()
	p.Id = &id

	j, err := p.MarshalJSON()
	if err != nil {
		return paciente{}, err
	}

	var birthDateTime *time.Time
	var genderCode *string

	if p.BirthDate != nil && *p.BirthDate != "" {
		t, err := time.Parse("2006-01-02", *p.BirthDate)
		if err != nil {
			return paciente{}, err
		}
		birthDateTime = &t
	}

	if p.Gender != nil {
		code := p.Gender.Code()
		genderCode = &code
	}

	// Campos omitidos serão criados pelo banco

	fn, err := proccessFullName(p.Name)

	if err != nil {
		return paciente{}, err
	}

	cpf, err := getPatientCPF(&p)

	if err != nil {
		return paciente{}, err
	}

	deceased := false
	active := true

	if p.Active != nil {
		active = *p.Active
	}

	if p.DeceasedBoolean != nil && *p.DeceasedBoolean {
		deceased = true
		active = false
	} else if p.DeceasedDateTime != nil && *p.DeceasedDateTime != "" {
		deceased = true
		active = false
	}

	pt := paciente{
		ID:           id,
		LastUpdated:  util.Ptr(time.Now()),
		CreatedAt:    time.Now(),
		Active:       active,
		Gender:       genderCode,
		BirthDate:    birthDateTime,
		FullName:     util.Ptr(fn),
		CPF:          util.Ptr(cpf),
		Deceased:     &deceased,
		ResourceJSON: j,
	}

	return pt, nil
}

func proccessFullName(n []fhir.HumanName) (string, error) {

	if len(n) == 0 {
		return "", errors.InvalidInput
	}

	// caso ele tenha apenas um nome
	if len(n) == 1 {
		val, err := proccessSingleName(n[0])

		if err != nil {
			return "", err
		}

		return val, nil
	}

	// caso ele tenha múltiplos nomes registrados,
	// Procuramos um com o nome "usual".
	// Caso não exista, o primeiro nome encontrado é usado

	var officialName *fhir.HumanName
	var usualName *fhir.HumanName

	for idx, name := range n {
		if name.Use == nil {
			continue
		}
		if name.Use.Code() == "usual" && usualName == nil {
			usualName = &n[idx]
		}
		if name.Use.Code() == "official" && officialName == nil {
			officialName = &n[idx]
		}
	}

	if usualName != nil {
		val, err := proccessSingleName(*usualName)
		if err != nil {
			return "", err
		}
		return val, nil
	}

	if officialName != nil {
		val, err := proccessSingleName(*officialName)
		if err != nil {
			return "", err
		}
		return val, nil
	}

	val, err := proccessSingleName(n[0])
	if err != nil {
		return "", err
	}
	return val, nil
}

func proccessSingleName(n fhir.HumanName) (string, error) {
	if n.Text != nil && *n.Text != "" {
		return *n.Text, nil
	}

	parts := []string{}

	for _, given := range n.Given {
		if given != "" {
			parts = append(parts, given)
		}
	}

	if n.Family != nil && *n.Family != "" {
		parts = append(parts, *n.Family)
	}

	if len(parts) == 0 {
		return "", errors.InvalidInput
	}

	return strings.Join(parts, " "), nil
}

func getPatientCPF(p *fhir.Patient) (string, error) {
	if len(p.Identifier) == 0 {
		return "CPF não encontrado", errors.InvalidInput
	}

	for _, identifier := range p.Identifier {
		if identifier.System == nil || identifier.Value == nil || *identifier.Value == "" {
			continue
		}

		systemURI := *identifier.System
		if systemURI == SystemCPFOficial || systemURI == SystemCPFAntigo {
			cpf := *identifier.Value
			cpf = strings.ReplaceAll(cpf, "-", "")
			cpf = strings.ReplaceAll(cpf, " ", "")
			cpf = strings.ReplaceAll(cpf, ".", "")
			fmt.Println("[INFO] CPF: ", cpf)
			fmt.Println("[INFO] CPF len: ", len(cpf))

			if len(cpf) != 11 {
				return "CPF inválido", errors.InvalidInput
			}

			return cpf, nil
		}
	}

	return "CPF não encontrado", errors.InvalidInput
}

// EXEMPLO DE PACIENTE fhir
// p := fhir.Patient{
// 	Id: Ptr("<uuid>"),
// 	Meta: &fhir.Meta{
// 		VersionId:   Ptr("1"),
// 		LastUpdated: Ptr(time.Now().String()), TODO: usar ISO 8601
// 		Source:      Ptr("https://interne.com.br"),
// 	},
// 	Text: &fhir.Narrative{
// 		Status: fhir.NarrativeStatusAdditional,
// 	},
// 	Name: []fhir.HumanName{
// 		{
// 			Text:   Ptr("Guilherme Henrique Freitas da Silva"),
// 			Family: Ptr("Silva"),
// 			Given:  []string{"Guilherme", "Henrique", "Freitas"},
// 		},
// 	},
// 	BirthDate:       Ptr("24/09/2002"), TODO: usar ano/mes/dia
// 	Gender:          Ptr(fhir.AdministrativeGenderMale),
// 	Active:          Ptr(true),
// 	DeceasedBoolean: Ptr(false),
//
//  	Address: []fhir.Address{
// 		{
// 			City:    Ptr("Vitória de Santo Antão"),
// 			State:   Ptr("Pernambuco"),
// 			Country: Ptr("Brasil"),
// 			Type:    Ptr(fhir.AddressTypePhysical),
// 		},
// 	},
// 	Communication: []fhir.PatientCommunication{
// 		{
// 			Language: fhir.CodeableConcept{
// 				Text: Ptr("pt-br"),
// 				Coding: []fhir.Coding{
// 					{
// 						UserSelected: Ptr(true),
// 						Code:         Ptr("pt-br"),
// 						Version:      Ptr("1"),
// 					},
// 				},
// 			},
// 			Preferred: Ptr(true),
// 		},
// 	},
// }

// jsonBytes, err := json.MarshalIndent(p, "", "  ")
// if err != nil {
// 	fmt.Println("Error marshaling:", err)
// 	return
// }

// fmt.Println(string(jsonBytes))
