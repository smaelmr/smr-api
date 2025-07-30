package entities

import "time"

type Freight struct {
	Id                      int64     `json:"id"`
	CarretaPlaca            string    `json:"carretaPlaca"`
	CavaloPlaca             string    `json:"cavaloPlaca"`
	ClienteId               int64     `json:"clienteId"`
	ClienteNome             string    `json:"clienteNome,omitempty"`
	OrigemId                int64     `json:"origemId"`
	OrigemNome              string    `json:"origemNome,omitempty"`
	DestinoFinalId          int64     `json:"destinoFinalId"`
	DestinoFinalNome        string    `json:"destinoFinalNome,omitempty"`
	FormaPagamentoId        int64     `json:"formaPagamentoId"`
	FormaPagamentoDescricao string    `json:"formaPagamentoDescricao,omitempty"`
	MotoristaId             int64     `json:"motoristaId"`
	MotoristaNome           string    `json:"motoristaNome,omitempty"`
	DataCarregamento        time.Time `json:"dataCarregamento"`
	DataEntrega             time.Time `json:"dataEntrega"`
	NumeroDocumento         string    `json:"numeroDocumento"`
	ValorAgenciamento       int64     `json:"valorAgenciamento"`
	ValorFrete              int64     `json:"valorFrete"`
	ValorPedagio            int64     `json:"valorPedagio"`
	Observacoes             string    `json:"observacoes"`
}
