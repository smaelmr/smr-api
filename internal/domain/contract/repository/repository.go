package repository

import (
	"github.com/smaelmr/finance-api/internal/domain/entities"
	"github.com/smaelmr/finance-api/internal/domain/entities/filter"
)

type RepoManager interface {
	Diesel() DieselRepository
	Person() PersonRepository
	City() CityRepository
	Freight() FreightRepository
}

type DieselRepository interface {
	Add(diesel entities.Diesel) error
	GetAll() ([]entities.Diesel, error)
	Update(diesel entities.Diesel) error
	Filter(params filter.DieselFilter) ([]entities.Diesel, error)
}

type CityRepository interface {
	Add(city entities.City) error
	GetAll() ([]entities.City, error)
}

type PersonRepository interface {
	Add(person entities.Person) error
	GetSuppliers() ([]entities.Person, error)
	GetCustomers() ([]entities.Person, error)
	GetDrivers() ([]entities.Person, error)
}

type FreightRepository interface {
	Add(freight entities.Freight) error
	GetAll() ([]entities.Freight, error)
	Update(freight entities.Freight) error
	Filter(params filter.FreightFilter) ([]entities.Freight, error)
}
