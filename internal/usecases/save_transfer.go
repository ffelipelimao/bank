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
	uow                UowI
}

func NewSaveTransfer(transferRepository TransferRepository, balanceRepository BalanceRepository, uow UowI) *SaveTransfer {
	return &SaveTransfer{
		transferRepository: transferRepository,
		balanceRepository:  balanceRepository,
		uow:                uow,
	}
}

func (s *SaveTransfer) Execute(ctx context.Context, transfer *entities.Transfer) (*entities.Balance, error) {
	err := transfer.Validate()
	if err != nil {
		return nil, err
	}

	tx, err := s.uow.Begin()
	if err != nil {
		return nil, err
	}

	balance, err := s.balanceRepository.Get(ctx, tx, transfer.UserID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var funds int64
	if transfer.IsDebit() {
		if balance.Value-transfer.Value < -balance.Limit {
			tx.Rollback()
			return nil, ErrInsufficientFunds
		}
		funds = balance.Value - transfer.Value
	} else {
		funds = balance.Value + transfer.Value
	}

	err = s.transferRepository.Save(ctx, tx, transfer, transfer.UserID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = s.balanceRepository.Update(ctx, tx, funds, transfer.UserID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &entities.Balance{Value: funds, Limit: balance.Limit}, nil
}
