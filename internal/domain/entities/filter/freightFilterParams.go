package filter

import "time"

type TripFilterParams struct {
	ClienteId      *string `json:"clienteId,omitempty"`
	MotoristaId    *string `json:"motoristaId,omitempty"`
	DataInicial    *string `json:"dataInicial,omitempty"`
	DataFinal      *string `json:"dataFinal,omitempty"`
	CavaloPlaca    *string `json:"cavaloPlaca,omitempty"`
	CarretaPlaca   *string `json:"carretaPlaca,omitempty"`
	OrigemId       *string `json:"origemId,omitempty"`
	DestinoFinalId *string `json:"destinoFinalId,omitempty"`
}

func (p *TripFilterParams) ToFilter() (*TripFilter, error) {
	filter := &TripFilter{}

	if p.ClienteId != nil && *p.ClienteId != "" {
		id, err := ParseStringToInt64(*p.ClienteId)
		if err != nil {
			return nil, err
		}
		filter.ClienteId = &id
	}

	if p.MotoristaId != nil && *p.MotoristaId != "" {
		id, err := ParseStringToInt64(*p.MotoristaId)
		if err != nil {
			return nil, err
		}
		filter.MotoristaId = &id
	}

	if p.OrigemId != nil && *p.OrigemId != "" {
		id, err := ParseStringToInt64(*p.OrigemId)
		if err != nil {
			return nil, err
		}
		filter.OrigemId = &id
	}

	if p.DestinoFinalId != nil && *p.DestinoFinalId != "" {
		id, err := ParseStringToInt64(*p.DestinoFinalId)
		if err != nil {
			return nil, err
		}
		filter.DestinoFinalId = &id
	}

	if p.CavaloPlaca != nil && *p.CavaloPlaca != "" {
		filter.CavaloPlaca = p.CavaloPlaca
	}

	if p.CarretaPlaca != nil && *p.CarretaPlaca != "" {
		filter.CarretaPlaca = p.CarretaPlaca
	}

	if p.DataInicial != nil && *p.DataInicial != "" {
		// Converte de DD/MM/YYYY para time.Time
		t, err := time.Parse("2006-01-02", *p.DataInicial)
		if err != nil {
			return nil, err
		}
		// Define hora como início do dia (00:00:00)
		t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
		filter.DataInicial = &t
	}

	if p.DataFinal != nil && *p.DataFinal != "" {
		// Converte de DD/MM/YYYY para time.Time
		t, err := time.Parse("2006-01-02", *p.DataFinal)
		if err != nil {
			return nil, err
		}
		// Define hora como fim do dia (23:59:59)
		t = time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, time.Local)
		filter.DataFinal = &t
	}

	// Validação adicional: DataFinal deve ser maior ou igual a DataInicial
	if filter.DataInicial != nil && filter.DataFinal != nil {
		if filter.DataFinal.Before(*filter.DataInicial) {
			return nil, ErrInvalidDateRange
		}
	}

	return filter, nil
}

func NewTripFilterParams(clienteId, motoristaId, dataInicial, dataFinal, cavaloPlaca *string) *TripFilterParams {
	return &TripFilterParams{
		ClienteId:   clienteId,
		MotoristaId: motoristaId,
		DataInicial: dataInicial,
		DataFinal:   dataFinal,
		CavaloPlaca: cavaloPlaca,
	}
}
