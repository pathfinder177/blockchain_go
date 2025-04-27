package gateway

import (
	"context"
	"wallet_server/internal/entity"
)

type (
	CliGateway interface {
		GetBalance(context.Context, entity.Wallet) (string, error)
	}
	TCPGateway interface {
		GetHistory(context.Context, entity.Wallet) (string, error)
	}
)
