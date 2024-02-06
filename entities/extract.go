package entities

import (
	"time"
)

type Extract struct {
	Balance  BalanceExtract `json:"saldo"`
	Tranfers []*Transfer    `json:"ultimas_transacoes"`
}

type BalanceExtract struct {
	Limit int64     `json:"limite"`
	Date  time.Time `json:"data_extrato"`
	Value int64     `json:"saldo"`
}
