package repository

import (
	"database/sql"
	"time"

	"github.com/smaelmr/finance-api/internal/domain/entities"
	"github.com/smaelmr/finance-api/internal/domain/entities/dto"
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
		litros, valor_unitario, valor_total, numero_documento, km, cheio)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

	result, err := r.conn.Exec(query,
		record.VeiculoId,
		record.PostoId,
		record.Data,
		record.TipoCombustivel,
		record.Litros,
		record.ValorUnitario,
		record.ValorTotal,
		record.NumeroDocumento,
		record.Km,
		record.Cheio)

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
				a.Km, a.numero_documento, a.cheio, a.created_at, a.updated_at
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
			&record.Cheio,
			&record.CreatedAt,
			&record.UpdatedAt); err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func (r *FuelingRepository) GetByDateRange(startDate, endDate time.Time) ([]entities.Fueling, error) {
	query := `SELECT 
        a.id, a.veiculo_id, a.posto_id, a.data_abastecimento, 
        a.tipo_combustivel, a.litros, a.valor_unitario, a.valor_total,
        a.Km, a.numero_documento, a.cheio, a.created_at, a.updated_at
    FROM abastecimento a
    INNER JOIN posto f ON a.posto_id = f.id
    INNER JOIN pessoa p ON p.id = f.pessoa_id
    WHERE a.data_abastecimento BETWEEN ? AND ?
    ORDER BY a.data_abastecimento`

	rows, err := r.conn.Query(query, startDate, endDate)
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
			&record.Cheio,
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
			numero_documento = ?,
			cheio = ?
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
		record.Cheio,
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

func (r *FuelingRepository) GetFuelConsumption(startDate, endDate time.Time) ([]dto.FuelingConsumption, error) {
	query := `
		WITH consumo_cheio AS (
			SELECT 
				veiculo_id,
				MIN(CASE WHEN cheio = true THEN km END) as km_inicial_cheio,
				MAX(CASE WHEN cheio = true THEN km END) as km_final_cheio
			FROM abastecimento
			WHERE data_abastecimento BETWEEN ? AND ?
				AND tipo_combustivel IN ('Diesel_S10', 'Diesel_S500')
				AND cheio = true
			GROUP BY veiculo_id
		)
		SELECT 
			a.veiculo_id,
			v.placa,
			SUM(a.litros) as total_litros,
			COALESCE(cc.km_final_cheio - cc.km_inicial_cheio, 0) as total_km,
			COUNT(*) as qtd_abastecimentos,
			COUNT(CASE WHEN a.cheio = true THEN 1 END) as qtd_abastecimentos_cheio,
			SUM(CASE WHEN a.cheio = false THEN a.litros ELSE 0 END) as litros_nao_cheio
		FROM abastecimento a
		INNER JOIN veiculo v ON v.id = a.veiculo_id
		LEFT JOIN consumo_cheio cc ON cc.veiculo_id = a.veiculo_id
		WHERE 
			a.data_abastecimento BETWEEN ? AND ?
			AND a.tipo_combustivel IN ('Diesel_S10', 'Diesel_S500')
		GROUP BY a.veiculo_id, v.placa, cc.km_inicial_cheio, cc.km_final_cheio
		HAVING COUNT(*) > 1
		ORDER BY v.placa`

	rows, err := r.conn.Query(query, startDate, endDate, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var consumos []dto.FuelingConsumption
	for rows.Next() {
		var consumo dto.FuelingConsumption
		var litrosNaoCheio float64
		var qtdAbastecimentosCheio int
		if err := rows.Scan(
			&consumo.VeiculoId,
			&consumo.Placa,
			&consumo.TotalLitros,
			&consumo.TotalKm,
			&consumo.QtdAbastecimentos,
			&qtdAbastecimentosCheio,
			&litrosNaoCheio); err != nil {
			return nil, err
		}

		// Calcula a média em km/l considerando apenas a quilometragem entre abastecimentos cheios
		if consumo.TotalLitros > 0 && qtdAbastecimentosCheio >= 2 {
			// Usa somente litros entre os abastecimentos cheios (total - litrosNaoCheio)
			litrosCheios := consumo.TotalLitros - litrosNaoCheio
			consumo.MediaConsumo = float64(consumo.TotalKm) / litrosCheios
		}

		consumos = append(consumos, consumo)
	}

	return consumos, nil
}
