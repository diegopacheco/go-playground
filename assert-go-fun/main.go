package main

// Person struct
type Person struct {
	Name string
}

// Sum does the basic math
func Sum(x int, y int) int {
	return x + y
}

func main() {
	Sum(5, 5)
}
