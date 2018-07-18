package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	// restore the echoing state when exiting
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()

	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		fmt.Println("I got the byte", b, "("+string(b)+")")
	}

}
