package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/ffelipelimao/bank/internal/entities"
	"github.com/ffelipelimao/bank/internal/uow"
	_ "github.com/go-sql-driver/mysql"
)

type TransferRepository struct {
}

func NewTransferRepository() *TransferRepository {
	return &TransferRepository{}
}

func (tr *TransferRepository) Save(ctx context.Context, tx *uow.Transaction, transfer *entities.Transfer, userID int64) error {
	query := "INSERT INTO transacoes (cliente_id, valor,tipo, descricao, realizada_em) VALUES (?, ?, ?, ?, NOW())"
	_, err := tx.Tx.ExecContext(ctx, query, transfer.UserID, transfer.Value, transfer.Type, transfer.Description)
	if err != nil {
		return fmt.Errorf("failed to save on db: %v", err)
	}
	return nil
}

func (r *TransferRepository) List(ctx context.Context, tx *uow.Transaction, userID int64) ([]*entities.Transfer, error) {
	rows, err := tx.Tx.QueryContext(ctx, "SELECT valor, tipo, descricao, realizada_em FROM transacoes WHERE cliente_id = ? LIMIT 10", userID)
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
