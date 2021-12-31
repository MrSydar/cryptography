package distribution

import (
	"fmt"
	"math/big"
	"mrsydar/cryptography/util"
	"testing"
)

func TestDH(t *testing.T) {
	util.RefreshRandomSeed()

	n := big.NewInt(util.RandomPrimeNumber(20))
	g, err := util.RandomPrimitiveRoot(n.Int64())
	if err != nil {
		t.Fatal()
	}

	dh1, err := NewDiffieHellman(*n, *big.NewInt(g), 5)
	if err != nil {
		t.Fatalf("no error expected, but got %q", err)
	}

	dh2, err := NewDiffieHellman(*n, *big.NewInt(g), 5)
	if err != nil {
		t.Fatalf("no error expected, but got %v", err)
	}

	dh1.SetSessionKey(dh2.PublicKey)
	dh2.SetSessionKey(dh1.PublicKey)

	if dh1.SessionKey.Int64() != dh2.SessionKey.Int64() {
		t.Fatalf("similar session key expected, but got %v and %v", dh1.SessionKey.Int64(), dh2.SessionKey.Int64())
	}

	fmt.Println(dh1.SessionKey)
}
