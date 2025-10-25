package services

import (
	"github.com/smaelmr/finance-api/internal/domain/contract/repository"
	"github.com/smaelmr/finance-api/internal/domain/entities"
)

type PersonService struct {
	RepoManager repository.RepoManager
}

func NewPersonService(repoManager repository.RepoManager) *PersonService {
	return &PersonService{
		RepoManager: repoManager,
	}
}

func (s *PersonService) AddClient(personAdd *entities.Person) error {
	return s.RepoManager.Person().AddClient(*personAdd)
}

func (s *PersonService) AddSupplier(personAdd *entities.Person) error {
	return s.RepoManager.Person().AddSupplier(*personAdd)
}

func (s *PersonService) AddDriver(driver *entities.Driver) error {
	return s.RepoManager.Person().AddDriver(*driver)
}

func (s *PersonService) AddGasStation(personAdd *entities.Person) error {
	return s.RepoManager.Person().AddGasStation(*personAdd)
}

func (s *PersonService) GetClients() ([]entities.Person, error) {
	records, err := s.RepoManager.Person().GetClients()
	if err != nil {
		return nil, err
	}

	var customerList []entities.Person
	customerList = append(customerList, records...)

	return customerList, nil
}

func (s *PersonService) GetSuppliers() ([]entities.Person, error) {
	records, err := s.RepoManager.Person().GetSuppliers()
	if err != nil {
		return nil, err
	}

	var suppliersList []entities.Person
	suppliersList = append(suppliersList, records...)

	return suppliersList, nil
}

func (s *PersonService) GetSupplierByCnpj(cnpj string) (*entities.Person, error) {
	person, err := s.RepoManager.Person().GetSupplierByCnpj(cnpj)
	if err != nil {
		return nil, err
	}

	return &person, nil
}

func (s *PersonService) GetDrivers() ([]entities.Driver, error) {
	return s.RepoManager.Person().GetDrivers()
}

func (s *PersonService) GetGasStation() ([]entities.Person, error) {
	records, err := s.RepoManager.Person().GetGasStations()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (s *PersonService) UpdateGasStation(person *entities.Person) error {
	return s.RepoManager.Person().UpdateGasStation(*person)
}

func (s *PersonService) DeleteGasStation(id int64) error {
	return s.RepoManager.Person().DeleteGasStation(id)
}

func (s *PersonService) UpdateClient(person *entities.Person) error {
	return s.RepoManager.Person().UpdateClient(*person)
}

func (s *PersonService) DeleteClient(id int64) error {
	return s.RepoManager.Person().DeleteClient(id)
}

func (s *PersonService) UpdateSupplier(person *entities.Person) error {
	return s.RepoManager.Person().UpdateSupplier(*person)
}

func (s *PersonService) DeleteSupplier(id int64) error {
	return s.RepoManager.Person().DeleteSupplier(id)
}

func (s *PersonService) UpdateDriver(driver *entities.Driver) error {
	return s.RepoManager.Person().UpdateDriver(*driver)
}

func (s *PersonService) DeleteDriver(id int64) error {
	return s.RepoManager.Person().DeleteDriver(id)
}
