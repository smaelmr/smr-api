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
	ValorDiesel     float64   `json:"valorDiesel"`
	ValorArla       float64   `json:"valorArla"`
	ValorDiversos   float64   `json:"valorDiversos"` // Valor de diversos (usado apenas para somar no total do financeiro, não salvo no banco)
	Km              int64     `json:"km"`
	NumeroDocumento string    `json:"numeroDocumento"`
	Cheio           bool      `json:"cheio"` //informar se o abastecimento foi completo ou parcial
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
