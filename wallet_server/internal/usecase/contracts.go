package usecase

import (
	"context"
	"wallet_server/internal/entity"
)

type (
	GetBalanceInteractor interface {
		GetBalance(context.Context, entity.Wallet) (string, error)
	}
	GetTransactionsHistoryInteractor interface {
		GetHistory(context.Context, entity.Wallet) (string, error)
	}
	SendCurrencyInteractor interface {
		SendCurrency(context.Context, entity.Wallet, string, string, string, string) (string, error)
	}
)
