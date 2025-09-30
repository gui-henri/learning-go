package patient

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	transport_encoding "github.com/gui-henri/learning-go/pkg/encoding"
	"github.com/jackc/pgx/v5"
)

func NewHttpTransportLayer(db *pgx.Conn, mux *http.ServeMux) *http.ServeMux {
	patientRepository := NewPatientRepository(db)
	patientService := NewPatientService(*patientRepository)

	insertPatient := httptransport.NewServer(
		makeInsertPatientEndpoint(patientService),
		transport_encoding.EncodeRequest[InsertPatientRequest],
		transport_encoding.EncodeResponse,
		httptransport.ServerErrorEncoder(transport_encoding.EncodeError),
	)

	listPatient := httptransport.NewServer(
		makeListPatientEndpoint(patientService),
		decodeListPatientRequest,
		transport_encoding.EncodeResponse,
		httptransport.ServerErrorEncoder(transport_encoding.EncodeError),
	)

	mux.HandleFunc("POST /Patient", insertPatient.ServeHTTP)
	mux.HandleFunc("GET /Patient", listPatient.ServeHTTP)

	return mux
}
