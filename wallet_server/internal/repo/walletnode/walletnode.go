package walletnode

type WalletNodeRepo struct {
	wNAddress string
	wNPort    string
}

func New(address, port string) *WalletNodeRepo {
	return &WalletNodeRepo{
		wNAddress: address,
		wNPort:    port,
	}
}

// func (wNode *WalletNodeRepo) GetHistory(ctx context.Context, w entity.Wallet) (string, error) {

// }
