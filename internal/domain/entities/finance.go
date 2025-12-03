package entities

import "time"

type Finance struct {
	Id              int64     `json:"id"`
	PessoaId        int64     `json:"pessoaId"`
	CategoriaId     int64     `json:"categoriaId"`
	OrigemId        int64     `json:"lancamentoId"`   // ID do lançamento pai, pode ser manutenção, abastecimento ou frete. Origem manual será null
	Origem          string    `json:"lancamentoTipo"` // Descrição do tipo de lançamento: manutenção, abastecimento, frete ou manual
	Valor           float64   `json:"valor"`
	NumeroParcela   int32     `json:"numeroParcela"`
	NumeroDocumento string    `json:"numeroDocumento"`
	DataLancamento  time.Time `json:"dataLancamento"`
	DataVencimento  time.Time `json:"dataVencimento"`
	DataRealizacao  time.Time `json:"dataRealizacao"`
	Observacao      string    `json:"observacao"`
	Realizado       bool      `json:"realizado"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

/*type Reader interface {
	Get(id int) (*Finance, error)
	GetAll() (*[]Finance, error)
}

// Writer book writer
type Writer interface {
	Create(e *Finance) (Finance, error)
	Update(e *Finance) (Finance, error)
}

// Repository interface
type FinanceRepository interface {
	Reader
	Writer
}

// Service interface
type FinanceService interface {
	Get(id int) (*Finance, error)
	GetAll() (*[]Finance, error)
	Create(finance Finance) (int, error)
}*/
