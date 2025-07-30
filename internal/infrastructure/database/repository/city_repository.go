package repository

import (
	"database/sql"

	"github.com/smaelmr/finance-api/internal/domain/entities"
)

type CityRepository struct {
	conn *sql.DB
}

func newCityRepository(conn *sql.DB) *CityRepository {
	return &CityRepository{
		conn: conn,
	}
}

func (r *CityRepository) Add(city entities.City) error {
	query := `INSERT INTO cidade (nome, estado) VALUES (?, ?)`
	_, err := r.conn.Exec(query, city.Name)
	return err
}

func (r *CityRepository) GetAll() ([]entities.City, error) {
	query := `SELECT id, nome, estado FROM cidade ORDER BY nome`
	rows, err := r.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cities []entities.City
	for rows.Next() {
		var city entities.City
		if err := rows.Scan(&city.Id, &city.Name, &city.State); err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cities, nil
}
