package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ffelipelimao/bank/internal/entities"
	_ "github.com/go-sql-driver/mysql"
)

type BalanceRepository struct {
	DB *sql.DB
}

func NewBalanceRepository(db *sql.DB) *BalanceRepository {
	return &BalanceRepository{DB: db}
}

func (br *BalanceRepository) Get(ctx context.Context, userID int64) (*entities.Balance, error) {
	query := "SELECT limite, saldo FROM cliente_saldo WHERE id = ?"
	row := br.DB.QueryRowContext(ctx, query, userID)

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

func (br *BalanceRepository) Update(ctx context.Context, value int64, userID int64) error {
	query := "UPDATE cliente_saldo SET saldo = ? WHERE id = ?"
	_, err := br.DB.ExecContext(ctx, query, value, userID)
	if err != nil {
		return fmt.Errorf("failed to update balance on db: %v", err)
	}
	return nil
}
