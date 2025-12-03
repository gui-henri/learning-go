package avaliation

import (
	"fmt"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	transport_encoding "github.com/gui-henri/learning-go/pkg/encoding"
	"github.com/jackc/pgx/v5"
)

func NewHttpTransportLayer(db *pgx.Conn, mux *http.ServeMux, gotenbergUrl string, templateInternePath string) (*http.ServeMux, error) {
	avaliationRepository := NewAvaliationRepository(db)
	avaliationService, err := NewAvaliationService(*avaliationRepository, gotenbergUrl, templateInternePath)

	if err != nil {
		fmt.Println(err)
		fmt.Println("[ERROR] Não foi possível ler o template para criação do AvaliationService. Avaliação fora do ar.")
		return nil, fmt.Errorf("[ERROR] Não foi possível ler o template para criação do AvaliationService")
	}

	saveAvaliation := httptransport.NewServer(
		MakeSaveFormEndpoint(*avaliationService),
		decodeSaveFormRequest,
		encodeResponse,
		httptransport.ServerErrorEncoder(transport_encoding.EncodeError),
	)

	exportAvaliation := httptransport.NewServer(
		MakeExportAvaliationEndpoint(*avaliationService),
		decodeExportAvaliationRequest,
		encodeExportAvaliationResponse,
		httptransport.ServerErrorEncoder(transport_encoding.EncodeError),
	)

	mux.HandleFunc("POST /Avaliation", saveAvaliation.ServeHTTP)
	mux.HandleFunc("GET /Avaliation/{id}/export", exportAvaliation.ServeHTTP)

	return mux, nil
}
