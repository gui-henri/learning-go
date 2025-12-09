package expansion

import (
	"context"
	"fmt"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gui-henri/learning-go/db"
	"github.com/gui-henri/learning-go/internal/expansion/requests"
	transport_encoding "github.com/gui-henri/learning-go/pkg/encoding"
)

func NewHttpTransportLayer(db db.IDB, mux *http.ServeMux, gotenbergUrl string, templateInternePath string) (*http.ServeMux, error) {
	avaliationRepository := NewAvaliationRepository(db)
	avaliationService, err := NewAvaliationService(avaliationRepository, gotenbergUrl, templateInternePath)

	if err != nil {
		fmt.Println(err)
		fmt.Println("[ERROR] Não foi possível ler o template para criação do AvaliationService. Avaliação fora do ar.")
		return nil, fmt.Errorf("[ERROR] Não foi possível ler o template para criação do AvaliationService")
	}

	saveAvaliation := httptransport.NewServer(
		MakeSaveAvaliationEndpoint(avaliationService),
		requests.DecodeSaveAvaliationRequest,
		requests.EncodeSaveAvaliationResponse,
		httptransport.ServerErrorEncoder(transport_encoding.EncodeError),
	)

	listAvaliation := httptransport.NewServer(
		MakeListAvaliationEndpoint(avaliationService),
		func(ctx context.Context, r *http.Request) (any, error) { return nil, nil },
		requests.EncodeListAvaliationResponse,
		httptransport.ServerErrorEncoder(transport_encoding.EncodeError),
	)

	exportAvaliation := httptransport.NewServer(
		MakeExportAvaliationEndpoint(avaliationService),
		requests.DecodeExportAvaliationRequest,
		requests.EncodeExportAvaliationResponse,
		httptransport.ServerErrorEncoder(transport_encoding.EncodeError),
	)

	mux.HandleFunc("POST /Avaliation", saveAvaliation.ServeHTTP)
	mux.HandleFunc("GET /Avaliation", listAvaliation.ServeHTTP)
	mux.HandleFunc("GET /Avaliation/{id}/export", exportAvaliation.ServeHTTP)

	return mux, nil
}
