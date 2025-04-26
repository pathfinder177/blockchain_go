package GetTransactionsHistoryInteractor

import (
	"context"
	"wallet_server/internal/entity"
	"wallet_server/internal/repo"
)

type UseCase struct {
	repo repo.WalletNodeRepo
}

func New(r repo.WalletNodeRepo) *UseCase {
	return &UseCase{
		repo: r,
	}
}

func (uc *UseCase) GetBalance(ctx context.Context, e entity.Wallet) (string, error) {
	return uc.repo.GetHistory(ctx, e)
}
