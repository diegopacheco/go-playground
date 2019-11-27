package main

import (
	"errors"
	"fmt"
)

func main() {

	var ErrNotFound = errors.New("not found")
	var e = fmt.Errorf("%q: %w", "I could not foudn the error - LOL", ErrNotFound)
	if e != nil {
		println("e is ErrNotFound")
		println(e)
		println(e.Error())
	}

}
