package shamir

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDemoShamir(t *testing.T) {
	parts, _ := GetPartsRand(955, 4, 3, 1523)

	fmt.Printf("\033[34mgenerated parts: %v\n\n\033[0m", parts)

	custom_parts := []Part{
		parts[0],
		parts[1],
		parts[3],
	}
	secret, _ := GetSecret(custom_parts, 3, 1523)

	fmt.Printf("\033[32msecret equals: %v\n\n\033[0m", secret)
}

func TestGetParts(t *testing.T) {
	pN := 4
	pA := []int{954, 352, 62}
	pP := 1523

	actual, err := GetParts(pA, pN, pP)

	if err != nil {
		t.Fatalf("error is not expected: %v", err)
	}

	if len(actual) != pN {
		t.Fatalf("expected parts length to be %v, but got %v", pN, len(actual))
	}

	expected := []Part{
		{1368, 1},
		{383, 2},
		{1045, 3},
		{308, 4},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected parts to be %v, but got %v", expected, actual)
	}
}

func TestGetSecret(t *testing.T) {
	selectedParts := []Part{
		{1368, 1},
		{1045, 3},
		{308, 4},
	}

	actual, err := GetSecret(selectedParts, 3, 1523)

	if err != nil {
		t.Fatal("unexpected error")
	}

	expected := 954
	if actual != expected {
		t.Fatalf("expected secret to be %v, but got %v", expected, actual)
	}
}
