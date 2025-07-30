package repository

import (
	"database/sql"

	"github.com/smaelmr/finance-api/internal/domain/entities"
)

type PersonRepository struct {
	conn *sql.DB
}

func newPersonRepository(conn *sql.DB) *PersonRepository {
	return &PersonRepository{
		conn: conn,
	}
}

func (r *PersonRepository) Add(person entities.Person) error {
	query :=
		`INSERT INTO pessoa
		(nome, cpf_cnpj)
		VALUES (?, ?);`

	_, err := r.conn.Exec(query,
		person.Name,
		person.CpfCnpj)

	return err
}

func (r *PersonRepository) GetCustomers() (retVal []entities.Person, err error) {
	query := `SELECT c.id, p.nome FROM pessoa p INNER JOIN cliente c ON p.id = c.pessoa_id ORDER BY p.nome`
	rows, err := r.conn.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		customer := entities.Person{}
		err = rows.Scan(
			&customer.Id,
			&customer.Name,
		)
		if err != nil {
			return nil, err
		}

		retVal = append(retVal, customer)
	}

	return retVal, nil
}

func (r *PersonRepository) GetSuppliers() (retVal []entities.Person, err error) {
	query := `SELECT f.id, p.nome FROM pessoa p INNER JOIN fornecedor f ON p.id = f.pessoa_id ORDER BY p.nome`
	rows, err := r.conn.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		customer := entities.Person{}
		err = rows.Scan(
			&customer.Id,
			&customer.Name,
		)
		if err != nil {
			return nil, err
		}

		retVal = append(retVal, customer)
	}

	return retVal, nil
}

func (r *PersonRepository) GetDrivers() (retVal []entities.Person, err error) {
	query := `SELECT m.id, p.nome FROM pessoa p INNER JOIN motorista m ON p.id = m.pessoa_id ORDER BY p.nome`
	rows, err := r.conn.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		driver := entities.Person{}
		err = rows.Scan(
			&driver.Id,
			&driver.Name,
		)
		if err != nil {
			return nil, err
		}

		retVal = append(retVal, driver)
	}

	return retVal, nil
}
