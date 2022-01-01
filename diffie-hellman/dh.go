package dh

import (
	"math/big"
	"mrsydar/cryptography/util"
)

type DiffieHellman struct {
	n, g, privateKey,
	SessionKey, PublicKey *big.Int
}

func (dh *DiffieHellman) SetSessionKey(publicKey *big.Int) {
	dh.SessionKey = new(big.Int).Exp(publicKey, dh.privateKey, dh.n)
}

func NewDiffieHellman(n, g big.Int, primeNumberVariety int) (*DiffieHellman, error) {
	dh := DiffieHellman{}
	dh.n = &n
	dh.g = &g

	dh.privateKey = big.NewInt(util.GetRandomN64(1, dh.n.Int64()-int64(1)))
	dh.PublicKey = new(big.Int).Exp(dh.g, dh.privateKey, dh.n)

	return &dh, nil
}
