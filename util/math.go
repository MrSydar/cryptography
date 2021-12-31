package util

import (
	"fmt"
	"math/big"
	"math/rand"
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

/*
func isPrimitiveRoot(g, p int64) bool {
	G := big.NewInt(g)
	P := big.NewInt(p)
	P1 := big.NewInt(p - 1)
	return new(big.Int).Exp(G, P1, P).Int64() == int64(1)
}
*/

func factors(m int) []int {
	var result []int
	for i := 2; int64(i)*int64(i) <= int64(m); i++ {
		if m%i == 0 {
			result = append(result, i)
		}
		for m%i == 0 {
			m /= i
		}
	}
	if m > 1 {
		result = append(result, m)
	}
	return result
}

func safeMod(x, m int64) int64 {
	x %= m
	if x < 0 {
		x += m
	}
	return x
}

func powMod(x, n int64, m int) int64 {
	if m == 1 {
		return 0
	}
	um := uint(m)
	r := uint64(1)
	y := uint64(safeMod(x, int64(m)))

	for n > 0 {
		if n&1 > 0 {
			r = (r * y) % uint64(um)
		}
		y = (y * y) % uint64(um)
		n >>= 1
	}
	return int64(r)
}

func isPrimitiveRoot(m, g int) bool {
	if !(1 <= g && g < m) {
		return false
	}
	for _, x := range factors(m - 1) {
		if powMod(int64(g), int64((m-1)/x), m) == 1 {
			return false
		}
	}
	return true
}

func RandomPrimitiveRoot(num int64, variety int) (int64, error) {
	primitiveRoots := make([]int, variety)

	latestIndex := 0
	primitiveRoot := int64(1)
	for latestIndex < variety && primitiveRoot < num {
		if isPrimitiveRoot(int(primitiveRoot), int(num)) {
			primitiveRoots[latestIndex] = int(primitiveRoot)
			latestIndex++
			fmt.Println("primitive root: ", primitiveRoot)
		}
		primitiveRoot++
	}

	if latestIndex == 0 {
		return 0, fmt.Errorf("no primitive roots for %v", num)
	} else if latestIndex != variety-1 {
		return 0, fmt.Errorf("there are not enough primitive roots for given variety %v", variety)
	} else {
		return int64(PickRandomN(primitiveRoots)), nil
	}
}
