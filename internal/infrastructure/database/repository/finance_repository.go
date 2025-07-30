package repository

import (
	"database/sql"

	"github.com/smaelmr/finance-api/internal/domain/entities"
)

type FinanceRepository struct {
	conn *sql.DB
}

func NewFinanceRepository(conn *sql.DB) *FinanceRepository {
	return &FinanceRepository{
		conn: conn,
	}
}

func (r *FinanceRepository) CreateFinanceRecord(record entities.Finance) error {
	query := `INSERT INTO financeiro 
	(valor, descricao, forma_pagamento, numero_documento, observacoes, data_lancamento, data_vencimento, categoria_id, conta_id, centro_id, tags_id)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.conn.Exec(query,
		record.Amount,
		record.Description,
		record.PaymentMethod,
		record.DocumentNumber,
		record.Notes,
		record.LaunchDate,
		record.DueDate,
		record.AccountId,
		record.CostCenterId,
		record.TagsId)

	return err
}

func (r *FinanceRepository) GetFinanceRecord(id int) (*entities.Finance, error) {
	query := `SELECT * FROM registros_financeiros WHERE id = ?`
	row := r.conn.QueryRow(query, id)

	var record entities.Finance
	err := row.Scan(&record.Id,
		&record.LaunchDate,
		&record.DueDate,
		&record.Amount,
		&record.Description,
		&record.PaymentMethod,
		&record.DocumentNumber,
		&record.Notes,
		&record.CategoryId,
		&record.AccountId,
		&record.CostCenterId,
		&record.TagsId)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

// Atualizar e deletar podem ser adicionados aqui...
