package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/smaelmr/finance-api/internal/domain/contract/repository"
	"github.com/smaelmr/finance-api/internal/domain/entities"
)

type FinanceService struct {
	RepoManager repository.RepoManager
}

func NewFinanceService(repoManager repository.RepoManager) *FinanceService {
	return &FinanceService{
		RepoManager: repoManager,
	}
}

func (s *FinanceService) Add(bill entities.Finance) error {
	// Se totalParcelas não foi informado ou é 1, criar apenas um registro
	if bill.TotalParcelas <= 1 {
		bill.NumeroParcela = 1
		bill.TotalParcelas = 1
		return s.RepoManager.Finance().Add(bill)
	}

	// Se há múltiplas parcelas, criar um registro para cada parcela
	totalParcelas := bill.TotalParcelas
	valorParcela := bill.Valor / float64(totalParcelas)
	valorOriginal := bill.Valor

	// Criar cada parcela
	for i := int32(1); i <= totalParcelas; i++ {
		parcela := bill
		parcela.NumeroParcela = i

		// Para a última parcela, ajustar o valor para compensar arredondamentos
		if i == totalParcelas {
			// Calcular o valor já lançado nas parcelas anteriores
			valorLancado := valorParcela * float64(i-1)
			parcela.ValorParcela = valorOriginal - valorLancado
		} else {
			parcela.ValorParcela = valorParcela
		}

		// Incrementar a data de vencimento baseado no número da parcela
		// Adiciona (i-1) meses à data de vencimento original
		parcela.DataVencimento = bill.DataVencimento.AddDate(0, int(i-1), 0)

		// Adicionar número da parcela no número do documento
		if bill.NumeroDocumento != "" {
			parcela.NumeroDocumento = fmt.Sprintf("%s-%d/%d", bill.NumeroDocumento, i, totalParcelas)
		}

		err := s.RepoManager.Finance().Add(parcela)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *FinanceService) GetAll(categoryType string, month int, year int) ([]entities.Finance, error) {
	if categoryType != "R" && categoryType != "D" {
		return nil, errors.New("type must be 'R' (receita) or 'D' (despesa)")
	}

	if month < 1 || month > 12 {
		return nil, errors.New("month must be between 1 and 12")
	}

	if year < 1900 {
		return nil, errors.New("year must be greater than 1900")
	}

	records, err := s.RepoManager.Finance().GetAll(categoryType, month, year)
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (s *FinanceService) GetReceipts(month int, year int) ([]entities.Finance, error) {
	if month < 1 || month > 12 {
		return nil, errors.New("month must be between 1 and 12")
	}

	if year < 1900 {
		return nil, errors.New("year must be greater than 1900")
	}

	records, err := s.RepoManager.Finance().GetAll("R", month, year)
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (s *FinanceService) GetPayments(month int, year int) ([]entities.Finance, error) {
	if month < 1 || month > 12 {
		return nil, errors.New("month must be between 1 and 12")
	}

	if year < 1900 {
		return nil, errors.New("year must be greater than 1900")
	}

	records, err := s.RepoManager.Finance().GetAll("D", month, year)
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (s *FinanceService) Update(dieselUpdate *entities.Finance) error {
	return s.RepoManager.Finance().Update(*dieselUpdate)
}

func (s *FinanceService) ProcessPayment(id int64, valorPago float64, dataRealizacao time.Time, formaPagamentoId int64, lancarDiferenca bool) error {
	// Buscar o lançamento original
	lancamento, err := s.RepoManager.Finance().Get(id)
	if err != nil {
		return err
	}

	if lancamento == nil {
		return errors.New("finance record not found")
	}

	// Verificar se já foi realizado
	if lancamento.Realizado {
		return errors.New("payment already processed")
	}

	// Validações
	if valorPago <= 0 {
		return errors.New("valorPago must be greater than zero")
	}

	if formaPagamentoId <= 0 {
		return errors.New("formaPagamentoId is required")
	}

	valorOriginal := lancamento.ValorParcela
	if valorOriginal == 0 {
		valorOriginal = lancamento.Valor
	}

	// Atualizar o lançamento original
	lancamento.ValorPago = &valorPago
	lancamento.DataRealizacao = &dataRealizacao
	lancamento.FormaPagamentoId = &formaPagamentoId
	lancamento.Realizado = true

	err = s.RepoManager.Finance().Update(*lancamento)
	if err != nil {
		return err
	}

	// Se solicitado, lançar a diferença
	if lancarDiferenca {
		diferenca := valorPago - valorOriginal

		// Se houver diferença (positiva ou negativa), criar novo lançamento
		if diferenca != 0 {
			novoLancamento := entities.Finance{
				PessoaId:         lancamento.PessoaId,
				CategoriaId:      lancamento.CategoriaId,
				OrigemId:         &id, // Referenciar o lançamento original
				Origem:           "ajuste_pagamento",
				Valor:            diferenca,
				ValorParcela:     diferenca,
				NumeroParcela:    1,
				TotalParcelas:    1,
				NumeroDocumento:  fmt.Sprintf("%s-AJUSTE", lancamento.NumeroDocumento),
				DataCompetencia:  dataRealizacao,
				DataVencimento:   dataRealizacao,
				DataRealizacao:   &dataRealizacao,
				FormaPagamentoId: &formaPagamentoId,
				Observacao:       fmt.Sprintf("Ajuste de pagamento - Ref: %s (Diferença: %.2f)", lancamento.NumeroDocumento, diferenca),
				Realizado:        true,
			}

			// Determinar se é acréscimo (receita) ou desconto (despesa)
			// Se pagou mais que o valor original, é uma despesa adicional
			// Se pagou menos, é um desconto (receita)
			if diferenca > 0 {
				// Pagou mais: criar como despesa (mesma categoria se for despesa)
				novoLancamento.Observacao += " - Acréscimo"
			} else {
				// Pagou menos: criar como desconto (receita)
				novoLancamento.Valor = -diferenca
				novoLancamento.ValorParcela = -diferenca
				novoLancamento.Observacao += " - Desconto"
			}

			err = s.RepoManager.Finance().Add(novoLancamento)
			if err != nil {
				return fmt.Errorf("payment processed but failed to create difference record: %v", err)
			}
		}
	}

	return nil
}

/*func (s *FinanceService) Filter(fornecedorId *string, placa *string, dataInicial *string, dataFinal *string) ([]entities.Finance, error) {

	filterParams := filter.NewFinanceFilterParams(fornecedorId, placa, dataInicial, dataFinal)

	dieselFilter, err := filterParams.ToFilter()
	if err != nil {
		return nil, err
	}

	return s.RepoManager.Finance().Filter(*dieselFilter)
}*/
