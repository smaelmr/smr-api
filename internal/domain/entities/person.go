package entities

type Person struct {
	Id           int64  `json:"id"`
	PessoaId     int64  `json:"pessoaId"`
	Name         string `json:"name"`
	Document     string `json:"document"` // CPF ou CNPJ
	Contact      string `json:"contact"`
	PhoneNumber  string `json:"phooneNumber"`
	Cep          string `json:"cep"`
	City         string `json:"city"`
	State        string `json:"state"`
	Street       string `json:"street"`
	StreetNumber string `json:"number"`
	Neighborhood string `json:"neighborhood"`
}
