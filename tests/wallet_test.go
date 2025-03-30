package tests

import (
	"blockchain/internal/app"
	"crypto/ecdsa"
	"reflect"
	"testing"
)

func TestNewWalletFieldTypes(t *testing.T) {
	walletType := reflect.TypeOf(app.Wallet{})

	expectedFields := map[string]reflect.Type{
		"PrivateKey": reflect.TypeOf(ecdsa.PrivateKey{}),
		"PublicKey":  reflect.TypeOf([]byte{}),
	}

	for i := range walletType.NumField() {
		field := walletType.Field(i)
		expectedType, exists := expectedFields[field.Name]

		if !exists {
			t.Errorf("Unexpected field: %s", field.Name)
		} else if field.Type != expectedType {
			t.Errorf("Field %s has type %s, expected %s", field.Name, field.Type, expectedType)
		}
	}
}

func TestWallet_GetAddress(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "cmp",
			want: test_wallet_address,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := string(test_wallet.GetAddress()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Wallet.GetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
