package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count("cats", c)
	/* Receiver */
	for {
		msg, open := <-c
		fmt.Println(msg)
		if !open {
			break
		}
	}
}

/* Sender */
func count(msg string, c chan string) {
	for i := 1; i <= 6; i++ {
		c <- msg
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
