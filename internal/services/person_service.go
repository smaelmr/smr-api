package services

import (
	"github.com/smaelmr/finance-api/api/commands"
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

func (s *PersonService) Add(personAdd *commands.PersonAdd) error {

	person := entities.Person{
		Name: personAdd.Name,
		Type: personAdd.Type,
	}

	return s.RepoManager.Person().Add(person)
}

func (s *PersonService) GetCustomers() ([]entities.Person, error) {
	records, err := s.RepoManager.Person().GetCustomers()
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

func (s *PersonService) GetDrivers() ([]entities.Person, error) {
	records, err := s.RepoManager.Person().GetDrivers()
	if err != nil {
		return nil, err
	}

	var driversList []entities.Person
	driversList = append(driversList, records...)

	return driversList, nil
}
