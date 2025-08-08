package main

import "fmt"

func ExIfElse() {
	fmt.Println("ex_if_else")

	x := 7
	if x%2 == 0 {
		fmt.Println("even")
	} else if x%3 == 0 {
		fmt.Println("divisible by three")
	} else {
		fmt.Println("odd")
	}
}
