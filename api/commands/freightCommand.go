package commands

type FreightAdd struct {
	CarretaPlaca      string `json:"carretaPlaca"`
	CavaloPlaca       string `json:"cavaloPlaca"`
	ClienteId         string `json:"clienteId"`
	OrigemId          string `json:"origemId"`
	DestinoFinalId    string `json:"destinoFinalId"`
	FormaPagamentoId  string `json:"formaPagamentoId"`
	MotoristaId       string `json:"motoristaId"`
	DataCarregamento  string `json:"dataCarregamento"`
	DataEntrega       string `json:"dataEntrega"`
	NumeroDocumento   string `json:"numeroDocumento"`
	ValorAgenciamento string `json:"valorAgenciamento"`
	ValorFrete        string `json:"valorFrete"`
	ValorPedagio      string `json:"valorPedagio"`
	Observacoes       string `json:"observacoes"`
}

type FreightUpdate struct {
	Id                string `json:"id"`
	CarretaPlaca      string `json:"carretaPlaca"`
	CavaloPlaca       string `json:"cavaloPlaca"`
	ClienteId         string `json:"clienteId"`
	OrigemId          int64  `json:"origemId"`
	DestinoFinalId    string `json:"destinoFinalId"`
	FormaPagamentoId  string `json:"formaPagamentoId"`
	MotoristaId       string `json:"motoristaId"`
	DataCarregamento  string `json:"dataCarregamento"`
	DataEntrega       string `json:"dataEntrega"`
	NumeroDocumento   string `json:"numeroDocumento"`
	ValorAgenciamento string `json:"valorAgenciamento"`
	ValorFrete        string `json:"valorFrete"`
	ValorPedagio      string `json:"valorPedagio"`
	Observacoes       string `json:"observacoes"`
}
