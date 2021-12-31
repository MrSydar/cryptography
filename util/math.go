package util

import (
	"errors"
	"math/big"
	"math/rand"
	"reflect"
	"time"
)

func RefreshRandomSeed() {
	rand.Seed(time.Now().UnixNano())
}

func GetRandomN(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func GetRandomN64(min, max int64) int64 {
	return rand.Int63n(max-min+1) + min
}

func PickRandomN(arr []int) int {
	return arr[GetRandomN(0, len(arr)-1)]
}

func PickRandomN64(arr []int64) int64 {
	return arr[GetRandomN(0, len(arr)-1)]
}

func isPrime(n int64) bool {
	return big.NewInt(n).ProbablyPrime(0)
}

func RandomPrimeNumber(variety int) int64 {
	primes := make([]int64, variety)

	latestIndex := 0
	p := int64(499)
	for latestIndex < variety {
		if isPrime(p) {
			primes[latestIndex] = p
			latestIndex++
		}
		p++
	}

	return PickRandomN64(primes)
}

func RandomPrimitiveRoot(modulo int64) (int64, error) {
	gcd := func(a, b int64) int64 {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	coprimes := make(map[int64]bool)
	for num := int64(1); num < modulo; num++ {
		if gcd(num, modulo) == 1 {
			coprimes[num] = true
		}
	}

	primRoots := make([]int64, 0)

	for potentialPrimRoot := int64(1); potentialPrimRoot < modulo; potentialPrimRoot++ {
		potentialPrimRootSet := make(map[int64]bool)

		for powers := int64(1); powers < modulo; powers++ {
			x := new(big.Int).Exp(big.NewInt(potentialPrimRoot), big.NewInt(powers), big.NewInt(modulo))
			potentialPrimRootSet[x.Int64()] = true
		}

		if reflect.DeepEqual(coprimes, potentialPrimRootSet) {
			primRoots = append(primRoots, potentialPrimRoot)
		}
	}

	if len(primRoots) == 0 {
		return 0, errors.New("no primitive roots found")
	} else {
		return PickRandomN64(primRoots), nil
	}
}
