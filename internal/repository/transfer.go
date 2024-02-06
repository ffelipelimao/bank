package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ffelipelimao/bank/internal/entities"
	_ "github.com/go-sql-driver/mysql"
)

type TransferRepository struct {
	DB *sql.DB
}

func NewTransferRepository(db *sql.DB) *TransferRepository {
	return &TransferRepository{DB: db}
}

func (tr *TransferRepository) Save(ctx context.Context, transfer *entities.Transfer, userID int64) error {
	query := "INSERT INTO transacoes (cliente_id, valor,tipo, descricao, realizada_em) VALUES (?, ?, ?, ?, NOW())"
	_, err := tr.DB.ExecContext(ctx, query, transfer.UserID, transfer.Value, transfer.Type, transfer.Description)
	if err != nil {
		return fmt.Errorf("failed to save on db: %v", err)
	}
	return nil
}

func (r *TransferRepository) List(ctx context.Context, userID int64) ([]*entities.Transfer, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT valor, tipo, descricao, realizada_em FROM transacoes WHERE cliente_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transfers []*entities.Transfer

	for rows.Next() {
		var transfer entities.Transfer
		var createdAt time.Time

		err := rows.Scan(&transfer.Value, &transfer.Type, &transfer.Description, &createdAt)
		if err != nil {
			return nil, err
		}

		transfer.CreatedAt = createdAt
		transfers = append(transfers, &transfer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transfers, nil
}
