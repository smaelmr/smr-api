package entities

type Finance struct {
	Id             int    `json:"id"`
	DueDate        string `json:"dueDate"`
	LaunchDate     string `json:"launchDate"`
	Amount         string `json:"amount"`
	Description    string `json:"description"`
	AccountId      string `json:"accountId"`
	PaymentMethod  string `json:"paymentMethod"`
	DocumentNumber string `json:"documentNumber"`
	Notes          string `json:"notes"`
	CategoryId     int    `json:"categotyId"`
	CostCenterId   int    `json:"costCenterId"`
	TagsId         int    `json:"tagsId"`
	Installments   int    `json:"instalments"`
	Repetetion     string `json:"repetetion"`
	SupplierId     int    `json:"supplierId"`
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
