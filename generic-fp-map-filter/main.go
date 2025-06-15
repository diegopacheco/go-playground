package main

import (
	"fmt"
)

func Filter[T any](items []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func Map[T any, R any](items []T, transform func(T) R) []R {
	var result []R
	for _, item := range items {
		result = append(result, transform(item))
	}
	return result
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	evens := Filter(numbers, func(n int) bool { return n%2 == 0 })
	squares := Map(evens, func(n int) int { return n * n })

	fmt.Println("Even Numbers:", evens)
	fmt.Println("Squares of Evens:", squares)
}
