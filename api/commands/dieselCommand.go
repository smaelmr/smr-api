package commands

type DieselAdd struct {
	Date        string `json:"data"`
	DieselPrice string `json:"preco"`
	Quantity    string `json:"quantidade"`
	DieselTotal string `json:"dieselTotal"`
	SupplierId  string `json:"fornecedorId"`
	ArlaTotal   string `json:"ArlaTotal"`
	Plate       string `json:"placa"`
	Km          string `json:"km"`
}
