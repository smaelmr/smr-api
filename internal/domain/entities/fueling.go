package entities

import (
	"time"
)

type Fueling struct {
	Id              int64     `json:"id"`
	VeiculoId       int64     `json:"veiculoId"`
	PostoId         int64     `json:"postoId"`
	Data            time.Time `json:"dataAbastecimento"`
	TipoCombustivel string    `json:"tipoCombustivel"`
	Litros          float64   `json:"litros"`
	ValorUnitario   float64   `json:"valorUnitario"`
	ValorTotal      float64   `json:"valorTotal"`
	Km              int64     `json:"km"`
	NumeroDocumento string    `json:"numeroDocumento"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
