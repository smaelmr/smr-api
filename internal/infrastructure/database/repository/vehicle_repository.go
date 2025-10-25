package repository

import (
	"database/sql"
	"time"

	"github.com/smaelmr/finance-api/internal/domain/entities"
)

type VehicleRepository struct {
	conn *sql.DB
}

func NewVehicleRepository(conn *sql.DB) *VehicleRepository {
	return &VehicleRepository{
		conn: conn,
	}
}

func (r *VehicleRepository) Create(record entities.Vehicle) error {
	now := time.Now()
	record.CreatedAt = now
	record.UpdatedAt = now

	query := `INSERT INTO veiculo 
		(placa, marca, modelo, ano, tipo, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)`

	result, err := r.conn.Exec(query,
		record.Placa,
		record.Marca,
		record.Modelo,
		record.Ano,
		record.Tipo,
		record.CreatedAt,
		record.UpdatedAt)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	record.Id = id
	return nil
}

func (r *VehicleRepository) Get(id int64) (*entities.Vehicle, error) {
	query := `SELECT 
		id, placa, marca, modelo, ano, tipo, created_at, updated_at
		FROM veiculo WHERE id = ?`

	row := r.conn.QueryRow(query, id)

	var record entities.Vehicle
	err := row.Scan(
		&record.Id,
		&record.Placa,
		&record.Marca,
		&record.Modelo,
		&record.Ano,
		&record.Tipo,
		&record.CreatedAt,
		&record.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (r *VehicleRepository) GetAll() ([]entities.Vehicle, error) {
	query := `SELECT 
		id, placa, marca, modelo, ano, tipo, created_at, updated_at
		FROM veiculo`

	rows, err := r.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []entities.Vehicle
	for rows.Next() {
		var record entities.Vehicle
		err := rows.Scan(
			&record.Id,
			&record.Placa,
			&record.Marca,
			&record.Modelo,
			&record.Ano,
			&record.Tipo,
			&record.CreatedAt,
			&record.UpdatedAt)

		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func (r *VehicleRepository) Update(record entities.Vehicle) error {
	record.UpdatedAt = time.Now()

	query := `UPDATE veiculo SET 
		placa = ?,
		marca = ?,
		modelo = ?,
		ano = ?,
		tipo = ?,
		updated_at = ?
		WHERE id = ?`

	result, err := r.conn.Exec(query,
		record.Placa,
		record.Marca,
		record.Modelo,
		record.Ano,
		record.Tipo,
		record.UpdatedAt,
		record.Id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *VehicleRepository) Delete(id int64) error {
	query := `DELETE FROM veiculo WHERE id = ?`

	result, err := r.conn.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
