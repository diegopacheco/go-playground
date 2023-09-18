package main

import (
	"context"
	"os"

	"github.com/a-h/templ-examples/hello-world/src/components"
)

func main() {
	component := components.Page()
	component.Render(context.Background(), os.Stdout)
}
