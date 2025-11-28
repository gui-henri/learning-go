package avaliation

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	transport_encoding "github.com/gui-henri/learning-go/pkg/encoding"
	"github.com/jackc/pgx/v5"
)

func NewHttpTransportLayer(db *pgx.Conn, mux *http.ServeMux) *http.ServeMux {
	avaliationRepository := NewAvaliationRepository(db)
	avaliationService := NewAvaliationService(*avaliationRepository)

	saveAvaliation := httptransport.NewServer(
		MakeSaveFormEndpoint(*avaliationService),
		decodeSaveFormRequest,
		encodeResponse,
		httptransport.ServerErrorEncoder(transport_encoding.EncodeError),
	)

	mux.HandleFunc("POST /Avaliation", saveAvaliation.ServeHTTP)

	return mux
}
