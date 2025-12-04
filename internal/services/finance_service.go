package services

import (
	"errors"

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

	return s.RepoManager.Finance().Add(bill)
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

/*func (s *FinanceService) Filter(fornecedorId *string, placa *string, dataInicial *string, dataFinal *string) ([]entities.Finance, error) {

	filterParams := filter.NewFinanceFilterParams(fornecedorId, placa, dataInicial, dataFinal)

	dieselFilter, err := filterParams.ToFilter()
	if err != nil {
		return nil, err
	}

	return s.RepoManager.Finance().Filter(*dieselFilter)
}*/
