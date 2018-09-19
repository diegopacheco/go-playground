package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go count("cats", c1, 500)
	go count("dogs", c2, 2000)
	/* Receiver */
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

/* Sender */
func count(msg string, c chan string, t int) {
	for i := 1; i <= 6; i++ {
		c <- msg
		time.Sleep(time.Millisecond * time.Duration(t))
	}
	close(c)
}
