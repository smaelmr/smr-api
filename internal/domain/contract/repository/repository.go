package repository

import (
	"time"

	"github.com/smaelmr/finance-api/internal/domain/entities"
	"github.com/smaelmr/finance-api/internal/domain/entities/dto"
	"github.com/smaelmr/finance-api/internal/domain/entities/filter"
)

type RepoManager interface {
	Fueling() FuelingRepository
	Person() PersonRepository
	City() CityRepository
	Trip() TripRepository
	Finance() FinanceRepository
	Vehicle() VehicleRepository
	Category() CategoryRepository
}

type FuelingRepository interface {
	Add(diesel entities.Fueling) error
	GetAll() ([]entities.Fueling, error)
	Update(diesel entities.Fueling) error
	Delete(id int64) error
	GetByDateRange(startDate, endDate time.Time) ([]entities.Fueling, error)
	GetFuelConsumption(startDate, endDate time.Time) ([]dto.FuelingConsumption, error)
}

type CityRepository interface {
	Add(city entities.City) error
	GetAll() ([]entities.City, error)
}

type PersonRepository interface {
	AddSupplier(person entities.Person) error
	AddClient(person entities.Person) error
	AddDriver(driver entities.Driver) error
	AddGasStation(person entities.Person) error
	GetSuppliers() ([]entities.Person, error)
	GetClients() ([]entities.Person, error)
	GetDrivers() ([]entities.Driver, error)
	GetGasStations() ([]entities.Person, error)
	UpdateGasStation(person entities.Person) error
	DeleteGasStation(id int64) error
	GetGasStationByCnpj(string) (entities.Person, error)
	UpdateClient(person entities.Person) error
	DeleteClient(id int64) error
	UpdateSupplier(person entities.Person) error
	DeleteSupplier(id int64) error
	UpdateDriver(driver entities.Driver) error
	DeleteDriver(id int64) error
}

type TripRepository interface {
	Add(trip entities.Trip) error
	GetAll() ([]entities.Trip, error)
	Update(trip entities.Trip) error
	GetByDateRange(startDate, endDate time.Time) ([]entities.Trip, error)
	Filter(params filter.TripFilter) ([]entities.Trip, error)
}

type FinanceRepository interface {
	Add(record entities.Finance) error
	Update(record entities.Finance) error
	Get(int64) (*entities.Finance, error)
	GetAll(categoryType string, month int, year int) ([]entities.Finance, error)
	Delete(int64) error
}

type VehicleRepository interface {
	Create(vehicle entities.Vehicle) error
	Get(id int64) (*entities.Vehicle, error)
	GetAll() ([]entities.Vehicle, error)
	Update(vehicle entities.Vehicle) error
	Delete(id int64) error
	GetByPlate(plate string) (*entities.Vehicle, error)
}

type CategoryRepository interface {
	Add(category entities.Category) error
	GetAll() ([]entities.Category, error)
	GetByType(categoryType string) ([]entities.Category, error)
	Get(id int64) (*entities.Category, error)
	Delete(id int64) error
}
