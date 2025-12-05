package repository

import (
	"database/sql"

	"github.com/smaelmr/finance-api/internal/domain/entities"
)

type PaymentMethodRepository struct {
	conn *sql.DB
}

func newPaymentMethodRepository(conn *sql.DB) *PaymentMethodRepository {
	return &PaymentMethodRepository{
		conn: conn,
	}
}

func (r *PaymentMethodRepository) Add(paymentMethod entities.PaymentMethod) error {
	query := `INSERT INTO forma_pagamento (descricao) VALUES (?)`
	_, err := r.conn.Exec(query, paymentMethod.Name)
	return err
}

func (r *PaymentMethodRepository) GetAll() ([]entities.PaymentMethod, error) {
	query := `SELECT id, descricao FROM forma_pagamento ORDER BY descricao`
	rows, err := r.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var paymentMethods []entities.PaymentMethod
	for rows.Next() {
		var paymentMethod entities.PaymentMethod
		if err := rows.Scan(&paymentMethod.Id, &paymentMethod.Name); err != nil {
			return nil, err
		}
		paymentMethods = append(paymentMethods, paymentMethod)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return paymentMethods, nil
}

func (r *PaymentMethodRepository) Get(id int64) (*entities.PaymentMethod, error) {
	query := `SELECT id, descricao FROM forma_pagamento WHERE id = ?`
	row := r.conn.QueryRow(query, id)

	var paymentMethod entities.PaymentMethod
	if err := row.Scan(&paymentMethod.Id, &paymentMethod.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &paymentMethod, nil
}

func (r *PaymentMethodRepository) Update(paymentMethod entities.PaymentMethod) error {
	query := `UPDATE forma_pagamento SET descricao = ? WHERE id = ?`
	result, err := r.conn.Exec(query, paymentMethod.Name, paymentMethod.Id)
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

func (r *PaymentMethodRepository) Delete(id int64) error {
	query := `DELETE FROM forma_pagamento WHERE id = ?`
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
