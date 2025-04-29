package math

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		valA     int
		ValB     int
		Expected int64
	}{
		{1, 2, 3},
		{2, 3, 5},
		{-3, 4, 1},
		{0, 0, 0},
	}

	for _, test := range tests {
		operation := IntOperation{
			ValueA: test.valA,
			ValueB: test.ValB,
		}
		result := operation.Add()
		fmt.Printf("Add(%d, %d) = %d; expected %d\n", test.valA, test.ValB, result, test.Expected)
		if result != IntResult(test.Expected) {
			t.Errorf("Add(%d, %d) = %d; expected %d", test.valA, test.ValB, result, test.Expected)
		}
	}
}

func TestDivision(t *testing.T) {
	tests := []struct {
		valA     int
		ValB     int
		Expected int64
	}{
		{1, 2, 0},
		{2, 3, 0},
		{-3, 4, 0},
		{0, 0, 0},
	}

	for _, test := range tests {
		io := IntOperation{
			ValueA: test.valA,
			ValueB: test.ValB,
		}
		result := io.Divide()
		fmt.Printf("Divide(%d, %d) = %d; expected %d\n", test.valA, test.ValB, result, test.Expected)
		if result != IntResult(test.Expected) {
			t.Errorf("Divide(%d, %d) = %d; expected %d", test.valA, test.ValB, result, test.Expected)
		}
	}
}

func TestMultiplyAndSub(t *testing.T) {
	tests := []struct {
		valA      int
		ValB      int
		Operation string
		Expected  int64
	}{
		{1, 2, "*", 2},
		{2, 3, "*", 6},
		{-3, 4, "*", -12},
		{0, 0, "*", 0},
		{1, 2, "-", -1},
		{2, 3, "-", -1},
		{-3, 4, "-", -7},
		{0, 0, "-", 0},
	}

	for _, test := range tests {
		io := IntOperation{
			ValueA: test.valA,
			ValueB: test.ValB,
		}
		result := io.FromString(test.Operation)
		fmt.Printf("%s (%d, %d) = %d; expected %d\n", test.Operation, test.valA, test.ValB, result, test.Expected)
		if result != IntResult(test.Expected) {
			t.Errorf("%s (%d, %d) = %d; expected %d", test.Operation, test.valA, test.ValB, result, test.Expected)
		}
	}
}
