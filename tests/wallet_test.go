package tests

import (
	"blockchain/internal/app"
	"crypto/ecdsa"
	"crypto/elliptic"
	"math/big"
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
			want: "17v3P2zeB49coGX2Sz6byFptKJRaDVzC6Y",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := new(big.Int).SetBits([]big.Word{
				18302347381835011368,
				10824219518206851950,
				14543274358396171484,
				15200568927396967186,
			})

			// Construct Y using the provided words.
			y := new(big.Int).SetBits([]big.Word{
				2320667325194657460,
				12320845091154170490,
				1777185806402393421,
				12513258465373944226,
			})

			// Construct D using the provided words.
			d := new(big.Int).SetBits([]big.Word{
				14601373157214410493,
				721469850043518372,
				18338694530410490814,
				12521284765453296601,
			})

			// Create the private key for testing.
			privateKey := &ecdsa.PrivateKey{
				PublicKey: ecdsa.PublicKey{
					Curve: elliptic.P256(),
					X:     x,
					Y:     y,
				},
				D: d,
			}

			pubKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)

			w := app.Wallet{
				PrivateKey: *privateKey,
				PublicKey:  pubKey,
			}

			if got := string(w.GetAddress()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Wallet.GetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
