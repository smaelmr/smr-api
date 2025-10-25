package repository

import (
	"database/sql"

	"github.com/smaelmr/finance-api/internal/domain/entities"
)

type FuelingRepository struct {
	conn *sql.DB
}

func newFuelingRepository(conn *sql.DB) *FuelingRepository {
	return &FuelingRepository{
		conn: conn,
	}
}

func (r *FuelingRepository) Add(record entities.Fueling) error {
	query :=
		`INSERT INTO abastecimento
		(veiculo_id, posto_id, data_abastecimento, tipo_combustivel,
		litros, valor_unitario, valor_total, numero_documento, km)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);`

	result, err := r.conn.Exec(query,
		record.VeiculoId,
		record.PostoId,
		record.Data,
		record.TipoCombustivel,
		record.Litros,
		record.ValorUnitario,
		record.ValorTotal,
		record.NumeroDocumento,
		record.Km)

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

func (r *FuelingRepository) GetAll() ([]entities.Fueling, error) {
	query := `SELECT 
				a.id, a.veiculo_id, a.posto_id, a.data_abastecimento, 
				a.tipo_combustivel, a.litros, a.valor_unitario, a.valor_total,
				a.Km, a.numero_documento, a.created_at, a.updated_at
			FROM abastecimento a
			INNER JOIN posto f ON a.posto_id = f.id
			INNER JOIN pessoa p ON p.id = f.pessoa_id`

	rows, err := r.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []entities.Fueling
	for rows.Next() {
		var record entities.Fueling
		if err := rows.Scan(
			&record.Id,
			&record.VeiculoId,
			&record.PostoId,
			&record.Data,
			&record.TipoCombustivel,
			&record.Litros,
			&record.ValorUnitario,
			&record.ValorTotal,
			&record.Km,
			&record.NumeroDocumento,
			&record.CreatedAt,
			&record.UpdatedAt); err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func (r *FuelingRepository) Update(record entities.Fueling) error {

	query := `UPDATE abastecimento 
		SET
			veiculo_id = ?,
			posto_id = ?,
			data_abastecimento = ?,
			tipo_combustivel = ?,
			km = ?,
			litros = ?,
			valor_unitario = ?,
			valor_total = ?,
			numero_documento = ?
		WHERE id = ?`

	_, err := r.conn.Exec(query,
		record.VeiculoId,
		record.PostoId,
		record.Data,
		record.TipoCombustivel,
		record.Km,
		record.Litros,
		record.ValorUnitario,
		record.ValorTotal,
		record.NumeroDocumento,
		record.Id)

	if err != nil {
		return err
	}

	//rowsAffected, err = result.RowsAffected()
	//if err != nil {
	//	return err
	//}

	//if rowsAffected == 0 {
	//	return sql.ErrNoRows
	//}

	return nil
}

func (r *FuelingRepository) Delete(id int64) error {
	query := `DELETE FROM abastecimento WHERE id = ?`

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

/*func (r *FuelingRepository) Filter(params filter.FuelingFilter) ([]entities.Fueling, error) {
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

	var records []entities.Fueling
	for rows.Next() {
		var record entities.Fueling
		if err := rows.Scan(
			&record.Id,
			&record.Data,
			&record.Quantidade,
			&record.FuelingTotal,
			&record.FornecedorId,
			&record.ArlaTotal,
			&record.Placa,
			&record.Km,
			&record.FuelingTotal,
			&record.FornecedorName); err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}*/
