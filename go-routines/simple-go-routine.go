package main

import (
	"fmt"
	"time"
)

func main() {
	go count("dogs")
	count("cats")
}

func count(msg string) {
	for i := 1; true; i++ {
		fmt.Println(i, msg)
		time.Sleep(time.Millisecond * 500)
	}

}
