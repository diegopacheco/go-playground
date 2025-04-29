package main

import (
	"fmt"

	"github.com/diegopacheco/dynamic-generic-tests-in-go/math"
)

func main() {
	var io = math.IntOperation{
		ValueA: 10,
		ValueB: 5,
	}
	result := io.FromString("+")
	fmt.Println(result)
}
