package services

import (
	"github.com/smaelmr/finance-api/internal/domain/contract/repository"
	"github.com/smaelmr/finance-api/internal/domain/entities"
	"github.com/smaelmr/finance-api/internal/domain/entities/filter"
)

type FreightService struct {
	RepoManager repository.RepoManager
}

func NewFreightService(repoManager repository.RepoManager) *FreightService {
	return &FreightService{
		RepoManager: repoManager,
	}
}

func (s *FreightService) Add(freightAdd *entities.Freight) error {

	/*dataCarregamento, err := ParseStringToTime(freightAdd.DataCarregamento, "02/01/2006")
	if err != nil {
		return err
	}*/

	return s.RepoManager.Freight().Add(*freightAdd)
}

func (s *FreightService) GetAll() ([]entities.Freight, error) {
	records, err := s.RepoManager.Freight().GetAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (s *FreightService) Update(freightUpdate *entities.Freight) error {
	return s.RepoManager.Freight().Update(*freightUpdate)
}

func (s *FreightService) Filter(clienteId, motoristaId, dataInicial, dataFinal, cavaloPlaca *string) ([]entities.Freight, error) {
	filterParams := filter.NewFreightFilterParams(clienteId, motoristaId, dataInicial, dataFinal, cavaloPlaca)

	freightFilter, err := filterParams.ToFilter()
	if err != nil {
		return nil, err
	}

	records, err := s.RepoManager.Freight().Filter(*freightFilter)
	if err != nil {
		return nil, err
	}

	return records, nil
}
