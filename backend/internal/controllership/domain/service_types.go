package domain

var ServiceTypes = map[string]float64{
	"Internação Domiciliar 6hrs":  50,
	"Internação Domiciliar 12hrs": 100,
	"Internação Domiciliar 24hrs": 200,
	"Assistência Domiciliar":      35,
	"Terapias":                    9,
	"Curabem":                     28.8,
	"Oxigenoterapia":              11.6,
	"Fiqbem":                      3.10,
}

var MapCustos = map[string]string{
	"BAIXA COMPLEXIDADE": "Internação Domiciliar 6hrs",
	"MEDIA COMPLEXIDADE": "Internação Domiciliar 12hrs",
	"ALTA COMPLEXIDADE":  "Internação Domiciliar 24hrs",
}

func CustoFixoTotal(serviceType string, period int) float64 {
	servicoCusto := ServiceTypes[serviceType]
	custoTotal := servicoCusto * float64(period)
	return custoTotal
}
