package entities

type Category struct {
	Id          int64  `json:"id"`
	Description string `json:"description"`
	Type        string `json:"type"` // receita ou despesa
}
