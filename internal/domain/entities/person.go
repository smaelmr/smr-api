package entities

type Person struct {
	Id      int    `json:"id"`
	Name    string `json:"nome"`
	CpfCnpj string `json:"cpfCnpj"` // CPF ou CNPJ
	Type    string `json:"tipo"`    // Tipo pode ser "cliente" ou "fornecedor"
}
