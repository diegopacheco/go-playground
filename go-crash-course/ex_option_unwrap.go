package main

import "fmt"

func ExOptionUnwrap() {
	fmt.Println("ex_option_unwrap")

	var a *int
	v := 10
	a = &v
	var b *int

	if a != nil {
		fmt.Println("a:", *a)
	}

	y := 0
	if b != nil {
		y = *b
	}
	fmt.Println("b unwrap_or(0):", y)

	if a != nil {
		z := (*a) * 2
		fmt.Println("a mapped *2:", z)
	}
}
