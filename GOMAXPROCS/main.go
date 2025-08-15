package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("=== Container-Aware GOMAXPROCS ===")

	// Before Go 1.25: GOMAXPROCS was always set to logical CPU count
	// After Go 1.25: GOMAXPROCS respects container CPU limits
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())

	// In a Docker container with --cpus=2, GOMAXPROCS will be 2
	// even if the host has 8 CPUs
	// You can disable this behavior with GODEBUG settings:
	// GODEBUG=containermaxprocs=0 (ignore cgroups CPU quota)
	// GODEBUG=updatemaxprocs=0 (prevent GOMAXPROCS auto-update)
}
