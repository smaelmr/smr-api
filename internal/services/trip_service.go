package services

import (
	"github.com/smaelmr/finance-api/internal/domain/contract/repository"
	"github.com/smaelmr/finance-api/internal/domain/entities"
	"github.com/smaelmr/finance-api/internal/domain/entities/filter"
)

type TripService struct {
	RepoManager repository.RepoManager
}

func NewTripService(repoManager repository.RepoManager) *TripService {
	return &TripService{
		RepoManager: repoManager,
	}
}

func (s *TripService) Add(tripAdd *entities.Trip) error {

	/*dataCarregamento, err := ParseStringToTime(tripAdd.DataCarregamento, "02/01/2006")
	if err != nil {
		return err
	}*/

	return s.RepoManager.Trip().Add(*tripAdd)
}

func (s *TripService) GetAll() ([]entities.Trip, error) {
	records, err := s.RepoManager.Trip().GetAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (s *TripService) Update(tripUpdate *entities.Trip) error {
	return s.RepoManager.Trip().Update(*tripUpdate)
}

func (s *TripService) Filter(clienteId, motoristaId, dataInicial, dataFinal, cavaloPlaca *string) ([]entities.Trip, error) {
	filterParams := filter.NewTripFilterParams(clienteId, motoristaId, dataInicial, dataFinal, cavaloPlaca)

	tripFilter, err := filterParams.ToFilter()
	if err != nil {
		return nil, err
	}

	records, err := s.RepoManager.Trip().Filter(*tripFilter)
	if err != nil {
		return nil, err
	}

	return records, nil
}
