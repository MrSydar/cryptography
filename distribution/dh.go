package distribution

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

func NewDiffieHellman(primeNumberVariety, primitiveRootVariety int) (*DiffieHellman, error) {
	util.RefreshRandomSeed()

	dh := DiffieHellman{}

	dh.n = big.NewInt(util.RandomPrimeNumber(primeNumberVariety))

	if g, err := util.RandomPrimitiveRoot(dh.n.Int64(), primitiveRootVariety); err != nil {
		dh.g = big.NewInt(g)
	} else {
		return nil, err
	}

	dh.privateKey = big.NewInt(util.GetRandomN64(1, dh.n.Int64()-int64(1)))
	dh.PublicKey = new(big.Int).Exp(dh.g, dh.privateKey, dh.n)

	// fmt.Println(dh.n, dh.g, dh.privateKey.Int64(), dh.PublicKey.Int64())

	return &dh, nil
}
