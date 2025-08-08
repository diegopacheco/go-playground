package main

import (
	"fmt"
	"sync"
	"time"
)

func ExCoroutines() {
	fmt.Println("ex_coroutines")

	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		time.Sleep(50 * time.Millisecond)
		ch <- "ping"
	}()

	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		ch <- "pong"
	}()

	for i := 0; i < 2; i++ {
		fmt.Println(<-ch)
	}

	close(ch)
	wg.Wait()
}
