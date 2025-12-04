package domain

type Endereco struct {
	CEP             string  `json:"cep"`
	Estado          string  `json:"estado"`
	Cidade          string  `json:"cidade"`
	Bairro          string  `json:"bairro"`
	Logradouro      string  `json:"logradouro"`
	Numero          string  `json:"numero"`
	Complemento     *string `json:"complemento"`
	PontoReferencia *string `json:"ponto_referencia"`
}
