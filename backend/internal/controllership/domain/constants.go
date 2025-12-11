package domain

const (
	TaxaTotal               = 0.02 + 0.0065 + 0.03 // 5,65%
	FatorLiquido            = 1 - TaxaTotal
	MultiplicadorCustoTotal = 1.1
	TaxaImpostos            = 0.34
)

var grupoProcedimentos = []string{
	"DIARIAS",
	"TAXAS",
	"SERV ENFERMAGEM",
	"SERVICOS DE FISIOTERAPIA",
	"SERVICOS DE FONOAUDIOLOGIA",
	"SERVICOS DE NUTRICAO",
	"SERVICOS DE TERAPIA OCUPACIONAL",
	"VISITAS E CONSULTAS MEDICAS - HC",
	"MEDICAMENTOS",
	"MEDICAMENTOS RESTRITO HOSP.",
	"ALIMENTACOES E DIETAS",
	"MATERIAIS",
	"CURATIVOS ESPECIAIS",
}
