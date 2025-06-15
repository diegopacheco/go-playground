package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func unreliableOperation() error {
	if rand.Float32() < 0.7 {
		return errors.New("operation failed")
	}
	return nil
}

func retryWithBackoff(ctx context.Context, maxRetries int, baseDelay time.Duration) error {
	for i := 0; i < maxRetries; i++ {
		if err := unreliableOperation(); err != nil {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(baseDelay * time.Duration(i+1)):
				fmt.Println("Retrying after delay...")
			}
		} else {
			return nil
		}
	}
	return errors.New("all retries failed")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := retryWithBackoff(ctx, 5, 500*time.Millisecond)
	if err != nil {
		fmt.Println("Final error:", err)
	} else {
		fmt.Println("Operation succeeded")
	}
}
