package domain

import "github.com/gui-henri/learning-go/pkg/util"

type OrcamentoItem struct {
	Descricao         string  `json:"descricao"`
	Quantidade        float64 `json:"quantidade"`
	Unidade           string  `json:"unidade"`
	ValorUnitario     string  `json:"valor_unitario"`
	ValorTotal        string  `json:"valor_total"`
	GrupoProcedimento string  `json:"grupo_procedimento"`
	CustoUnitario     string  `json:"custo_unitario"`
	CustoTotal        string  `json:"custo_total"`
	Pacote            string  `json:"pacote"`
}

// ChecarPacientePacote checks if the patient is a package patient.
func ChecarPacientePacote(itensOrcamento []OrcamentoItem) bool {
	for _, item := range itensOrcamento {
		if item.GrupoProcedimento == "DIARIAS" {
			return true
		}
	}
	return false
}

func ProccessPacientePacote(itensOrcamento []OrcamentoItem) []OrcamentoItem {
	for i := range itensOrcamento {
		if itensOrcamento[i].GrupoProcedimento == "DIARIAS" {
			itensOrcamento[i].Pacote = "NÃO"
		} else {
			itensOrcamento[i].Pacote = "SIM"
		}
	}
	return itensOrcamento
}

func ProccessPacienteNaoPacote(itensOrcamento []OrcamentoItem) []OrcamentoItem {
	for i := range itensOrcamento {
		itensOrcamento[i].Pacote = "NÃO"
	}
	return itensOrcamento
}

func CustosMV(itensOrcamento []OrcamentoItem, categoria string) float64 {
	total := 0.0
	for _, item := range itensOrcamento {
		if item.GrupoProcedimento == categoria {
			total += util.ParseValor(item.CustoTotal)
		}
	}
	return total
}

func ReceitaBrutaCategoria(itensOrcamento []OrcamentoItem, categoria string) float64 {
	total := 0.0
	for _, item := range itensOrcamento {
		if item.GrupoProcedimento == categoria && item.Pacote == "NÃO" {
			total += util.ParseValor(item.ValorTotal)
		}
	}
	return total
}
