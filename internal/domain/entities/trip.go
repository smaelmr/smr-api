package entities

import "time"

type Trip struct {
	Id                int64     `json:"id"`
	VeiculoId         int64     `json:"veiculoId"`
	ClienteId         int64     `json:"clienteId"`
	OrigemId          int64     `json:"origemId"`
	DestinoId         int64     `json:"destinoId"`
	MotoristaId       int64     `json:"motoristaId"`
	DataColeta        time.Time `json:"dataColeta"`
	DataEntrega       time.Time `json:"dataEntrega"`
	NumeroDocumento   string    `json:"numeroDocumento"`
	ValorAgenciamento float64   `json:"valorAgenciamento"`
	ValorFrete        float64   `json:"valorFrete"`
	ValorPedagio      float64   `json:"valorPedagio"`
	Observacao        string    `json:"observacao"`
}
