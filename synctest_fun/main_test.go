package main

import (
	"fmt"
	"sync"
	"testing"
	"testing/synctest"
	"time"
)

func demonstrateSynctest() {
	fmt.Println("\n=== Testing/Synctest Package ===")
	testFunc := func(t *testing.T) {
		synctest.Test(t, func(t *testing.T) {
			start := time.Now()
			var wg sync.WaitGroup
			wg.Add(2)

			go func() {
				defer wg.Done()
				time.Sleep(1 * time.Hour)
			}()

			go func() {
				defer wg.Done()
				time.Sleep(30 * time.Minute)
			}()

			wg.Wait()
			elapsed := time.Since(start)
			fmt.Printf("Elapsed time: %v\n", elapsed)
		})
	}
	fmt.Printf("Synctest demo function created: %T\n", testFunc)
}

func TestSynctest(t *testing.T) {
	demonstrateSynctest()

	realStart := time.Now()
	synctest.Test(t, func(t *testing.T) {
		virtualStart := time.Now()

		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Hour)
		}()

		go func() {
			defer wg.Done()
			time.Sleep(30 * time.Minute)
		}()

		wg.Wait()
		virtualElapsed := time.Since(virtualStart)
		fmt.Printf("Virtual elapsed time: %v\n", virtualElapsed)

		if virtualElapsed < 1*time.Hour {
			t.Errorf("Expected virtual time to be at least 1 hour, but got %v", virtualElapsed)
		}
	})

	realElapsed := time.Since(realStart)
	fmt.Printf("Real elapsed time: %v\n", realElapsed)
	if realElapsed > 10*time.Second {
		t.Errorf("Expected real test to complete quickly, but took %v", realElapsed)
	}
}
