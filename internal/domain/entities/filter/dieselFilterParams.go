package filter

import (
	"time"
)

type DieselFilterParams struct {
	FornecedorId *string `json:"fornecedorId,omitempty"`
	Placa        *string `json:"placa,omitempty"`
	DataInicial  *string `json:"dataInicial,omitempty"`
	DataFinal    *string `json:"dataFinal,omitempty"`
}

func (p *DieselFilterParams) ToFilter() (*DieselFilter, error) {
	filter := &DieselFilter{}

	if p.FornecedorId != nil && *p.FornecedorId != "" {
		id, err := ParseStringToInt64(*p.FornecedorId)
		if err != nil {
			return nil, err
		}
		filter.FornecedorId = &id
	}

	if p.Placa != nil && *p.Placa != "" {
		filter.Placa = p.Placa
	}

	if p.DataInicial != nil && *p.DataInicial != "" {
		data, err := time.Parse("2006-01-02", *p.DataInicial)
		if err != nil {
			return nil, err
		}
		filter.DataInicial = &data
	}

	if p.DataFinal != nil && *p.DataFinal != "" {
		data, err := time.Parse("2006-01-02", *p.DataFinal)
		if err != nil {
			return nil, err
		}
		filter.DataFinal = &data
	}

	// Validação adicional: DataFinal deve ser maior ou igual a DataInicial
	if filter.DataInicial != nil && filter.DataFinal != nil {
		if filter.DataFinal.Before(*filter.DataInicial) {
			return nil, ErrInvalidDateRange
		}
	}

	return filter, nil
}

func NewDieselFilterParams(fornecedorId, placa, dataInicial, dataFinal *string) *DieselFilterParams {
	return &DieselFilterParams{
		FornecedorId: fornecedorId,
		Placa:        placa,
		DataInicial:  dataInicial,
		DataFinal:    dataFinal,
	}
}
