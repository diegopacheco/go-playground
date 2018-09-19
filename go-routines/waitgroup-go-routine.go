package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	func() {
		count("cats")
		wg.Done()
	}()

	wg.Wait()
}

func count(msg string) {
	for i := 1; i <= 6; i++ {
		fmt.Println(i, msg)
		time.Sleep(time.Millisecond * 500)
	}

}
