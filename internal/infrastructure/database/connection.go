package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/smaelmr/finance-api/config"
)

type MySQLConnection struct {
	DB     *sql.DB
	config *config.Config
}

// NewMySQLConnection cria uma nova conexão com o banco de dados MySQL.
func NewMySQLConnection(config *config.Config) (*MySQLConnection, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.Database.User,
		config.Database.Pass,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// Testa a conexão
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &MySQLConnection{DB: db}, nil
}

// Close encerra a conexão com o banco de dados.
func (conn *MySQLConnection) Close() error {
	return conn.DB.Close()
}
