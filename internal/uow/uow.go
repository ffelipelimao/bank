package uow

import "database/sql"

type UnitOfWork struct {
	DB *sql.DB
}

func NewUnityOfWork(DB *sql.DB) *UnitOfWork {
	return &UnitOfWork{
		DB: DB,
	}
}

func (u *UnitOfWork) Begin() (*Transaction, error) {
	tx, err := u.DB.Begin()
	if err != nil {
		return nil, err
	}
	return &Transaction{Tx: tx}, nil
}

type SQLTransaction struct {
	Tx *sql.Tx
}

func (t *SQLTransaction) Commit() error {
	return t.Tx.Commit()
}

func (t *SQLTransaction) Rollback() error {
	return t.Tx.Rollback()
}
