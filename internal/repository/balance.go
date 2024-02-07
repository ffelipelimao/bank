package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ffelipelimao/bank/internal/entities"
	"github.com/ffelipelimao/bank/internal/uow"
	_ "github.com/go-sql-driver/mysql"
)

type BalanceRepository struct {
}

func NewBalanceRepository() *BalanceRepository {
	return &BalanceRepository{}
}

func (br *BalanceRepository) Get(ctx context.Context, tx *uow.Transaction, userID int64) (*entities.Balance, error) {
	query := "SELECT limite, saldo FROM cliente_saldo WHERE id = ?"
	row := tx.Tx.QueryRowContext(ctx, query, userID)

	balance := &entities.Balance{}
	err := row.Scan(&balance.Limit, &balance.Value)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user does not exists")
		}
		return nil, fmt.Errorf("failed to get balance on db: %v", err)
	}

	return balance, nil
}

func (br *BalanceRepository) Update(ctx context.Context, tx *uow.Transaction, value int64, userID int64) error {
	query := "UPDATE cliente_saldo SET saldo = ? WHERE id = ?"
	_, err := tx.Tx.ExecContext(ctx, query, value, userID)
	if err != nil {
		return fmt.Errorf("failed to update balance on db: %v", err)
	}
	return nil
}
