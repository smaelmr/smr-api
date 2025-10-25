package repository

import (
	"database/sql"
	"time"

	"github.com/smaelmr/finance-api/internal/domain/entities"
)

type FinanceRepository struct {
	conn *sql.DB
}

func newFinanceRepository(conn *sql.DB) *FinanceRepository {
	return &FinanceRepository{
		conn: conn,
	}
}

func (r *FinanceRepository) Add(record entities.Finance) error {
	now := time.Now()
	record.CreatedAt = now
	record.UpdatedAt = now

	query := `INSERT INTO financeiro 
	(pessoa_id, valor, numero_documento, data_lancamento, data_vencimento, 
	data_realizacao, origem, observacao, recebido, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := r.conn.Exec(query,
		record.PessoaId,
		record.Valor,
		record.NumeroDocumento,
		record.DataLancamento,
		record.DataVencimento,
		record.DataRealizacao,
		record.Origem,
		record.Observacao,
		record.Recebido,
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

func (r *FinanceRepository) Get(id int64) (*entities.Finance, error) {
	query := `SELECT 
		id, pessoa_id, valor, numero_documento, data_lancamento, 
		data_vencimento, data_realizacao, origem, observacao, 
		recebido, created_at, updated_at
	FROM financeiro WHERE id = ?`

	row := r.conn.QueryRow(query, id)

	var record entities.Finance
	err := row.Scan(
		&record.Id,
		&record.PessoaId,
		&record.Valor,
		&record.NumeroDocumento,
		&record.DataLancamento,
		&record.DataVencimento,
		&record.DataRealizacao,
		&record.Origem,
		&record.Observacao,
		&record.Recebido,
		&record.CreatedAt,
		&record.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (r *FinanceRepository) GetAll() ([]entities.Finance, error) {
	query := `SELECT 
		id, pessoa_id, valor, numero_documento, data_lancamento, 
		data_vencimento, data_realizacao, origem, observacao, 
		recebido, created_at, updated_at
	FROM financeiro`

	rows, err := r.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []entities.Finance
	for rows.Next() {
		var record entities.Finance
		err := rows.Scan(
			&record.Id,
			&record.PessoaId,
			&record.Valor,
			&record.NumeroDocumento,
			&record.DataLancamento,
			&record.DataVencimento,
			&record.DataRealizacao,
			&record.Origem,
			&record.Observacao,
			&record.Recebido,
			&record.CreatedAt,
			&record.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func (r *FinanceRepository) Update(record entities.Finance) error {
	query := `UPDATE financeiro SET 
		pessoa_id = ?,
		valor = ?,
		numero_documento = ?,
		data_lancamento = ?,
		data_vencimento = ?,
		data_realizacao = ?,
		origem = ?,
		observacao = ?,
		recebido = ?,
		updated_at = ?
	WHERE id = ?`

	result, err := r.conn.Exec(query,
		record.PessoaId,
		record.Valor,
		record.NumeroDocumento,
		record.DataLancamento,
		record.DataVencimento,
		record.DataRealizacao,
		record.Origem,
		record.Observacao,
		record.Recebido,
		time.Now(),
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

func (r *FinanceRepository) Delete(id int64) error {
	query := `DELETE FROM financeiro WHERE id = ?`

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
