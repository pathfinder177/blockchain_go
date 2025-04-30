package GetTransactionsHistoryInteractor

import (
	"context"
	"wallet_server/internal/entity"
	"wallet_server/internal/gateway"
)

type UseCase struct {
	gateway gateway.TCPGateway
}

func New(g gateway.TCPGateway) *UseCase {
	return &UseCase{
		gateway: g,
	}
}

func (uc *UseCase) GetHistory(ctx context.Context, e entity.Wallet) (string, error) {
	return uc.gateway.GetHistory(ctx, e)
}
