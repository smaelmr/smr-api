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

func (f *FuelingImport) DataTransacaoTime() (time.Time, error) {
	// Ajuste o layout conforme o formato real da data
	layout := "2006-01-02T15:04:05" // exemplo: 2025-10-25T14:30:00
	return time.Parse(layout, f.DataTransacao)
}

func (f *FuelingImport) HodometroInt64() (int64, error) {
	return strconv.ParseInt(f.Hodometro, 10, 64)
}

func (f *FuelingImport) ValorTotalFloat64() (float64, error) {
	return strconv.ParseFloat(strings.ReplaceAll(f.ValorTotal, ",", "."), 64)
}

func (f *FuelingImport) QuantidadeFloat64() (float64, error) {
	val, err := strconv.ParseFloat(strings.ReplaceAll(f.Quantidade, ",", "."), 64)
	if err != nil {
		return 0, err
	}
	// Arredonda para 3 casas decimais
	return float64(int(val*1000)) / 1000, nil
}

func (f *FuelingImport) ProdutoMapped() string {
	switch strings.ToUpper(strings.TrimSpace(f.Produto)) {
	case "DIESEL S-10 COMUM":
		return "Diesel_S10"
	case "DIESEL S-500 COMUM":
		return "Diesel_S500"
	case "ARLA GRANEL":
		return "Arla"
	default:
		return f.Produto
	}
}
