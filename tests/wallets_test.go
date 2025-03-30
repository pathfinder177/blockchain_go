package tests

import (
	"blockchain/internal/app"
	"reflect"
	"testing"
)

func TestWallets_GetAddresses(t *testing.T) {
	test_wallets.Wallets[test_wallet_address] = &test_wallet

	tests := []struct {
		name string
		want []string
	}{
		{
			name: "cmp",
			want: []string{test_wallet_address},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := test_wallets.GetAddresses(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Wallets.GetAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWallets_GetWallet(t *testing.T) {
	test_wallets.Wallets[test_wallet_address] = &test_wallet

	tests := []struct {
		name string
		want app.Wallet
	}{
		{
			name: "cmp",
			want: test_wallet,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := test_wallets.GetWallet(test_wallet_address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Wallets.GetWallet() = %v, want %v", got, tt.want)
			}
		})
	}
}
