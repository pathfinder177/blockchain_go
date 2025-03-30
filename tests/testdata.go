package tests

import (
	"blockchain/internal/app"
	"crypto/ecdsa"
	"crypto/elliptic"
	"math/big"
)

var test_wallet_address string = "17v3P2zeB49coGX2Sz6byFptKJRaDVzC6Y"

var test_privkey_x = new(big.Int).SetBits([]big.Word{
	18302347381835011368,
	10824219518206851950,
	14543274358396171484,
	15200568927396967186,
})
var test_privkey_y = new(big.Int).SetBits([]big.Word{
	2320667325194657460,
	12320845091154170490,
	1777185806402393421,
	12513258465373944226,
})
var test_privkey_d = new(big.Int).SetBits([]big.Word{
	14601373157214410493,
	721469850043518372,
	18338694530410490814,
	12521284765453296601,
})

var test_privateKey *ecdsa.PrivateKey = &ecdsa.PrivateKey{
	PublicKey: ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     test_privkey_x,
		Y:     test_privkey_y,
	},
	D: test_privkey_d,
}

var test_wallet app.Wallet = app.Wallet{
	PrivateKey: *test_privateKey,
	PublicKey:  append(test_privateKey.PublicKey.X.Bytes(), test_privateKey.PublicKey.Y.Bytes()...),
}

var test_wallets = app.Wallets{Wallets: make(map[string]*app.Wallet)}
