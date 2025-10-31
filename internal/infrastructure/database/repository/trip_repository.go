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
		f.id, f.carreta_placa, f.cavalo_placa, f.cliente_id, 
		f.origem_id, f.destino_final_id, f.forma_pagamento_id, 
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
			&record.CarretaPlaca,
			&record.CavaloPlaca,
			&record.ClienteId,
			&record.OrigemId,
			&record.DestinoFinalId,
			&record.FormaPagamentoId,
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
			(carreta_placa, cavalo_placa, cliente_id, origem_id, 
            destino_final_id, forma_pagamento_id, motorista_id, 
            data_carregamento, data_entrega, numero_documento, 
            valor_agenciamento, valor_frete, valor_pedagio, observacoes)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.conn.Exec(query,
		trip.CarretaPlaca,
		trip.CavaloPlaca,
		trip.ClienteId,
		trip.OrigemId,
		trip.DestinoFinalId,
		trip.FormaPagamentoId,
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
		`SELECT id, carreta_placa, cavalo_placa, cliente_id, origem_id,
			 destino_final_id, forma_pagamento_id, motorista_id,
			 data_carregamento, data_entrega, numero_documento,
			 valor_agenciamento, valor_frete, valor_pedagio, observacoes
		 FROM frete LIMIT 1;`

	row := r.conn.QueryRow(query)

	var record entities.Trip
	err := row.Scan(
		&record.Id,
		&record.CarretaPlaca,
		&record.CavaloPlaca,
		&record.ClienteId,
		&record.OrigemId,
		&record.DestinoFinalId,
		&record.FormaPagamentoId,
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
				f.id, f.carreta_placa, f.cavalo_placa, f.cliente_id, pc.nome AS cliente_nome,
				f.origem_id, o.nome AS origem_nome, f.destino_final_id, d.nome AS destino_final_nome, 
				f.forma_pagamento_id, fp.descricao AS forma_pagamento_descricao, f.motorista_id, 
				pm.nome AS motorista_nome, f.data_carregamento, f.data_entrega, f.numero_documento, 
				f.valor_agenciamento, f.valor_frete, f.valor_pedagio, f.observacoes
				FROM frete f
				INNER JOIN cliente c ON f.cliente_id = c.id
				INNER JOIN pessoa pc ON pc.id = c.pessoa_id
				INNER JOIN motorista m ON m.id = f.motorista_id
				INNER JOIN pessoa pm ON pm.id = m.pessoa_id
				INNER JOIN cidade o ON f.origem_id = o.id
				INNER JOIN cidade d ON f.destino_final_id = d.id
				INNER JOIN forma_pagamento fp ON f.forma_pagamento_id = fp.id`

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
			&record.CarretaPlaca,
			&record.CavaloPlaca,
			&record.ClienteId,
			&record.ClienteNome,
			&record.OrigemId,
			&record.OrigemNome,
			&record.DestinoFinalId,
			&record.DestinoFinalNome,
			&record.FormaPagamentoId,
			&record.FormaPagamentoDescricao,
			&record.MotoristaId,
			&record.MotoristaNome,
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
		carreta_placa = ?, 
		cavalo_placa = ?, 
		cliente_id = ?, 
		origem_id = ?,
		destino_final_id = ?, 
		forma_pagamento_id = ?, 
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
		trip.CarretaPlaca,
		trip.CavaloPlaca,
		trip.ClienteId,
		trip.OrigemId,
		trip.DestinoFinalId,
		trip.FormaPagamentoId,
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
	query := `SELECT 
		f.id, f.carreta_placa, f.cavalo_placa, f.cliente_id, pc.nome AS cliente_nome,
		f.origem_id, o.nome AS origem_nome, f.destino_final_id, d.nome AS destino_final_nome, 
		f.forma_pagamento_id, fp.descricao AS forma_pagamento_descricao, f.motorista_id, 
		pm.nome AS motorista_nome, f.data_carregamento, f.data_entrega, f.numero_documento, 
		f.valor_agenciamento, f.valor_frete, f.valor_pedagio, f.observacoes
		FROM frete f
		INNER JOIN cliente c ON f.cliente_id = c.id
		INNER JOIN pessoa pc ON pc.id = c.pessoa_id
		INNER JOIN motorista m ON m.id = f.motorista_id
		INNER JOIN pessoa pm ON pm.id = m.pessoa_id
		INNER JOIN cidade o ON f.origem_id = o.id
		INNER JOIN cidade d ON f.destino_final_id = d.id
		INNER JOIN forma_pagamento fp ON f.forma_pagamento_id = fp.id
		WHERE 1=1`

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
			&record.CarretaPlaca,
			&record.CavaloPlaca,
			&record.ClienteId,
			&record.ClienteNome,
			&record.OrigemId,
			&record.OrigemNome,
			&record.DestinoFinalId,
			&record.DestinoFinalNome,
			&record.FormaPagamentoId,
			&record.FormaPagamentoDescricao,
			&record.MotoristaId,
			&record.MotoristaNome,
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
