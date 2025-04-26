package GetBalanceInteractor

import (
	"context"
	"wallet_server/internal/entity"
	"wallet_server/internal/repo"
)

type UseCase struct {
	repo repo.CliRepo
}

func New(r repo.CliRepo) *UseCase {
	return &UseCase{
		repo: r,
	}
}

func (uc *UseCase) GetBalance(ctx context.Context, e entity.Wallet) (string, error) {
	return uc.repo.GetBalance(ctx, e)
}
