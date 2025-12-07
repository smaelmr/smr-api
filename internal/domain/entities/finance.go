package entities

import (
	"encoding/json"
	"time"
)

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
	CreatedAt        time.Time  `json:"createdAt"`
	UpdatedAt        time.Time  `json:"updatedAt"`
}

// GetStatus retorna o status do lançamento baseado na data de realização e vencimento
func (f *Finance) GetStatus() string {
	// Se tem data de realização, está pago/recebido
	if f.DataRealizacao != nil {
		return "Pago/Recebido"
	}

	// Se não tem data de realização, verificar vencimento
	hoje := time.Now()

	// Normalizar as datas para comparação apenas de dia (sem hora)
	hojeNormalizado := time.Date(hoje.Year(), hoje.Month(), hoje.Day(), 0, 0, 0, 0, hoje.Location())
	vencimentoNormalizado := time.Date(f.DataVencimento.Year(), f.DataVencimento.Month(), f.DataVencimento.Day(), 0, 0, 0, 0, f.DataVencimento.Location())

	// Se já passou do vencimento
	if hojeNormalizado.After(vencimentoNormalizado) {
		return "Em Atraso"
	}

	// Caso contrário, está em aberto dentro do prazo
	return "Em Aberto"
}

// MarshalJSON customiza a serialização JSON para incluir o campo Status dinamicamente
func (f *Finance) MarshalJSON() ([]byte, error) {
	type Alias Finance
	return json.Marshal(&struct {
		*Alias
		Status string `json:"status"`
	}{
		Alias:  (*Alias)(f),
		Status: f.GetStatus(),
	})
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
