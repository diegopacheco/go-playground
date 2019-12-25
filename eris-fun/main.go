package main

import (
	"fmt"
	"github.com/rotisserie/eris"
)

func main() {
	var err = eris.New("not found")
	fmt.Println(err)
	if err != nil {
		err = eris.Wrapf(err, "error getting resource '%v'", 35)
	}
	fmt.Printf("%v", err)  // print without the stack trace
	fmt.Printf("%+v", err) // print with the stack trace
}
