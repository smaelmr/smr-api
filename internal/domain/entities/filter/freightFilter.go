package filter

import "time"

type TripFilter struct {
	ClienteId      *int64     `json:"clienteId,omitempty"`
	MotoristaId    *int64     `json:"motoristaId,omitempty"`
	DataInicial    *time.Time `json:"dataInicial,omitempty"`
	DataFinal      *time.Time `json:"dataFinal,omitempty"`
	CavaloPlaca    *string    `json:"cavaloPlaca,omitempty"`
	CarretaPlaca   *string    `json:"carretaPlaca,omitempty"`
	OrigemId       *int64     `json:"origemId,omitempty"`
	DestinoFinalId *int64     `json:"destinoFinalId,omitempty"`
}
