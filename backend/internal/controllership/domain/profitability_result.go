package domain

import "strconv"

type ProfitabilityResult struct {
	CategoryResults     []CategoryResult
	TotalReceitaBruta   float64
	TotalReceitaLiquida float64
	TotalCustosMV       float64
	TotalCustosFixos    float64
	TotalMargemBruta    float64
	TotalImpostos       float64
	TotalMargemLiquida  float64
	Rentabilidade       float64
}

func RunProfitabilityAnalysis(cabecalho Cabecalho, orcamento []OrcamentoItem) ProfitabilityResult {
	result := ProfitabilityResult{}

	pacote := ChecarPacientePacote(orcamento)
	var itensOrcamento []OrcamentoItem
	if pacote {
		itensOrcamento = ProccessPacientePacote(orcamento)
	} else {
		itensOrcamento = ProccessPacienteNaoPacote(orcamento)
	}

	quantidadeDias, _ := strconv.Atoi(cabecalho.Internacao.QuantidadeDias)
	cfixoTotal := CustoFixoTotal(MapCustos[cabecalho.Internacao.Complexidade], quantidadeDias) * MultiplicadorCustoTotal

	var totalMargemBrutaPositiva float64

	for _, categoria := range grupoProcedimentos {
		categoryResult := CategoryResult{
			GrupoProcedimento: categoria,
		}

		rectBruta := ReceitaBrutaCategoria(itensOrcamento, categoria)
		categoryResult.ReceitaBruta = rectBruta
		result.TotalReceitaBruta += rectBruta

		categoryResult.ReceitaLiquida = rectBruta * FatorLiquido
		result.TotalReceitaLiquida += categoryResult.ReceitaLiquida

		custos := CustosMV(itensOrcamento, categoria)
		categoryResult.CustosMV = custos
		result.TotalCustosMV += custos

		if categoria == "TAXAS" {
			categoryResult.CustosFixos = cfixoTotal
			result.TotalCustosFixos += cfixoTotal
		}

		categoryResult.MargemBruta = categoryResult.ReceitaLiquida - (categoryResult.CustosMV + categoryResult.CustosFixos)
		if categoryResult.MargemBruta > 0 {
			totalMargemBrutaPositiva += categoryResult.MargemBruta
		}
		result.CategoryResults = append(result.CategoryResults, categoryResult)
	}

	totalImpostos := 0.0
	if totalMargemBrutaPositiva > 0 {
		totalMargemBruta := result.TotalReceitaLiquida - (result.TotalCustosMV + result.TotalCustosFixos)
		if totalMargemBruta > 0 {
			totalImpostos = totalMargemBruta * TaxaImpostos
		}
		taxaRateio := totalImpostos / totalMargemBrutaPositiva

		for i := range result.CategoryResults {
			if result.CategoryResults[i].MargemBruta > 0 {
				result.CategoryResults[i].Impostos = result.CategoryResults[i].MargemBruta * taxaRateio
			}
			result.CategoryResults[i].MargemLiquida = result.CategoryResults[i].MargemBruta - result.CategoryResults[i].Impostos
		}
	}

	result.TotalMargemBruta = result.TotalReceitaLiquida - (result.TotalCustosMV + result.TotalCustosFixos)
	result.TotalImpostos = totalImpostos
	result.TotalMargemLiquida = result.TotalMargemBruta - result.TotalImpostos

	if result.TotalReceitaLiquida != 0 {
		result.Rentabilidade = (result.TotalMargemLiquida / result.TotalReceitaLiquida) * 100
	} else {
		result.Rentabilidade = 0.0
	}

	return result
}
