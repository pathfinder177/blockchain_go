package SendCurrencyInteractor

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

func (uc *UseCase) SendCurrency(ctx context.Context, e entity.Wallet, amount, currency, receiver, mine string) (string, error) {
	return uc.gateway.SendCurrency(ctx, e, amount, currency, receiver, mine)
}
