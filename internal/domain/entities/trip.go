package entities

import "time"

type Trip struct {
	Id                int64     `json:"id"`
	CavaloPlaca       string    `json:"cavaloPlaca"`
	ClienteId         int64     `json:"clienteId"`
	OrigemId          int64     `json:"origemId"`
	DestinoId         int64     `json:"destinoId"`
	MotoristaId       int64     `json:"motoristaId"`
	DataCarregamento  time.Time `json:"dataCarregamento"`
	DataEntrega       time.Time `json:"dataEntrega"`
	NumeroDocumento   string    `json:"numeroDocumento"`
	ValorAgenciamento int64     `json:"valorAgenciamento"`
	ValorFrete        int64     `json:"valorFrete"`
	ValorPedagio      int64     `json:"valorPedagio"`
	Observacoes       string    `json:"observacoes"`
}
