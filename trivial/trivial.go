package trivial

import (
	"errors"
	"fmt"
	"math/rand"
)

func checkPredefinedInputParams(s []int, k int) error {
	if k <= 1 {
		return errors.New("the k parameter can't equals or less than 1")
	} else if len(s) == 0 {
		return errors.New("n can't be zero")
	} else {
		for i := 0; i < len(s); i++ {
			if s[i] >= k {
				return fmt.Errorf("element on index %v can't be greater or equal to the k value", i)
			}
		}
		return nil
	}
}

func GetPartsRandom(s int, n int, k int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("n can't be negative")
	}

	arr := make([]int, n)

	arr[0] = s
	for i := 1; i < n; i++ {
		arr[i] = rand.Intn(k)
	}

	fmt.Printf("Generated %v\n", arr)

	return GetParts(arr, k)
}

func GetParts(s []int, k int) ([]int, error) {
	if iErr := checkPredefinedInputParams(s, k); iErr != nil {
		return nil, iErr
	}

	diff := s[0]
	for _, val := range s[1:] {
		diff -= val
	}

	lastPart := ((diff % k) + k) % k

	return append(s[1:], lastPart), nil
}

func GetSecret(s []int, k int) (int, error) {
	if iErr := checkPredefinedInputParams(s, k); iErr != nil {
		return -1, iErr
	}

	sum := s[0]
	for _, val := range s[1:] {
		sum += val
	}

	return ((sum % k) + k) % k, nil
}
