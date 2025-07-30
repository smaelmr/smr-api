package services

import (
	"github.com/smaelmr/finance-api/internal/domain/contract/repository"
	"github.com/smaelmr/finance-api/internal/domain/entities"
	"github.com/smaelmr/finance-api/internal/domain/entities/filter"
)

type DieselService struct {
	RepoManager repository.RepoManager
}

func NewDieselService(repoManager repository.RepoManager) *DieselService {
	return &DieselService{
		RepoManager: repoManager,
	}
}

func (s *DieselService) Add(dieselAdd *entities.Diesel) error {

	/*date, err := ParseStringToTime(dieselAdd.Date, "02/01/2006")
	if err != nil {
		return err
	}*/

	return s.RepoManager.Diesel().Add(*dieselAdd)
}

func (s *DieselService) GetAll() ([]entities.Diesel, error) {
	records, err := s.RepoManager.Diesel().GetAll()
	if err != nil {
		return nil, err
	}

	var dieselList []entities.Diesel
	dieselList = append(dieselList, records...)

	return dieselList, nil
}

func (s *DieselService) Update(dieselUpdate *entities.Diesel) error {
	return s.RepoManager.Diesel().Update(*dieselUpdate)
}

func (s *DieselService) Filter(fornecedorId *string, placa *string, dataInicial *string, dataFinal *string) ([]entities.Diesel, error) {

	filterParams := filter.NewDieselFilterParams(fornecedorId, placa, dataInicial, dataFinal)

	dieselFilter, err := filterParams.ToFilter()
	if err != nil {
		return nil, err
	}

	return s.RepoManager.Diesel().Filter(*dieselFilter)
}
