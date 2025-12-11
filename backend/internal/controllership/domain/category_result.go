package domain

type CategoryResult struct {
	GrupoProcedimento string
	ReceitaBruta      float64
	ReceitaLiquida    float64
	CustosMV          float64
	CustosFixos       float64
	MargemBruta       float64
	Impostos          float64
	MargemLiquida     float64
}
