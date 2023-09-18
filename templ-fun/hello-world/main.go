package main

import (
	"context"
	"os"
)

func main() {
	component := hello("John Doe")
	component.Render(context.Background(), os.Stdout)
}
