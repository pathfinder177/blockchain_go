package repo

import (
	"context"
	"wallet_server/internal/entity"
)

type (
	CliRepo interface {
		GetBalance(context.Context, entity.Wallet) (string, error)
	}
	WalletNodeRepo interface {
		GetHistory(context.Context, entity.Wallet) (string, error)
	}
)
