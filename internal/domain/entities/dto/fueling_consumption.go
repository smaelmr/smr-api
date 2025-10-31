package dto

// FuelingConsumption representa o consumo médio de combustível por veículo
type FuelingConsumption struct {
	VeiculoId         int64   `json:"veiculoId"`
	Placa             string  `json:"placa"`
	TotalLitros       float64 `json:"totalLitros"`
	TotalKm           int64   `json:"totalKm"`
	MediaConsumo      float64 `json:"mediaConsumo"` // km/l
	QtdAbastecimentos int     `json:"qtdAbastecimentos"`
}
