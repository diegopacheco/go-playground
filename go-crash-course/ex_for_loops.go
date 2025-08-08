package main

import "fmt"

func ExForLoops() {
	fmt.Println("ex_for_loops")

	sum := 0
	for i := 0; i < 5; i++ {
		sum += i
	}
	fmt.Println("sum 0..4:", sum)

	v := []int{1, 2, 3}
	for _, x := range v {
		fmt.Println(x)
	}

	for i := len(v) - 1; i >= 0; i-- {
		fmt.Println(v[i])
	}
}
