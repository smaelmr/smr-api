package repository

import (
	"database/sql"

	"github.com/smaelmr/finance-api/internal/domain/contract/repository"
)

type Repo struct {
	conn        *sql.DB
	vehicleRepo repository.VehicleRepository
}

func NewRepo(conn *sql.DB) *Repo {
	repo := &Repo{
		conn: conn,
	}
	repo.vehicleRepo = NewVehicleRepository(conn)
	return repo
}

func (c *Repo) Fueling() repository.FuelingRepository {
	return newFuelingRepository(c.conn)
}

func (c *Repo) Person() repository.PersonRepository {
	return newPersonRepository(c.conn)
}

func (c *Repo) City() repository.CityRepository {
	return newCityRepository(c.conn)
}

func (c *Repo) Vehicle() repository.VehicleRepository {
	return c.vehicleRepo
}

func (c *Repo) Trip() repository.TripRepository {
	return newTripRepository(c.conn)
}

func (c *Repo) Finance() repository.FinanceRepository {
	return newFinanceRepository(c.conn)
}

func (c *Repo) Category() repository.CategoryRepository {
	return newCategoryRepository(c.conn)
}
