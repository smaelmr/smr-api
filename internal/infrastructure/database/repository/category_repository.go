package repository

import (
	"database/sql"

	"github.com/smaelmr/finance-api/internal/domain/entities"
)

type CategoryRepository struct {
	conn *sql.DB
}

func newCategoryRepository(conn *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		conn: conn,
	}
}

func (r *CategoryRepository) Add(category entities.Category) error {
	query := `INSERT INTO categoria (descricao, tipo) VALUES (?, ?)`
	_, err := r.conn.Exec(query, category.Description, category.Type)
	return err
}

func (r *CategoryRepository) GetAll() ([]entities.Category, error) {
	query := `SELECT id, descricao, tipo FROM categoria ORDER BY descricao`
	rows, err := r.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []entities.Category
	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(&category.Id, &category.Description, &category.Type); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepository) GetByType(categoryType string) ([]entities.Category, error) {
	query := `SELECT id, descricao, tipo FROM categoria WHERE tipo = ? ORDER BY descricao`
	rows, err := r.conn.Query(query, categoryType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []entities.Category
	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(&category.Id, &category.Description, &category.Type); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepository) Delete(id int64) error {
	query := `DELETE FROM categoria WHERE id = ?`
	_, err := r.conn.Exec(query, id)
	return err
}

func (r *CategoryRepository) Get(id int64) (*entities.Category, error) {
	query := `SELECT id, descricao, tipo FROM categoria WHERE id = ?`
	row := r.conn.QueryRow(query, id)

	var category entities.Category
	if err := row.Scan(&category.Id, &category.Description, &category.Type); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &category, nil
}
