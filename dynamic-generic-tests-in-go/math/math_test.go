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
