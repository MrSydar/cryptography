package util

import (
	"errors"
	"fmt"
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

	var coprimes []int64
	for num := int64(1); num < modulo; num++ {
		if gcd(num, modulo) == 1 {
			coprimes = append(coprimes, num)
		}
	}

	primitiveRoots := make([]int64, 0)

	for g := int64(1); g < modulo; g++ {
		var actual_set []int64
		for powers := int64(1); powers < modulo; powers++ {
			x := new(big.Int).Exp(big.NewInt(g), big.NewInt(powers), big.NewInt(modulo))
			actual_set = append(actual_set, x.Int64())
		}

		if reflect.DeepEqual(coprimes, actual_set) {
			primitiveRoots = append(primitiveRoots, g)
		} else {
			fmt.Println("For primitive root", g, "set is", actual_set)
		}
	}

	fmt.Println("For modulo", modulo, "coprimes are", coprimes)

	if len(primitiveRoots) == 0 {
		return 0, errors.New("no primitive roots found")
	} else {
		return PickRandomN64(primitiveRoots), nil
	}

}
