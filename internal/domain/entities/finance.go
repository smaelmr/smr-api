package entities

import "time"

type Finance struct {
	Id              int64     `json:"id"`
	PessoaId        int64     `json:"pessoaId"`
	Valor           float64   `json:"valor"`
	NumeroDocumento string    `json:"numeroDocumento"`
	DataLancamento  time.Time `json:"dataLancamento"`
	DataVencimento  time.Time `json:"dataVencimento"`
	DataRealizacao  time.Time `json:"dataRealizacao"`
	Origem          string    `json:"origem"`
	Observacao      string    `json:"observacao"`
	Recebido        bool      `json:"recebido"`
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
