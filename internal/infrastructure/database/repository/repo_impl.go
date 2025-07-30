package repository

import (
	"database/sql"

	"github.com/smaelmr/finance-api/internal/domain/contract/repository"
)

type Repo struct {
	conn *sql.DB
}

func NewRepo(conn *sql.DB) *Repo {
	return &Repo{
		conn: conn,
	}
}

func (c *Repo) Diesel() repository.DieselRepository {
	return newDieselRepository(c.conn)
}

func (c *Repo) Person() repository.PersonRepository {
	return newPersonRepository(c.conn)
}

func (c *Repo) City() repository.CityRepository {
	return newCityRepository(c.conn)
}

func (c *Repo) Freight() repository.FreightRepository {
	return newFreightRepository(c.conn)
}
