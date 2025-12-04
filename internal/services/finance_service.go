package services

import (
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

func (s *FinanceService) GetAll() ([]entities.Finance, error) {
	records, err := s.RepoManager.Finance().GetAll()
	if err != nil {
		return nil, err
	}

	var dieselList []entities.Finance
	dieselList = append(dieselList, records...)

	return dieselList, nil
}

func (s *FinanceService) GetReceipts() ([]entities.Finance, error) {
	records, err := s.RepoManager.Finance().GetAll()
	if err != nil {
		return nil, err
	}

	var dieselList []entities.Finance
	dieselList = append(dieselList, records...)

	return dieselList, nil
}

func (s *FinanceService) GetPayments() ([]entities.Finance, error) {
	records, err := s.RepoManager.Finance().GetAll()
	if err != nil {
		return nil, err
	}

	var dieselList []entities.Finance
	dieselList = append(dieselList, records...)

	return dieselList, nil
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
