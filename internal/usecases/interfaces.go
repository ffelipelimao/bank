package usecases

import (
	"context"

	"github.com/ffelipelimao/bank/internal/entities"
	"github.com/ffelipelimao/bank/internal/uow"
)

type TransferRepository interface {
	Save(ctx context.Context, tx *uow.Transaction, transfer *entities.Transfer, UserID int64) error
	List(ctx context.Context, tx *uow.Transaction, userID int64) ([]*entities.Transfer, error)
}

type BalanceRepository interface {
	Get(ctx context.Context, tx *uow.Transaction, userID int64) (*entities.Balance, error)
	Update(ctx context.Context, tx *uow.Transaction, Value int64, userID int64) error
}

type UowI interface {
	Begin() (*uow.Transaction, error)
}
