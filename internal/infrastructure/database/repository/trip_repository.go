package repository

import (
	"database/sql"
	"time"

	"github.com/smaelmr/finance-api/internal/domain/entities"
	"github.com/smaelmr/finance-api/internal/domain/entities/filter"
)

type TripRepository struct {
	conn *sql.DB
}

func newTripRepository(conn *sql.DB) *TripRepository {
	return &TripRepository{
		conn: conn,
	}
}

func (r *TripRepository) GetByDateRange(startDate, endDate time.Time) ([]entities.Trip, error) {
	query := `SELECT 
		f.id, f.cavalo_placa, f.cliente_id, f.origem_id, f.destino_final_id, 
		f.motorista_id, f.data_carregamento, f.data_entrega, 
		f.numero_documento, f.valor_agenciamento, f.valor_frete, 
		f.valor_pedagio, f.observacoes
	FROM frete f
	WHERE f.data_carregamento BETWEEN ? AND ?
	ORDER BY f.data_carregamento`

	rows, err := r.conn.Query(query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []entities.Trip
	for rows.Next() {
		var record entities.Trip
		if err := rows.Scan(
			&record.Id,
			&record.CavaloPlaca,
			&record.ClienteId,
			&record.OrigemId,
			&record.DestinoId,
			&record.MotoristaId,
			&record.DataCarregamento,
			&record.DataEntrega,
			&record.NumeroDocumento,
			&record.ValorAgenciamento,
			&record.ValorFrete,
			&record.ValorPedagio,
			&record.Observacoes); err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func (r *TripRepository) Add(trip entities.Trip) error {
	query :=
		`INSERT INTO frete 
			(cavalo_placa, cliente_id, origem_id, destino_final_id, 
			motorista_id, data_carregamento, data_entrega, numero_documento, 
			valor_agenciamento, valor_frete, valor_pedagio, observacoes)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.conn.Exec(query,
		trip.CavaloPlaca,
		trip.ClienteId,
		trip.OrigemId,
		trip.DestinoId,
		trip.MotoristaId,
		trip.DataCarregamento,
		trip.DataEntrega,
		trip.NumeroDocumento,
		trip.ValorAgenciamento,
		trip.ValorFrete,
		trip.ValorPedagio,
		trip.Observacoes)
	return err
}

func (r *TripRepository) GetTripRecord() (*entities.Trip, error) {
	query :=
		`SELECT id, cavalo_placa, cliente_id, origem_id,
			 destino_final_id, forma_pagamento_id, motorista_id,
			 data_carregamento, data_entrega, numero_documento,
			 valor_agenciamento, valor_frete, valor_pedagio, observacoes
		 FROM frete LIMIT 1;`

	row := r.conn.QueryRow(query)

	var record entities.Trip
	err := row.Scan(
		&record.Id,
		&record.CavaloPlaca,
		&record.ClienteId,
		&record.OrigemId,
		&record.DestinoId,
		&record.MotoristaId,
		&record.DataCarregamento,
		&record.DataEntrega,
		&record.NumeroDocumento,
		&record.ValorAgenciamento,
		&record.ValorFrete,
		&record.ValorPedagio,
		&record.Observacoes)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (r *TripRepository) GetAll() ([]entities.Trip, error) {
	query := `SELECT 
				f.id, f.cavalo_placa, f.cliente_id, origem_id, f.destino_final_id, 
				f.motorista_id, f.data_carregamento, f.data_entrega, f.numero_documento, 
				f.valor_agenciamento, f.valor_frete, f.valor_pedagio, f.observacoes
				FROM frete f`

	rows, err := r.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []entities.Trip
	for rows.Next() {
		var record entities.Trip
		if err := rows.Scan(
			&record.Id,
			&record.CavaloPlaca,
			&record.ClienteId,
			&record.OrigemId,
			&record.DestinoId,
			&record.MotoristaId,
			&record.DataCarregamento,
			&record.DataEntrega,
			&record.NumeroDocumento,
			&record.ValorAgenciamento,
			&record.ValorFrete,
			&record.ValorPedagio,
			&record.Observacoes); err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func (r *TripRepository) Update(trip entities.Trip) error {
	query := `UPDATE frete SET 
		valo_placa = ?, 
		cliente_id = ?, 
		origem_id = ?,
		destino_final_id = ?, 
		motorista_id = ?,
		data_carregamento = ?, 
		data_entrega = ?, 
		numero_documento = ?,
		valor_agenciamento = ?, 
		valor_frete = ?, 
		valor_pedagio = ?, 
		observacoes = ?
		WHERE id = ?`

	_, err := r.conn.Exec(query,
		trip.CavaloPlaca,
		trip.ClienteId,
		trip.OrigemId,
		trip.DestinoId,
		trip.MotoristaId,
		trip.DataCarregamento,
		trip.DataEntrega,
		trip.NumeroDocumento,
		trip.ValorAgenciamento,
		trip.ValorFrete,
		trip.ValorPedagio,
		trip.Observacoes,
		trip.Id)

	return err
}

func (r *TripRepository) Filter(params filter.TripFilter) ([]entities.Trip, error) {
	query := `f.id, f.cavalo_placa, f.cliente_id, origem_id, f.destino_final_id, 
				f.motorista_id, f.data_carregamento, f.data_entrega, f.numero_documento, 
				f.valor_agenciamento, f.valor_frete, f.valor_pedagio, f.observacoes
				FROM frete f WHERE 1=1`

	args := []interface{}{}

	if params.ClienteId != nil {
		query += " AND f.cliente_id = ?"
		args = append(args, *params.ClienteId)
	}

	if params.MotoristaId != nil {
		query += " AND f.motorista_id = ?"
		args = append(args, *params.MotoristaId)
	}

	if params.DataInicial != nil {
		query += " AND f.data_carregamento >= ?"
		args = append(args, *params.DataInicial)
	}

	if params.DataFinal != nil {
		query += " AND f.data_carregamento <= ?"
		args = append(args, *params.DataFinal)
	}

	if params.CavaloPlaca != nil {
		query += " AND f.cavalo_placa = ?"
		args = append(args, *params.CavaloPlaca)
	}

	rows, err := r.conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []entities.Trip
	for rows.Next() {
		var record entities.Trip
		if err := rows.Scan(
			&record.Id,
			&record.CavaloPlaca,
			&record.ClienteId,
			&record.OrigemId,
			&record.MotoristaId,
			&record.DataCarregamento,
			&record.DataEntrega,
			&record.NumeroDocumento,
			&record.ValorAgenciamento,
			&record.ValorFrete,
			&record.ValorPedagio,
			&record.Observacoes); err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func (r *TripRepository) Delete(id int64) error {
	query := "DELETE FROM frete WHERE id = ?"
	_, err := r.conn.Exec(query, id)
	return err
}
