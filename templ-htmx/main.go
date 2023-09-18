package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/a-h/templ-examples/hello-world/src/components"
)

func main() {
	component := components.Page(1, 1)
	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on http://127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
