package main

import (
	"errors"
	"fmt"
)

func main() {

	var ErrNotFound = errors.New("not found")
	var e = fmt.Errorf("%q: %w", "I could not foudn the error - LOL", ErrNotFound)
	if errors.Is(e, ErrNotFound) {
		println("e is ErrNotFound")
		println(e)
		println(e.Error())
	}

	type MyError struct {
		Env   string
		Blame string
		Err   error
	}
	var e2 = MyError{Env: "production", Blame: "No One", Err: errors.New("MyError jush happend")}
	println("Env: " + e2.Env + " blame: " + e2.Blame + " err: " + e2.Err.Error())
}
