package main

import "fmt"

func ExFunctional() {
	fmt.Println("ex_functional")

	v := []int{1, 2, 3, 4, 5}
	tmp := make([]int, 0, len(v))
	for _, x := range v {
		tmp = append(tmp, x*2)
	}

	f := make([]int, 0, len(tmp))
	for _, x := range tmp {
		if x%3 == 0 {
			f = append(f, x)
		}
	}

	sum := 0
	for _, x := range f {
		sum += x
	}
	fmt.Println("sum:", sum)

	a := []string{"a", "b", "c"}
	j := ""
	for i, s := range a {
		if i > 0 {
			j += "-"
		}
		j += s
	}
	fmt.Println("joined:", j)

	f1 := func(x int) int { return x + 1 }
	y := f1(9)
	fmt.Println("y:", y)
}
