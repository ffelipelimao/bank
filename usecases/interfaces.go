package usecases

import (
	"context"

	"github.com/ffelipelimao/bank/entities"
)

type TransferRepository interface {
	Save(ctx context.Context, transfer *entities.Transfer, UserID int64) error
	List(ctx context.Context, userID int64) ([]*entities.Transfer, error)
}

type BalanceRepository interface {
	Get(ctx context.Context, userID int64) (*entities.Balance, error)
	Update(ctx context.Context, Value int64, userID int64) error
}
