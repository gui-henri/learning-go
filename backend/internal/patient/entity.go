package patient

import (
	"encoding/json"
	"time"

	"github.com/gui-henri/learning-go/pkg/fhir"
	"github.com/gui-henri/learning-go/pkg/util"
)

type paciente struct {
	InternalID   int64           `db:"internal_id" json:"internal_id"` //criado no banco
	ID           string          `db:"id" json:"id"`                   //criado no banco
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

func NewPaciente(p fhir.Patient) (paciente, error) {
	j, err := p.MarshalJSON()
	if err != nil {
		return paciente{}, err
	}

	t, err := time.Parse("", *p.BirthDate)
	if err != nil {
		return paciente{}, err
	}

	// Campos omitidos serão criados pelo banco
	pt := paciente{
		LastUpdated:  util.Ptr(time.Now()),
		Active:       true,
		Gender:       util.Ptr(p.Gender.Code()),
		BirthDate:    &t,
		FullName:     util.Ptr("a"),  // TODO: adicionar pré-processamento correto de nome
		CPF:          util.Ptr("aa"), // TODO: adicionar pré-processamento correto de CPF
		ResourceJSON: j,
	}

	return pt, nil
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
