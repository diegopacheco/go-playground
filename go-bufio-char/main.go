package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {

	// avoid buffer and avoid to print :(
	/*
		exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
		exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
		defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()
	*/
	Up := []byte{27, 91, 65}
	Down := []byte{27, 91, 66}
	Right := []byte{27, 91, 67}
	Left := []byte{27, 91, 68}

	r := bufio.NewReader(os.Stdin)
	strLine, _ := r.ReadString('\n')
	strLine = strings.Replace(strLine, "\n", "", 0)
	strLine = strings.Replace(strLine, "\r", "", 0)

	fmt.Printf("HERE Up Means: % x] \n", Up)
	fmt.Printf("Byte String % x] ", bytes.NewBufferString(strLine).Bytes()[0:3])

	if bytes.Equal(bytes.NewBufferString(strLine).Bytes()[0:3], Up) == true {
		fmt.Println("UP pressed")
	} else if bytes.Equal(bytes.NewBufferString(strLine).Bytes()[0:3], Down) == true {
		fmt.Println("Down pressed")
	} else if bytes.Equal(bytes.NewBufferString(strLine).Bytes()[0:3], Left) == true {
		fmt.Println("Left pressed")
	} else if bytes.Equal(bytes.NewBufferString(strLine).Bytes()[0:3], Right) == true {
		fmt.Println("Right pressed")
	} else {
		fmt.Print("Unkonw: ")
		fmt.Println(strLine)
	}
}
