package repository

import "database/sql"

// DataManager holds the methods that manipulates the main data.
type DataManager interface {
	RepoManager
	Begin() (TransactionManager, error)
	Close() error
}

// TransactionManager holds the methods that manipulates the main
// data, from within a transaction.
type TransactionManager interface {
	RepoManager
	Rollback() error
	Commit() error
	GetDBTransaction() *sql.Tx
}
