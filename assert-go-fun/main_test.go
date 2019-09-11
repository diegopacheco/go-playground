package main

import (
	"testing"

	"github.com/deliveroo/assert-go"
)

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func Test(t *testing.T) {
	message := "bar"
	assert.Equal(t, message, "bar")

	p := Person{Name: "Alice"}
	assert.Equal(t, p, Person{Name: "Alice"})
}
