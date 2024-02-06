package usecases

import (
	"context"
	"time"

	"github.com/ffelipelimao/bank/entities"
)

type GetExtract struct {
	transferRepository TransferRepository
	balanceRepository  BalanceRepository
}

func NewGetExtract(transferRepository TransferRepository, balanceRepository BalanceRepository) *GetExtract {
	return &GetExtract{
		transferRepository: transferRepository,
		balanceRepository:  balanceRepository,
	}
}

func (g *GetExtract) Execute(ctx context.Context, userID *int64) (*entities.Extract, error) {
	balance, err := g.balanceRepository.Get(ctx, *userID)
	if err != nil {
		return nil, err
	}

	transfers, err := g.transferRepository.List(ctx, *userID)
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
