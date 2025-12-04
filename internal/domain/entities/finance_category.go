package entities

type Category struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"` // receita ou despesa
}
