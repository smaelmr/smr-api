package entities

import "time"

type Diesel struct {
	Id             int64     `json:"id"`
	Data           time.Time `json:"data"`
	Quantidade     int       `json:"quantidade"`
	DieselTotal    int64     `json:"dieselTotal"`
	FornecedorId   int64     `json:"fornecedorId"`
	FornecedorName string    `json:"fornecedorNome,omitempty"`
	ArlaTotal      int64     `json:"arlaTotal"`
	Placa          string    `json:"placa"`
	Km             int64     `json:"km"`
}

func (d *Diesel) PrecoPorLitro() int64 {
	if d.Quantidade == 0 {
		return 0
	}
	return d.DieselTotal / int64(d.Quantidade)
}
