package repository

import (
	"database/sql"
	"strings"

	"github.com/smaelmr/finance-api/internal/domain/entities"
	"github.com/smaelmr/finance-api/internal/domain/entities/filter"
)

type DieselRepository struct {
	conn *sql.DB
}

func newDieselRepository(conn *sql.DB) *DieselRepository {
	return &DieselRepository{
		conn: conn,
	}
}

func (r *DieselRepository) Add(record entities.Diesel) error {
	query :=
		`INSERT INTO abastecimento
		(fornecedor_id, data_abastecimento, qtd_diesel,
		total_arla, placa_veiculo, km, total_diesel)
		VALUES (?, ?, ?, ?, ?, ?, ?);`

	_, err := r.conn.Exec(query,
		record.FornecedorId,
		record.Data,
		record.Quantidade,
		record.ArlaTotal,
		record.Placa,
		record.Km,
		record.DieselTotal)
	return err
}

func (r *DieselRepository) GetAll() ([]entities.Diesel, error) {
	query := `SELECT 
			a.id, 
			a.data_abastecimento, 
			a.qtd_diesel,
			a.total_diesel, 
			a.fornecedor_id, 
			a.total_arla, 
			a.placa_veiculo, 
			a.km,
			a.total_diesel,
			p.nome
		FROM abastecimento a
		INNER JOIN fornecedor f ON a.fornecedor_id = f.id
		INNER JOIN pessoa p ON p.id = f.pessoa_id`

	rows, err := r.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []entities.Diesel
	for rows.Next() {
		var record entities.Diesel
		if err := rows.Scan(
			&record.Id,
			&record.Data,
			&record.Quantidade,
			&record.DieselTotal,
			&record.FornecedorId,
			&record.ArlaTotal,
			&record.Placa,
			&record.Km,
			&record.DieselTotal,
			&record.FornecedorName); err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func (r *DieselRepository) Update(diesel entities.Diesel) error {
	query := `UPDATE abastecimento SET
		fornecedor_id = ?,
		data_abastecimento = ?,
		qtd_diesel = ?,
		total_arla = ?,
		placa_veiculo = ?,
		km = ?,
		total_diesel = ?
		WHERE id = ?`

	_, err := r.conn.Exec(query,
		diesel.FornecedorId,
		diesel.Data,
		diesel.Quantidade,
		diesel.ArlaTotal,
		diesel.Placa,
		diesel.Km,
		diesel.DieselTotal,
		diesel.Id)

	return err
}

func (r *DieselRepository) Filter(params filter.DieselFilter) ([]entities.Diesel, error) {
	query := `SELECT 
        a.id, 
        a.data_abastecimento, 
        a.qtd_diesel,
        a.total_diesel, 
        a.fornecedor_id, 
        a.total_arla, 
        a.placa_veiculo, 
        a.km,
        a.total_diesel,
        p.nome
    FROM abastecimento a
    INNER JOIN fornecedor f ON a.fornecedor_id = f.id
    INNER JOIN pessoa p ON p.id = f.pessoa_id
    WHERE 1=1`

	var conditions []string
	var args []interface{}

	if params.FornecedorId != nil {
		conditions = append(conditions, "a.fornecedor_id = ?")
		args = append(args, *params.FornecedorId)
	}

	if params.Placa != nil {
		conditions = append(conditions, "a.placa_veiculo = ?")
		args = append(args, *params.Placa)
	}

	if params.DataInicial != nil {
		conditions = append(conditions, "a.data_abastecimento >= ?")
		args = append(args, *params.DataInicial)
	}

	if params.DataFinal != nil {
		conditions = append(conditions, "a.data_abastecimento <= ?")
		args = append(args, *params.DataFinal)
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	query += " ORDER BY a.data_abastecimento DESC"

	rows, err := r.conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []entities.Diesel
	for rows.Next() {
		var record entities.Diesel
		if err := rows.Scan(
			&record.Id,
			&record.Data,
			&record.Quantidade,
			&record.DieselTotal,
			&record.FornecedorId,
			&record.ArlaTotal,
			&record.Placa,
			&record.Km,
			&record.DieselTotal,
			&record.FornecedorName); err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}
