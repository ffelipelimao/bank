package uow

import "database/sql"

type Transaction struct {
	Tx *sql.Tx
}

func (t *Transaction) Commit() error {
	return t.Tx.Commit()
}

func (t *Transaction) Rollback() error {
	return t.Tx.Rollback()
}
