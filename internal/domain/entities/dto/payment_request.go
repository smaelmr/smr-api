package dto

import "time"

// PaymentRequest representa a requisição para realizar um pagamento
type PaymentRequest struct {
	ValorPago        float64   `json:"valorPago"`
	DataRealizacao   time.Time `json:"dataRealizacao"`
	FormaPagamentoId int64     `json:"formaPagamentoId"`
	LancarDiferenca  bool      `json:"lancarDiferenca"` // Se true, cria novo lançamento com a diferença
}
