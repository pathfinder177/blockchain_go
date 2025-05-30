package cli

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"wallet_server/internal/entity"
	gw "wallet_server/internal/gateway"
)

func (*cliGateway) GetBalance(ctx context.Context, w entity.Wallet) (string, error) {
	address := w.Address

	wb, err := _getBalance(address)
	if err != nil {
		return "", err
	}
	if wb == "" {
		return "", errors.New("wallet balance is empty")
	}

	return wb, nil
}

func _getBalance(address string) (string, error) {
	args := []string{"getbalance", "-address", address}
	cmd := exec.Command(gw.BlockchainBin, args...)
	cmd.Env = append(os.Environ(), gw.BlockchainWalletNodeID)
	cmd.Dir = gw.BlockchainBinPath

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func (*cliGateway) SendCurrency(ctx context.Context, e entity.Wallet, amount, currency, receiver, mine string) (string, error) {
	//FIXME validate args
	sender := e.Address

	args := []string{"send", "-from", sender, "-to", receiver, "-currency", currency, "-amount", amount, "-mine", mine}
	cmd := exec.Command(gw.BlockchainBin, args...)
	cmd.Env = append(os.Environ(), gw.BlockchainWalletNodeID)
	cmd.Dir = gw.BlockchainBinPath

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
