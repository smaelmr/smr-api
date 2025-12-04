package entities

import "time"

type Finance struct {
	Id               int64      `json:"id"`
	PessoaId         int64      `json:"pessoaId"`
	CategoriaId      int64      `json:"categoriaId"`
	FormaPagamentoId *int64     `json:"formaPagamentoId"` // ID da forma de pagamento (null se não pago)
	OrigemId         *int64     `json:"OrigemId"`         // ID do lançamento pai, pode ser manutenção, abastecimento ou frete. Origem manual será null
	Origem           string     `json:"origem"`           // Descrição do tipo de lançamento: manutenção, abastecimento, frete ou manual
	Valor            float64    `json:"valor"`
	ValorPago        *float64   `json:"valorPago"` // Valor efetivamente pago (null se não pago)
	ValorParcela     float64    `json:"valorParcela"`
	NumeroParcela    int32      `json:"numeroParcela"`
	TotalParcelas    int32      `json:"totalParcelas"` // Número total de parcelas (usado na criação para gerar múltiplos registros)
	NumeroDocumento  string     `json:"numeroDocumento"`
	DataCompetencia  time.Time  `json:"dataCompetencia"`
	DataVencimento   time.Time  `json:"dataVencimento"`
	DataRealizacao   *time.Time `json:"dataRealizacao"`
	Observacao       string     `json:"observacao"`
	Realizado        bool       `json:"realizado"`
	CreatedAt        time.Time  `json:"createdAt"`
	UpdatedAt        time.Time  `json:"updatedAt"`
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
