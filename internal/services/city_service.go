package services

import (
	"github.com/smaelmr/finance-api/internal/domain/contract/repository"
	"github.com/smaelmr/finance-api/internal/domain/entities"
)

type CityService struct {
	RepoManager repository.RepoManager
}

func NewCityService(repoManager repository.RepoManager) *CityService {
	return &CityService{
		RepoManager: repoManager,
	}
}

func (s *CityService) Add(city entities.City) error {
	/*if city == nil {
		return repository.ErrInvalidEntity
	}

	if city.Name == "" || city.State == "" {
		return repository.ErrInvalidEntity
	}*/

	return s.RepoManager.City().Add(city)
}

func (s *CityService) GetAll() ([]entities.City, error) {
	cities, err := s.RepoManager.City().GetAll()
	if err != nil {
		return nil, err
	}

	var cityList []entities.City
	cityList = append(cityList, cities...)

	return cityList, nil
}
