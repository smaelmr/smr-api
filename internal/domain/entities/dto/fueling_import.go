package dto

import (
	"strconv"
	"strings"
	"time"
)

type FuelingImport struct {
	DataTransacao string `json:"dataTransacao"`
	NumeroCupom   string `json:"numeroCupom"`
	CnpjPosto     string `json:"cnpjPosto"`
	Placa         string `json:"placa"`
	CpfMotorista  string `json:"cpfMotorista"`
	Hodometro     string `json:"hodometro"`
	Produto       string `json:"produto"`
	ValorUnitario string `json:"valorUnitario"`
	Quantidade    string `json:"quantidade"`
	ValorTotal    string `json:"valorTotal"`
}

func (f *FuelingImport) DataTransacaoTime() time.Time {
	// Ajuste o layout conforme o formato real da data
	layout := "2006-01-02T15:04:05" // exemplo: 2025-10-25T14:30:00
	dt, _ := time.Parse(layout, f.DataTransacao)

	return dt
}

func (f *FuelingImport) HodometroInt64() int64 {
	km, _ := strconv.ParseInt(f.Hodometro, 10, 64)

	return km
}

func (f *FuelingImport) ValorTotalFloat64() float64 {
	// Remove os pontos de milhar
	withoutDots := strings.ReplaceAll(f.ValorTotal, ".", "")
	// Substitui a vírgula decimal por ponto
	withDot := strings.ReplaceAll(withoutDots, ",", ".")

	vl, _ := strconv.ParseFloat(withDot, 64)
	return vl
}

func (f *FuelingImport) QuantidadeFloat64() float64 {
	// Remove os pontos de milhar
	withoutDots := strings.ReplaceAll(f.Quantidade, ".", "")
	// Substitui a vírgula decimal por ponto
	withDot := strings.ReplaceAll(withoutDots, ",", ".")

	val, err := strconv.ParseFloat(withDot, 64)
	if err != nil {
		return 0
	}
	// Arredonda para 3 casas decimais
	return float64(int(val*1000)) / 1000
}

func (f *FuelingImport) ProdutoMappedRussi() (tipoProduto string) {

	switch strings.ToUpper(strings.TrimSpace(f.Produto)) {
	case "DIESEL S-10 COMUM":
		tipoProduto = "Diesel_S10"
	case "DIESEL S-500 COMUM":
		tipoProduto = "Diesel_S500"
	case "ARLA GRANEL":
		tipoProduto = "Arla"
	default:
		tipoProduto = f.Produto
	}

	return tipoProduto

}

func (f *FuelingImport) ProdutoMappedGraal() (tipoProduto string) {

	switch strings.ToUpper(strings.TrimSpace(f.Produto)) {
	case "OLEO DIESEL BS10":
		tipoProduto = "Diesel_S10"
	case "OLEO DIESEL B S 500":
		tipoProduto = "Diesel_S500"
	case "ARLA32 GRANELL":
		tipoProduto = "Arla"
	default:
		tipoProduto = f.Produto
	}

	return tipoProduto

}
