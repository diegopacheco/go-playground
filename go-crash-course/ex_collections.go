package main

import (
	"fmt"
	"maps"
	"slices"
)

func ExCollections() {
	fmt.Println("ex_collections")

	v := []int{3, 1, 2}
	slices.Sort(v)
	fmt.Println("sorted slice:", v)

	h := map[string]int{"a": 1}
	fmt.Println("map get a:", h["a"])

	h2 := maps.Clone(h)
	fmt.Println("map cloned equals:", maps.Equal(h, h2))

	keys := make([]string, 0, len(h))
	for k := range h {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, k := range keys {
		fmt.Println(k, ":", h[k])
	}
}
