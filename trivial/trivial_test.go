package trivial

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDemoTrivial(t *testing.T) {
	n := 3
	k := 1000
	s := 392

	parts, _ := GetPartsRandom(s, n, k)

	fmt.Printf("\033[34mgenerated parts: %v\033[0m\n\n", parts)

	secret, _ := GetSecret(parts, k)

	fmt.Printf("\033[32msecret equals: %v\n\n\033[0m", secret)
}

func TestGenerateTrivialPartsForKEq1000AndNEq3AndSRandom(t *testing.T) {
	n := 3
	k := 1000
	s := 393

	actual, err := GetPartsRandom(s, n, k)
	if err != nil {
		t.Fatalf("error was not expected, but got %v", err)
	}

	if len(actual) != n {
		t.Fatalf("expected parts length to be %v, but got %v", n, len(actual))
	}

	for _, val := range actual {
		if val >= k {
			t.Fatalf("expected all values to be less than k, but got %v", val)
		}
	}
}

func TestGenerateTrivialPartsForNegativeN(t *testing.T) {
	n := -3
	k := 1000

	actual, err := GetPartsRandom(100, n, k)
	if err == nil {
		t.Fatalf("expected error")
	}

	if actual != nil {
		t.Fatalf("expected parts to be nil, but got %v", actual)
	}
}

func TestGenerateTrivialPartsForKEq1000AndNEq3AndSEq456AndS1Eq856AndS2Eq231(t *testing.T) {
	arr := []int{456, 856, 231}
	k := 1000

	actual, err := GetParts(arr, k)
	if expected := []int{856, 231, 369}; !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected parts %v, but got %v", expected, actual)
	}

	if err != nil {
		t.Fatalf("error was not expected, but got %v", err)
	}
}

func TestGenerateTrivialPartsEmptyParts(t *testing.T) {
	arr := []int{}
	k := 1000

	actual, err := GetParts(arr, k)

	if actual != nil {
		t.Fatalf("expected parts to be nil, but got %v", actual)
	}

	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestGenerateTrivialPartsOneElementEqualsK(t *testing.T) {
	arr := []int{456, 1000, 231}
	k := 1000

	actual, err := GetParts(arr, k)

	if actual != nil {
		t.Fatalf("expected parts to be nil, but got %v", actual)
	}

	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestGenerateTrivialPartsOneElementGreaterThanK(t *testing.T) {
	arr := []int{456, 1005, 231}
	k := 1000

	actual, err := GetParts(arr, k)

	if actual != nil {
		t.Fatalf("expected parts to be nil, but got %v", actual)
	}

	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestGenerateTrivialPartsKEq1(t *testing.T) {
	arr := []int{456, 856, 231}
	k := 1

	actual, err := GetParts(arr, k)

	if actual != nil {
		t.Fatalf("expected parts to be nil, but got %v", actual)
	}

	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestGenerateTrivialPartsKEqMinus1(t *testing.T) {
	arr := []int{456, 856, 231}
	k := -1

	actual, err := GetParts(arr, k)

	if actual != nil {
		t.Fatalf("expected parts to be nil, but got %v", actual)
	}

	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestGenerateTrivialSecretForKEq1000AndNEq3AndS1Eq856AndS2Eq231AndS3Eq369(t *testing.T) {
	parts := []int{856, 231, 369}
	k := 1000

	actual, err := GetSecret(parts, k)
	if expected := 456; actual != expected {
		t.Fatalf("expected secret %d, but got %d", expected, actual)
	}

	if err != nil {
		t.Fatalf("error was not expected, but got %v", err)
	}
}

func TestGenerateTrivialSecretEmptyParts(t *testing.T) {
	parts := []int{}
	k := 1000

	actual, err := GetSecret(parts, k)

	if actual != -1 {
		t.Fatalf("expected secret to be -1, but got %v", actual)
	}

	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestGenerateTrivialSecretOneElementEqualsK(t *testing.T) {
	parts := []int{856, 1000, 269}
	k := 1000

	actual, err := GetSecret(parts, k)

	if actual != -1 {
		t.Fatalf("expected secret to be -1, but got %v", actual)
	}

	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestGenerateTrivialSecretOneElementGreaterThanK(t *testing.T) {
	parts := []int{856, 1005, 269}
	k := 1000

	actual, err := GetSecret(parts, k)

	if actual != -1 {
		t.Fatalf("expected secret to be -1, but got %v", actual)
	}

	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestGenerateTrivialSecretKEq1(t *testing.T) {
	parts := []int{856, 231, 269}
	k := 1

	actual, err := GetSecret(parts, k)

	if actual != -1 {
		t.Fatalf("expected secret to be -1, but got %v", actual)
	}

	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestGenerateTrivialSecretKEqMinusOne(t *testing.T) {
	parts := []int{856, 231, 269}
	k := -1

	actual, err := GetSecret(parts, k)

	if actual != -1 {
		t.Fatalf("expected secret to be -1, but got %v", actual)
	}

	if err == nil {
		t.Fatalf("expected error")
	}
}
