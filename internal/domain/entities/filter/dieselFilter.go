package filter

import (
	"time"
)

type FuelingFilter struct {
	FornecedorId *int64     `json:"fornecedorId,omitempty"`
	Placa        *string    `json:"placa,omitempty"`
	DataInicial  *time.Time `json:"dataInicial,omitempty"`
	DataFinal    *time.Time `json:"dataFinal,omitempty"`
}
