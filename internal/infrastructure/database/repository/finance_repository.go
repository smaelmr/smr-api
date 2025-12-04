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

	// Tratar OrigemId: se for nil ou zero, salvar como NULL
	var origemId interface{}
	if record.OrigemId == nil || *record.OrigemId == 0 {
		origemId = nil
	} else {
		origemId = *record.OrigemId
	}

	// Tratar DataRealizacao: se for nil ou zero, salvar como NULL
	var dataRealizacao interface{}
	if record.DataRealizacao == nil || record.DataRealizacao.IsZero() {
		dataRealizacao = nil
	} else {
		dataRealizacao = *record.DataRealizacao
	}

	query := `INSERT INTO financeiro 
	(pessoa_id, valor_original, numero_documento, data_competencia, data_vencimento, 
	data_realizacao, origem, origem_id, observacao, numero_parcela, categoria_id)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := r.conn.Exec(query,
		record.PessoaId,
		record.ValorParcela,
		record.NumeroDocumento,
		record.DataCompetencia,
		record.DataVencimento,
		dataRealizacao,
		record.Origem,
		origemId,
		record.Observacao,
		record.NumeroParcela,
		record.CategoriaId)

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
		id, pessoa_id, valor_original, numero_documento, data_competencia, 
		data_vencimento, data_realizacao, origem, origem_id, observacao, 
		categoria_id, created_at, updated_at
	FROM financeiro WHERE id = ?`

	row := r.conn.QueryRow(query, id)

	var record entities.Finance
	var dataRealizacao sql.NullTime
	var origemId sql.NullInt64

	err := row.Scan(
		&record.Id,
		&record.PessoaId,
		&record.Valor,
		&record.NumeroDocumento,
		&record.DataCompetencia,
		&record.DataVencimento,
		&dataRealizacao,
		&record.Origem,
		&origemId,
		&record.Observacao,
		&record.CategoriaId,
		&record.CreatedAt,
		&record.UpdatedAt)
	if err != nil {
		return nil, err
	}

	// Converter sql.NullTime para *time.Time
	if dataRealizacao.Valid {
		record.DataRealizacao = &dataRealizacao.Time
	}

	// Converter sql.NullInt64 para *int64
	if origemId.Valid {
		record.OrigemId = &origemId.Int64
	}

	return &record, nil
}

func (r *FinanceRepository) GetAll(categoryType string, month int, year int) ([]entities.Finance, error) {
	query := `SELECT 
		f.id, f.pessoa_id, f.valor_original, f.numero_documento, f.data_competencia, 
		f.data_vencimento, f.data_realizacao, f.origem, f.origem_id, f.observacao, 
		categoria_id, f.numero_parcela, f.created_at, f.updated_at
	FROM financeiro f
	INNER JOIN categoria c ON f.categoria_id = c.id
	WHERE c.tipo = ?
		AND MONTH(f.data_competencia) = ?
		AND YEAR(f.data_competencia) = ?
	ORDER BY f.data_competencia DESC`

	rows, err := r.conn.Query(query, categoryType, month, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []entities.Finance
	for rows.Next() {
		var record entities.Finance
		var dataRealizacao sql.NullTime
		var origemId sql.NullInt64

		err := rows.Scan(
			&record.Id,
			&record.PessoaId,
			&record.Valor,
			&record.NumeroDocumento,
			&record.DataCompetencia,
			&record.DataVencimento,
			&dataRealizacao,
			&record.Origem,
			&origemId,
			&record.Observacao,
			&record.CategoriaId,
			&record.NumeroParcela,
			&record.CreatedAt,
			&record.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Converter sql.NullTime para *time.Time
		if dataRealizacao.Valid {
			record.DataRealizacao = &dataRealizacao.Time
		}

		// Converter sql.NullInt64 para *int64
		if origemId.Valid {
			record.OrigemId = &origemId.Int64
		}

		records = append(records, record)
	}

	return records, nil
}

func (r *FinanceRepository) Update(record entities.Finance) error {
	// Tratar OrigemId: se for nil ou zero, salvar como NULL
	var origemId interface{}
	if record.OrigemId == nil || *record.OrigemId == 0 {
		origemId = nil
	} else {
		origemId = *record.OrigemId
	}

	// Tratar DataRealizacao: se for nil ou zero, salvar como NULL
	var dataRealizacao interface{}
	if record.DataRealizacao == nil || record.DataRealizacao.IsZero() {
		dataRealizacao = nil
	} else {
		dataRealizacao = *record.DataRealizacao
	}

	query := `UPDATE financeiro SET 
		pessoa_id = ?,
		valor_original = ?,
		numero_documento = ?,
		data_competencia = ?,
		data_vencimento = ?,
		data_realizacao = ?,
		origem = ?,
		origem_id = ?,
		observacao = ?,
		categoria_id = ?,
		updated_at = ?
	WHERE id = ?`

	result, err := r.conn.Exec(query,
		record.PessoaId,
		record.Valor,
		record.NumeroDocumento,
		record.DataCompetencia,
		record.DataVencimento,
		dataRealizacao,
		record.Origem,
		origemId,
		record.Observacao,
		record.CategoriaId,
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
