package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"math/big"
	"reflect"
	"testing"
)

func TestWallets_GetAddresses(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "cmp",
			want: []string{"17v3P2zeB49coGX2Sz6byFptKJRaDVzC6Y"},
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

			w_address := "17v3P2zeB49coGX2Sz6byFptKJRaDVzC6Y"
			w := &Wallet{
				PrivateKey: *privateKey,
				PublicKey:  append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...),
			}

			ws := Wallets{}
			ws.Wallets = make(map[string]*Wallet)
			ws.Wallets[w_address] = w

			if got := ws.GetAddresses(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Wallets.GetAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWallets_GetWallet(t *testing.T) {
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

	w := &Wallet{
		PrivateKey: *privateKey,
		PublicKey:  append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...),
	}
	w_address := string("17v3P2zeB49coGX2Sz6byFptKJRaDVzC6Y")

	ws := Wallets{}
	ws.Wallets = make(map[string]*Wallet)
	ws.Wallets[w_address] = w

	tests := []struct {
		name string
		want Wallet
	}{
		{
			name: "cmp",
			want: *w,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ws.GetWallet(w_address); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Wallets.GetWallet() = %v, want %v", got, tt.want)
			}
		})
	}
}
