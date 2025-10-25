package entities

import "time"

type Vehicle struct {
	Id        int64     `json:"id"`
	Placa     string    `json:"placa"`
	Marca     string    `json:"marca"`
	Modelo    string    `json:"modelo"`
	Ano       int       `json:"ano"`
	Tipo      string    `json:"tipo"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
