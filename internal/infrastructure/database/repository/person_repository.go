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

func (r *PersonRepository) addPerson(person entities.Person) (int64, error) {
	queryPessoa := `INSERT INTO pessoa (
	       nome, cpf_cnpj, nome_contato, telefone, cep, cidade, estado, rua, numero, bairro
       ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
	result, err := r.conn.Exec(queryPessoa,
		person.Name,
		person.Document,
		person.Contact,
		person.PhoneNumber,
		person.Cep,
		person.City,
		person.State,
		person.Street,
		person.StreetNumber,
		person.Neighborhood,
	)
	if err != nil {
		return 0, err
	}

	// Obter o ID gerado
	pessoaID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return pessoaID, nil
}

func (r *PersonRepository) AddClient(person entities.Person) error {
	// Insere na tabela pessoa com todos os campos
	pessoaID, err := r.addPerson(person)
	if err != nil {
		return err
	}

	query := `INSERT INTO cliente (pessoa_id) VALUES (?);`
	_, err = r.conn.Exec(query, pessoaID)
	if err != nil {
		return err
	}

	return nil
}

func (r *PersonRepository) AddSupplier(person entities.Person) error {
	// Insere na tabela pessoa com todos os campos
	pessoaID, err := r.addPerson(person)
	if err != nil {
		return err
	}

	query := `INSERT INTO fornecedor (pessoa_id) VALUES (?);`
	_, err = r.conn.Exec(query, pessoaID)
	if err != nil {
		return err
	}

	return nil
}

func (r *PersonRepository) AddDriver(driver entities.Driver) error {
	tx, err := r.conn.Begin()
	if err != nil {
		return err
	}

	// Insere na tabela pessoa com todos os campos
	pessoaID, err := r.addPerson(driver.Person)
	if err != nil {
		tx.Rollback()
		return err
	}

	query := `INSERT INTO motorista (pessoa_id, cnh_categoria, cnh_validade, observacao) VALUES (?, ?, ?, ?);`
	_, err = tx.Exec(query, pessoaID, driver.CnhCategory, driver.CnhValidity, driver.Observation)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *PersonRepository) AddGasStation(person entities.Person) error {
	// Insere na tabela pessoa com todos os campos
	pessoaID, err := r.addPerson(person)
	if err != nil {
		return err
	}

	query := `INSERT INTO posto (pessoa_id) VALUES (?);`
	_, err = r.conn.Exec(query, pessoaID)
	if err != nil {
		return err
	}

	return nil
}

func (r *PersonRepository) GetClients() (retVal []entities.Person, err error) {
	query := `SELECT m.id, p.nome, p.cpf_cnpj, p.nome_contato, p.telefone, p.cep, 
		p.cidade, p.estado, p.rua, p.numero, p.bairro 
		FROM pessoa p 
		INNER JOIN cliente m ON p.id = m.pessoa_id ORDER BY p.nome`

	rows, err := r.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var client entities.Person
		var document, contact, phone, cep, city, state, street, number, neighborhood sql.NullString

		err = rows.Scan(
			&client.Id,
			&client.Name,
			&document,
			&contact,
			&phone,
			&cep,
			&city,
			&state,
			&street,
			&number,
			&neighborhood,
		)
		if err != nil {
			return nil, err
		}

		// Converter os campos NULL para string vazia
		if document.Valid {
			client.Document = document.String
		}
		if contact.Valid {
			client.Contact = contact.String
		}
		if phone.Valid {
			client.PhoneNumber = phone.String
		}
		if cep.Valid {
			client.Cep = cep.String
		}
		if city.Valid {
			client.City = city.String
		}
		if state.Valid {
			client.State = state.String
		}
		if street.Valid {
			client.Street = street.String
		}
		if number.Valid {
			client.StreetNumber = number.String
		}
		if neighborhood.Valid {
			client.Neighborhood = neighborhood.String
		}

		retVal = append(retVal, client)
	}

	return retVal, nil
}

func (r *PersonRepository) GetSuppliers() (retVal []entities.Person, err error) {
	query := `SELECT f.id, p.nome, p.cpf_cnpj, p.nome_contato, p.telefone, p.cep, 
		p.cidade, p.estado, p.rua, p.numero, p.bairro 
		FROM pessoa p 
		INNER JOIN fornecedor f ON p.id = f.pessoa_id ORDER BY p.nome`

	rows, err := r.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var supplier entities.Person
		var document, contact, phone, cep, city, state, street, number, neighborhood sql.NullString

		err = rows.Scan(
			&supplier.Id,
			&supplier.Name,
			&document,
			&contact,
			&phone,
			&cep,
			&city,
			&state,
			&street,
			&number,
			&neighborhood,
		)
		if err != nil {
			return nil, err
		}

		// Converter os campos NULL para string vazia
		if document.Valid {
			supplier.Document = document.String
		}
		if contact.Valid {
			supplier.Contact = contact.String
		}
		if phone.Valid {
			supplier.PhoneNumber = phone.String
		}
		if cep.Valid {
			supplier.Cep = cep.String
		}
		if city.Valid {
			supplier.City = city.String
		}
		if state.Valid {
			supplier.State = state.String
		}
		if street.Valid {
			supplier.Street = street.String
		}
		if number.Valid {
			supplier.StreetNumber = number.String
		}
		if neighborhood.Valid {
			supplier.Neighborhood = neighborhood.String
		}

		retVal = append(retVal, supplier)
	}

	return retVal, nil
}

func (r *PersonRepository) GetGasStationByCnpj(cnpj string) (entities.Person, error) {
	query := `
		SELECT f.id, p.nome
		FROM posto f 
		INNER JOIN pessoa p ON f.pessoa_id = p.id 
		WHERE p.cpf_cnpj = ?`

	row := r.conn.QueryRow(query, cnpj)

	supplier := entities.Person{}
	if err := row.Scan(
		&supplier.Id,
		&supplier.Name); err != nil {
		return entities.Person{}, err
	}

	return supplier, nil
}

func (r *PersonRepository) GetDrivers() (retVal []entities.Driver, err error) {
	query := `SELECT m.id, p.nome, p.cpf_cnpj, p.nome_contato, p.telefone, p.cep, 
		p.cidade, p.estado, p.rua, p.numero, p.bairro, m.cnh_categoria, m.cnh_validade, m.observacao
		FROM pessoa p 
		INNER JOIN motorista m ON p.id = m.pessoa_id 
		ORDER BY p.nome`

	rows, err := r.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var driver entities.Driver
		var document, contact, phone, cep, city, state, street, number, neighborhood, observation sql.NullString

		err = rows.Scan(
			&driver.Id,
			&driver.Name,
			&document,
			&contact,
			&phone,
			&cep,
			&city,
			&state,
			&street,
			&number,
			&neighborhood,
			&driver.CnhCategory,
			&driver.CnhValidity,
			&observation,
		)
		if err != nil {
			return nil, err
		}

		// Converter os campos NULL para string vazia
		if document.Valid {
			driver.Document = document.String
		}
		if contact.Valid {
			driver.Contact = contact.String
		}
		if phone.Valid {
			driver.PhoneNumber = phone.String
		}
		if cep.Valid {
			driver.Cep = cep.String
		}
		if city.Valid {
			driver.City = city.String
		}
		if state.Valid {
			driver.State = state.String
		}
		if street.Valid {
			driver.Street = street.String
		}
		if number.Valid {
			driver.StreetNumber = number.String
		}
		if neighborhood.Valid {
			driver.Neighborhood = neighborhood.String
		}
		if observation.Valid {
			driver.Observation = observation.String
		}

		retVal = append(retVal, driver)
	}

	return retVal, nil
}

func (r *PersonRepository) GetGasStations() (retVal []entities.Person, err error) {
	query := `SELECT m.id, p.nome, p.cpf_cnpj, p.nome_contato, p.telefone, p.cep, 
		p.cidade, p.estado, p.rua, p.numero, p.bairro 
		FROM pessoa p 
		INNER JOIN posto m ON p.id = m.pessoa_id ORDER BY p.nome`

	rows, err := r.conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var gasStation entities.Person
		var document, contact, phone, cep, city, state, street, number, neighborhood sql.NullString

		err = rows.Scan(
			&gasStation.Id,
			&gasStation.Name,
			&document,
			&contact,
			&phone,
			&cep,
			&city,
			&state,
			&street,
			&number,
			&neighborhood,
		)
		if err != nil {
			return nil, err
		}

		// Converter os campos NULL para string vazia
		if document.Valid {
			gasStation.Document = document.String
		}
		if contact.Valid {
			gasStation.Contact = contact.String
		}
		if phone.Valid {
			gasStation.PhoneNumber = phone.String
		}
		if cep.Valid {
			gasStation.Cep = cep.String
		}
		if city.Valid {
			gasStation.City = city.String
		}
		if state.Valid {
			gasStation.State = state.String
		}
		if street.Valid {
			gasStation.Street = street.String
		}
		if number.Valid {
			gasStation.StreetNumber = number.String
		}
		if neighborhood.Valid {
			gasStation.Neighborhood = neighborhood.String
		}

		retVal = append(retVal, gasStation)
	}

	return retVal, nil
}

func (r *PersonRepository) UpdateGasStation(person entities.Person) error {
	tx, err := r.conn.Begin()
	if err != nil {
		return err
	}

	// Primeiro atualiza a tabela pessoa
	queryPessoa := `UPDATE pessoa p 
		INNER JOIN posto m ON p.id = m.pessoa_id
		SET p.nome = ?, p.cpf_cnpj = ?, p.nome_contato = ?, p.telefone = ?, 
			p.cep = ?, p.cidade = ?, p.estado = ?, p.rua = ?, p.numero = ?, p.bairro = ?
		WHERE m.id = ?`

	_, err = tx.Exec(queryPessoa,
		person.Name,
		person.Document,
		person.Contact,
		person.PhoneNumber,
		person.Cep,
		person.City,
		person.State,
		person.Street,
		person.StreetNumber,
		person.Neighborhood,
		person.Id,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *PersonRepository) DeleteGasStation(id int64) error {
	tx, err := r.conn.Begin()
	if err != nil {
		return err
	}

	// Primeiro deleta da tabela posto
	queryPosto := `DELETE p FROM posto p WHERE p.id = ?`
	_, err = tx.Exec(queryPosto, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Depois deleta da tabela pessoa
	queryPessoa := `DELETE p FROM pessoa p 
		LEFT JOIN posto m ON p.id = m.pessoa_id
		WHERE m.id IS NULL`
	_, err = tx.Exec(queryPessoa)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *PersonRepository) UpdateClient(person entities.Person) error {
	tx, err := r.conn.Begin()
	if err != nil {
		return err
	}

	queryPessoa := `UPDATE pessoa p 
		INNER JOIN cliente c ON p.id = c.pessoa_id
		SET p.nome = ?, p.cpf_cnpj = ?, p.nome_contato = ?, p.telefone = ?, 
			p.cep = ?, p.cidade = ?, p.estado = ?, p.rua = ?, p.numero = ?, p.bairro = ?
		WHERE c.id = ?`

	_, err = tx.Exec(queryPessoa,
		person.Name,
		person.Document,
		person.Contact,
		person.PhoneNumber,
		person.Cep,
		person.City,
		person.State,
		person.Street,
		person.StreetNumber,
		person.Neighborhood,
		person.Id,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *PersonRepository) DeleteClient(id int64) error {
	tx, err := r.conn.Begin()
	if err != nil {
		return err
	}

	queryCliente := `DELETE c FROM cliente c WHERE c.id = ?`
	_, err = tx.Exec(queryCliente, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	queryPessoa := `DELETE p FROM pessoa p 
		LEFT JOIN cliente c ON p.id = c.pessoa_id
		WHERE c.id IS NULL`
	_, err = tx.Exec(queryPessoa)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *PersonRepository) UpdateSupplier(person entities.Person) error {
	tx, err := r.conn.Begin()
	if err != nil {
		return err
	}

	queryPessoa := `UPDATE pessoa p 
		INNER JOIN fornecedor f ON p.id = f.pessoa_id
		SET p.nome = ?, p.cpf_cnpj = ?, p.nome_contato = ?, p.telefone = ?, 
			p.cep = ?, p.cidade = ?, p.estado = ?, p.rua = ?, p.numero = ?, p.bairro = ?
		WHERE f.id = ?`

	_, err = tx.Exec(queryPessoa,
		person.Name,
		person.Document,
		person.Contact,
		person.PhoneNumber,
		person.Cep,
		person.City,
		person.State,
		person.Street,
		person.StreetNumber,
		person.Neighborhood,
		person.Id,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *PersonRepository) DeleteSupplier(id int64) error {
	tx, err := r.conn.Begin()
	if err != nil {
		return err
	}

	queryFornecedor := `DELETE f FROM fornecedor f WHERE f.id = ?`
	_, err = tx.Exec(queryFornecedor, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	queryPessoa := `DELETE p FROM pessoa p 
		LEFT JOIN fornecedor f ON p.id = f.pessoa_id
		WHERE f.id IS NULL`
	_, err = tx.Exec(queryPessoa)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *PersonRepository) UpdateDriver(driver entities.Driver) error {
	tx, err := r.conn.Begin()
	if err != nil {
		return err
	}

	// Atualiza a tabela pessoa
	queryPessoa := `UPDATE pessoa p 
		INNER JOIN motorista m ON p.id = m.pessoa_id
		SET p.nome = ?, p.cpf_cnpj = ?, p.nome_contato = ?, p.telefone = ?, 
			p.cep = ?, p.cidade = ?, p.estado = ?, p.rua = ?, p.numero = ?, p.bairro = ?
		WHERE m.id = ?`

	_, err = tx.Exec(queryPessoa,
		driver.Name,
		driver.Document,
		driver.Contact,
		driver.PhoneNumber,
		driver.Cep,
		driver.City,
		driver.State,
		driver.Street,
		driver.StreetNumber,
		driver.Neighborhood,
		driver.Id,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Atualiza a tabela motorista
	queryMotorista := `UPDATE motorista 
		SET cnh_categoria = ?, cnh_validade = ?, observacao = ?
		WHERE id = ?`

	_, err = tx.Exec(queryMotorista,
		driver.CnhCategory,
		driver.CnhValidity,
		driver.Observation,
		driver.Id,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *PersonRepository) DeleteDriver(id int64) error {
	tx, err := r.conn.Begin()
	if err != nil {
		return err
	}

	queryMotorista := `DELETE m FROM motorista m WHERE m.id = ?`
	_, err = tx.Exec(queryMotorista, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	queryPessoa := `DELETE p FROM pessoa p 
		LEFT JOIN motorista m ON p.id = m.pessoa_id
		WHERE m.id IS NULL`
	_, err = tx.Exec(queryPessoa)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
