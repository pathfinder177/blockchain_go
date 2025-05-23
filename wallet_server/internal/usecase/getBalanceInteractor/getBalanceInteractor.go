package GetBalanceInteractor

import (
	"context"
	"wallet_server/internal/entity"
	"wallet_server/internal/gateway"
)

type UseCase struct {
	gateway gateway.CliGateway
}

func New(g gateway.CliGateway) *UseCase {
	return &UseCase{
		gateway: g,
	}
}

func (uc *UseCase) GetBalance(ctx context.Context, e entity.Wallet) (string, error) {
	return uc.gateway.GetBalance(ctx, e)
}
