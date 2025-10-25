package services

import (
	"github.com/smaelmr/finance-api/internal/domain/contract/repository"
	"github.com/smaelmr/finance-api/internal/domain/entities"
)

type VehicleService struct {
	RepoManager repository.RepoManager
}

func NewVehicleService(repoManager repository.RepoManager) *VehicleService {
	return &VehicleService{
		RepoManager: repoManager,
	}
}

func (s *VehicleService) Create(record entities.Vehicle) error {
	return s.RepoManager.Vehicle().Create(record)
}

func (s *VehicleService) Get(id int64) (*entities.Vehicle, error) {
	return s.RepoManager.Vehicle().Get(id)
}

func (s *VehicleService) GetAll() ([]entities.Vehicle, error) {
	return s.RepoManager.Vehicle().GetAll()
}

func (s *VehicleService) Update(record entities.Vehicle) error {
	return s.RepoManager.Vehicle().Update(record)
}

func (s *VehicleService) Delete(id int64) error {
	return s.RepoManager.Vehicle().Delete(id)
}

func (s *VehicleService) GetByPlate(plate string) (*entities.Vehicle, error) {
	return s.RepoManager.Vehicle().GetByPlate(plate)
}
