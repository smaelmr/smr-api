package commands

type FuelingAdd struct {
	Date         string `json:"data"`
	FuelingPrice string `json:"preco"`
	Quantity     string `json:"quantidade"`
	FuelingTotal string `json:"dieselTotal"`
	SupplierId   string `json:"fornecedorId"`
	ArlaTotal    string `json:"ArlaTotal"`
	Plate        string `json:"placa"`
	Km           string `json:"km"`
}
