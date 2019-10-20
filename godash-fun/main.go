package main

import (
	"fmt"

	"github.com/thecasualcoder/godash"
)

// Person is a simple pojo
type Person struct {
	Name string
	Age  int
}

func main() {
	input := []Person{
		{Name: "John", Age: 22},
		{Name: "Doe", Age: 23},
	}
	var output int
	godash.Reduce(input, &output, func(sum int, person Person) int {
		return sum + person.Age
	})
	fmt.Println(output) // prints 45

	input2 := []int{1, 2, 3, 4, 5}
	var output2 []int
	godash.Map(input2, &output2, func(el int) int {
		return el * el
	})
	fmt.Println(output2) // pint [1 4 9 16 25]
}
