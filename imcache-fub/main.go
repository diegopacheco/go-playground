package main

import (
	"fmt"

	"github.com/erni27/imcache"
)

func main() {
	// Zero value Cache is a valid non-sharded cache
	// with no expiration, no sliding expiration,
	// no entry limit and no eviction callback.
	var c imcache.Cache[uint32, string]
	// Set a new value with no expiration time.
	c.Set(1, "one", imcache.WithNoExpiration())
	// Get the value for the key 1.
	value, ok := c.Get(1)
	if !ok {
		panic("value for the key '1' not found")
	}
	fmt.Println(value)
}
