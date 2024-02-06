package usecases

import (
	"context"
	"errors"

	"github.com/ffelipelimao/bank/internal/entities"
)

var ErrInsufficientFunds = errors.New("insufficient funds")

type SaveTransfer struct {
	transferRepository TransferRepository
	balanceRepository  BalanceRepository
}

func NewSaveTransfer(transferRepository TransferRepository, balanceRepository BalanceRepository) *SaveTransfer {
	return &SaveTransfer{
		transferRepository: transferRepository,
		balanceRepository:  balanceRepository,
	}
}

func (s *SaveTransfer) Execute(ctx context.Context, transfer *entities.Transfer) (*entities.Balance, error) {
	err := transfer.Validate()
	if err != nil {
		return nil, err
	}

	balance, err := s.balanceRepository.Get(ctx, transfer.UserID)
	if err != nil {
		return nil, err
	}

	var funds int64
	if transfer.IsDebit() {
		if balance.Value-transfer.Value < -balance.Limit {
			return nil, ErrInsufficientFunds
		}
		funds = balance.Value - transfer.Value
	} else {
		funds = balance.Value + transfer.Value
	}

	err = s.transferRepository.Save(ctx, transfer, transfer.UserID)
	if err != nil {
		return nil, err
	}

	err = s.balanceRepository.Update(ctx, funds, transfer.UserID)
	if err != nil {
		return nil, err
	}

	return &entities.Balance{Value: funds, Limit: balance.Limit}, nil
}
