package usecase

import (
	"context"
	"wallet_server/internal/entity"
)

type (
	GetBalanceInteractor interface {
		GetBalance(context.Context, entity.Wallet) (string, error)
	}
)
