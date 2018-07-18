package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()

	Up := []byte{27, 91, 65}

	r := bufio.NewReader(os.Stdin)
	strLine, _ := r.ReadString('\n')
	strLine = strings.Replace(strLine, "\n", "", 0)
	strLine = strings.Replace(strLine, "\r", "", 0)

	fmt.Printf("HERE Up Means: % x] \n", Up)
	fmt.Printf("Byte String % x] ", bytes.NewBufferString(strLine).Bytes()[0:3])
	fmt.Printf("Comp Up and Str: %t \n", bytes.Equal(bytes.NewBufferString(strLine).Bytes()[0:3], Up))

	if bytes.Equal(bytes.NewBufferString(strLine).Bytes()[0:3], Up) == true {
		fmt.Println("UP pressed")
	} else {
		fmt.Print("Unkonw: ")
		fmt.Println(strLine)
	}
}
