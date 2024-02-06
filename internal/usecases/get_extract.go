package usecases

import (
	"context"
	"time"

	"github.com/ffelipelimao/bank/internal/entities"
)

type GetExtract struct {
	transferRepository TransferRepository
	balanceRepository  BalanceRepository
	uow                UowI
}

func NewGetExtract(transferRepository TransferRepository, balanceRepository BalanceRepository, uow UowI) *GetExtract {
	return &GetExtract{
		transferRepository: transferRepository,
		balanceRepository:  balanceRepository,
		uow:                uow,
	}
}

func (g *GetExtract) Execute(ctx context.Context, userID *int64) (*entities.Extract, error) {
	tx, err := g.uow.Begin()
	if err != nil {
		return nil, err
	}

	balance, err := g.balanceRepository.Get(ctx, tx, *userID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	transfers, err := g.transferRepository.List(ctx, tx, *userID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &entities.Extract{
		Balance: entities.BalanceExtract{
			Limit: balance.Limit,
			Value: balance.Value,
			Date:  time.Now(),
		},
		Tranfers: transfers,
	}, nil

}
